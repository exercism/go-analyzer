package raindrops

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/tehsphinx/astrav"
)

// Register registers all suggestion functions for this exercise.
var Register = sugg.Register{
	Funcs: []sugg.SuggestionFunc{
		examMainFunc,
		examAllCombis,
		examMultipleLoops,
		examLoopMap,
		examExtensiveForLoop,
		examItoa,
		examStringsBuilder,
		examBytesBuffer,
		examFmtPrintf,
		examRemoveExtraBool,
		sugg.ExamExtraVariable,
		examExtraFunction,
		examPlusEqual,
	},
	Severity: severity,
}

func examMainFunc(pkg *astrav.Package, suggs sugg.Suggester) {
	main := pkg.FuncDeclByName("Convert")
	if main == nil {
		suggs.AppendUnique(MissingEntryFunc)
		return
	}

	params := main.Params()
	if params == nil || len(params.Children()) != 1 {
		suggs.AppendUnique(FuncSignatureChanged)
	}
	results := main.Results()
	if results == nil || len(results.Children()) != 1 {
		suggs.AppendUnique(FuncSignatureChanged)
	}
}

var fmtPrintfRegex = regexp.MustCompile(`Pl.ng`)

func examAllCombis(pkg *astrav.Package, suggs sugg.Suggester) {
	lits := pkg.FindByNodeType(astrav.NodeTypeBasicLit)
	for _, lit := range lits {
		if lit.ValueType() == nil || lit.ValueType().String() != "string" {
			continue
		}

		occs := fmtPrintfRegex.FindAllString(lit.(*astrav.BasicLit).Value, -1)
		if 1 < len(occs) {
			suggs.AppendUnique(AllCombinations)
			return
		}
	}

	rets := pkg.FindByNodeType(astrav.NodeTypeReturnStmt)
	if 6 < len(rets) {
		suggs.AppendUnique(AllCombinations)
	}
}

func examFmtPrintf(pkg *astrav.Package, suggs sugg.Suggester) {
	nodes := pkg.FindByName("fmt.Sprintf")
	for _, node := range nodes {
		lits := node.Parent().FindByNodeType(astrav.NodeTypeBasicLit)
		for _, lit := range lits {
			if lit.ValueType().String() != "string" {
				continue
			}
			if fmtPrintfRegex.MatchString(lit.(*astrav.BasicLit).Value) {
				suggs.AppendUniquePH(ConcatFMT, map[string]string{
					"function": "fmt.Sprintf",
				})
			}
		}
	}
}

func examRemoveExtraBool(pkg *astrav.Package, suggs sugg.Suggester) {
	nodes := pkg.FindByValueType("bool")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeIdent) {
			continue
		}
		if node.Parent().IsNodeType(astrav.NodeTypeAssignStmt) {
			continue
		}
		if ifParent := node.NextParentByType(astrav.NodeTypeIfStmt); ifParent == nil {
			continue
		} else if ifParent.Level()+2 < node.Level() {
			continue
		} else if block := node.NextParentByType(astrav.NodeTypeBlockStmt); block != nil &&
			ifParent.Level() < block.Level() {
			continue
		} else if !ifParent.Parent().Parent().IsNodeType(astrav.NodeTypeFuncDecl) {
			continue
		}

		name := node.(*astrav.Ident).Name
		suggs.AppendUniquePH(RemoveExtraBool, map[string]string{
			"name": name,
		})
		break
	}
}

func examLoopMap(pkg *astrav.Package, suggs sugg.Suggester) {
	if suggs.HasSuggestion(MultipleLoops) {
		return
	}

	loops := pkg.FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, l := range loops {
		loop := l.(*astrav.RangeStmt)

		ident, ok := loop.X().(*astrav.Ident)
		if !ok {
			continue
		}
		if strings.HasPrefix(ident.ValueType().String(), "map") {
			suggs.AppendUnique(LoopMap)
			return
		}
	}
}

