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
		examExtraVariable,
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
			op != "<=" && op != "<" {
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
		)
		for _, ident := range idents {
			id := ident.(*astrav.Ident)
			if id.NodeName().String() == "len" {
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
		if basicVal != "0" {
			continue
		}
		if op == "<=" {
			suggs.AppendUnique(LengthSmallerZero)
		}
		if foundString {
			suggs.AppendUnique(LenOfStringEqual)
		}
	}
}

// check if an empty error message was provided
func examNoErrorMsg(pkg *astrav.Package, suggs Suggester) {
	nodes := pkg.FindByName("New")
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
	nodes := pkg.FindByName("Errorf")
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

func examExtraVariable(pkg *astrav.Package, suggs Suggester) {
	decls := pkg.FindVarDeclarations()
	for _, decl := range decls {
		_, declScope := decl.GetScope()
		usages := pkg.FindUsages(decl)
		usageCount := len(usages)
		for _, usage := range usages {
			if !usage.Parent().IsNodeType(astrav.NodeTypeAssignStmt) {
				continue
			}
			_, usageScope := usage.GetScope()
			lhs := usage.Parent().(*astrav.AssignStmt).LHS()
			if len(lhs) == 1 && lhs[0] == usage && usageScope == declScope {
				usageCount--
				suggs.AppendUniquePH(UseVarAssignment, map[string]string{
					"name": decl.Name,
					"line": decl.Parent().GetSourceString(),
				})
			}
		}
		if usageCount < 2 {
			suggs.AppendUniquePH(ExtraVar, map[string]string{
				"name": decl.Name,
			})
		}
	}
}
