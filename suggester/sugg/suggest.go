package sugg

import (
	"strings"

	"github.com/tehsphinx/astrav"
)

// GeneralRegister registers all suggestion functions for this exercise.
var GeneralRegister = Register{
	Funcs: []SuggestionFunc{
		examGoFmt,
		examGoLint,
		examMainFunction,
		examStringLenComparison,
		examNoErrorMsg,
		examErrorfWithoutParams,
		examCustomError,
		examStringsCompare,
	},
	Severity: severity,
}

// checks if gofmt was applied
func examGoFmt(pkg *astrav.Package, suggs Suggester) {
	files := pkg.GetRawFiles()
	res := fmtCode(files)
	if res == "" {
		return
	}

	suggs.AppendUniquePH(GoFmt, map[string]string{
		"gofmt": res,
	})
}

// checks if golint is happy
func examGoLint(pkg *astrav.Package, suggs Suggester) {
	files := pkg.GetRawFiles()
	res := lintCode(files)
	if res == "" {
		return
	}

	suggs.AppendUniquePH(GoLint, map[string]string{
		"golint": res,
	})
}

// checks if a `main` function was declared
func examMainFunction(pkg *astrav.Package, suggs Suggester) {
	mainFunc := pkg.FuncDeclByName("main")
	if mainFunc != nil {
		suggs.AppendUnique(MainFunction)
	}
}

// examins if emptyness of a string was checked by taking its length
func examStringLenComparison(pkg *astrav.Package, suggs Suggester) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeBinaryExpr)
	for _, node := range nodes {
		bin := node.(*astrav.BinaryExpr)
		op := bin.Op.String()
		if op != "==" && op != "!=" &&
			op != "<=" && op != "<" &&
			op != ">=" && op != ">" {
			continue
		}
		// check if there are 2 idents ("len" and variable name)
		idents := bin.FindByNodeType(astrav.NodeTypeIdent)
		if len(idents) < 2 {
			continue
		}
		// check if one of the idents is "len" and the other one is of type string
		var (
			foundLen    bool
			foundString bool
			varName     string
		)
		for _, ident := range idents {
			id := ident.(*astrav.Ident)
			if id.NodeName() != "len" {
				varName = id.NodeName()
			}
			if id.NodeName() == "len" {
				foundLen = true
			} else if id.ValueType().String() == "string" {
				foundString = true
			}
		}
		if !foundLen {
			continue
		}
		// Check if a basicLit exists and it is 0
		basic := bin.FindByNodeType(astrav.NodeTypeBasicLit)
		if len(basic) != 1 {
			continue
		}

		basicVal := basic[0].(*astrav.BasicLit).Value
		if op == "<" && basicVal == "1" {
			suggs.AppendUnique(LengthSmallerZero)
		}
		if op == ">=" && basicVal == "1" {
			suggs.AppendUniquePH(LenOfStringEqual, map[string]string{
				"name": varName,
			})
		}
		if basicVal != "0" {
			continue
		}
		if op == "<=" {
			suggs.AppendUnique(LengthSmallerZero)
		}
		if foundString {
			suggs.AppendUniquePH(LenOfStringEqual, map[string]string{
				"name": varName,
			})
		}
	}
}

// check if an empty error message was provided
func examNoErrorMsg(pkg *astrav.Package, suggs Suggester) {
	nodes := pkg.FindByName("errors.New")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeSelectorExpr) {
			continue
		}
		selExpr := node.(*astrav.SelectorExpr)
		if selExpr.PackageName().Name != "errors" {
			continue
		}

		bLit := selExpr.Parent().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if bLit == nil {
			continue
		}

		if bLit.(*astrav.BasicLit).Value == `""` {
			suggs.AppendUnique(OmittedErrorMsg)
		}
	}
}

// checks if fmt.Errorf was used without params instead of errors.New
func examErrorfWithoutParams(pkg *astrav.Package, suggs Suggester) {
	nodes := pkg.FindByName("fmt.Errorf")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeSelectorExpr) {
			continue
		}
		var count int
		for _, ch := range node.Parent().Children() {
			if ch.IsNodeType(astrav.NodeTypeSelectorExpr) {
				continue
			}
			count++
		}
		if count == 1 {
			suggs.AppendUnique(ErrorfWithoutParam)
		}
	}
}

// checks a custom error was created
func examCustomError(pkg *astrav.Package, suggs Suggester) {
	nodes := pkg.FindByName("Error")
	for _, node := range nodes {
		if !node.IsNodeType(astrav.NodeTypeFuncDecl) {
			continue
		}
		funcType := node.ChildByNodeType(astrav.NodeTypeFuncType)
		if funcType == nil {
			continue
		}
		if strings.HasSuffix(funcType.GetSourceString(), "Error() string") {
			suggs.AppendUnique(CustomErrorCreated)
		}
	}
}

var excludeVarTypes = []string{
	"*bufio.Reader",
}

// ExamExtraVariable checks for a variable that can be inlined
func ExamExtraVariable(pkg *astrav.Package, suggs Suggester) {
	decls := pkg.FindVarDeclarations()
	for _, decl := range decls {
		if isExcludeType(decl) {
			continue
		}

		var (
			usageCount = len(pkg.FindUsages(decl))
			firstUsage = pkg.FindFirstUsage(decl)
		)
		if firstUsage == nil {
			continue
		}
		if canBeCombined(decl, firstUsage) {
			usageCount--
			suggs.AppendUniquePH(UseVarAssignment, map[string]string{
				"name": decl.Name,
				"line": decl.NextParentByType(astrav.NodeTypeGenDecl).GetSourceString(),
			})
		}

		if usageCount == 1 {
			_, declScope := decl.GetScope()
			_, usageScope := firstUsage.GetScope()
			if usageScope != declScope {
				continue
			}
			if IsMultiAssignment(decl.Parent()) {
				continue
			}
			suggs.AppendUniquePH(ExtraVar, map[string]string{
				"name": decl.Name,
			})
		}
	}
}

// IsMultiAssignment declaration assigns to multiple variables
func IsMultiAssignment(decl astrav.Node) bool {
	assign, ok := decl.(*astrav.AssignStmt)
	if !ok {
		return false
	}

	return len(assign.LHS()) != 1
}

func canBeCombined(decl, firstUsage *astrav.Ident) bool {
	// only if the declaration itself is not an assign statement,
	// check if we can make it an assign statement because we assign it after anyway
	if decl.Parent().IsNodeType(astrav.NodeTypeAssignStmt) {
		return false
	}
	if !firstUsage.Parent().IsNodeType(astrav.NodeTypeAssignStmt) {
		return false
	}
	_, usageScope := firstUsage.GetScope()
	_, declScope := decl.GetScope()
	if usageScope != declScope {
		return false
	}

	lhs := firstUsage.Parent().(*astrav.AssignStmt).LHS()
	return len(lhs) == 1 && lhs[0] == firstUsage
}

func isExcludeType(node astrav.Node) bool {
	for _, varType := range excludeVarTypes {
		if node.IsValueType(varType) {
			return true
		}
	}
	return false
}

func examStringsCompare(pkg *astrav.Package, suggs Suggester) {
	cmp := pkg.FindByName("strings.Compare")
	if len(cmp) != 0 {
		suggs.AppendUnique(StringsCompare)
	}
}