func examExtensiveForLoop(pkg *astrav.Package, suggs sugg.Suggester) {
	f := pkg.FindFirstByName("Convert")
	params := f.(*astrav.FuncDecl).Params()
	if len(params.List) == 0 || len(params.List[0].Names) == 0 {
		return
	}
	paramName := params.List[0].Names[0].Name

	loops := pkg.FindByNodeType(astrav.NodeTypeForStmt)
	for _, l := range loops {
		loop := l.(*astrav.ForStmt)

		// check if loop goes up to input var
		if loop.Cond().FindFirstByName(paramName) != nil {
			suggs.AppendUnique(ExtensiveFor)
			return
		}

		// if using a basiclit it should be 7 or 8 depending on the operator
		basicLit := loop.Cond().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if basicLit != nil {
			val := basicLit.(*astrav.BasicLit).Value
			if n, err := strconv.Atoi(val); err == nil && n < 7 {
				return
			}
			if val != "7" && val != "8" {
				suggs.AppendUnique(ExtensiveFor)
				return
			}
		}

		// check if loop starts with 3
		if loop.Init() == nil {
			// Probably a condition-only loop
			return
		}
		basicLit = loop.Init().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if basicLit != nil {
			if basicLit.(*astrav.BasicLit).Value != "3" {
				suggs.AppendUnique(ExtensiveFor)
				return
			}
		}

		// check if loop uses steps of 2: 3, 5, 7
		if loop.Post() == nil {
			// Probably a condition-only loop
			return
		}
		if loop.Post().IsNodeType(astrav.NodeTypeIncDecStmt) {
			suggs.AppendUnique(ExtensiveFor)
			return
		}
		if loop.Post().IsNodeType(astrav.NodeTypeAssignStmt) {
			basicLit := loop.Post().FindFirstByNodeType(astrav.NodeTypeBasicLit)
			if basicLit != nil {
				if basicLit.(*astrav.BasicLit).Value != "2" {
					suggs.AppendUnique(ExtensiveFor)
					return
				}
			}
		}
	}
}

func examItoa(pkg *astrav.Package, suggs sugg.Suggester) {
	itoa := pkg.FindFirstByName("strconv.Itoa")
	if itoa != nil {
		return
	}

	suggs.AppendUnique(UseItoa)
}

func examPlusEqual(pkg *astrav.Package, suggs sugg.Suggester) {
	assigns := pkg.FindByNodeType(astrav.NodeTypeAssignStmt)
	for _, assign := range assigns {
		token := assign.(*astrav.AssignStmt).Tok.String()
		if token != "=" {
			continue
		}

		binExpr := assign.FindFirstByNodeType(astrav.NodeTypeBinaryExpr)
		if binExpr == nil {
			continue
		}

		if binExpr.(*astrav.BinaryExpr).Op.String() == "+" {
			suggs.AppendUnique(PlusEqual)
			return
		}
	}
}

func examMultipleLoops(pkg *astrav.Package, suggs sugg.Suggester) {
	var count int
	count += len(pkg.FindByNodeType(astrav.NodeTypeForStmt))
	count += len(pkg.FindByNodeType(astrav.NodeTypeRangeStmt))

	if 1 < count {
		suggs.AppendUnique(MultipleLoops)
	}
}

func examStringsBuilder(pkg *astrav.Package, suggs sugg.Suggester) {
	builder := pkg.FindByName("Builder")
	if builder != nil {
		suggs.AppendUniquePH(StringBuilder, map[string]string{
			"function": "strings.Builder",
		})
	}
}

func examBytesBuffer(pkg *astrav.Package, suggs sugg.Suggester) {
	buffer := pkg.FindByName("Buffer")
	if buffer != nil {
		suggs.AppendUniquePH(StringBuilder,
			map[string]string{
				"function": "bytes.Buffer",
			})
	}
}

func examExtraFunction(pkg *astrav.Package, suggs sugg.Suggester) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeFuncDecl)
	main := pkg.FuncDeclByName("main")
	if 1 < len(nodes) && main == nil {
		suggs.AppendUnique(sugg.ExtraFunction)
	}
}
