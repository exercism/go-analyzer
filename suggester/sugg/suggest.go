package sugg

import (
	"github.com/tehsphinx/astrav"
)

// GeneralRegister registers all suggestion functions for this exercise.
var GeneralRegister = Register{
	Funcs: []SuggestionFunc{
		examGoFmt,
		examGoLint,
		examMainFunction,
		examEmptyByLenOfString,
	},
	Severity: severity,
}

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

func examMainFunction(pkg *astrav.Package, suggs Suggester) {
	mainFunc := pkg.FuncDeclByName("main")
	if mainFunc != nil {
		suggs.AppendUnique(MainFunction)
	}
}

func examEmptyByLenOfString(pkg *astrav.Package, suggs Suggester) {
	nodes := pkg.FindByNodeType(astrav.NodeTypeBinaryExpr)
	for _, node := range nodes {
		bin := node.(*astrav.BinaryExpr)
		op := bin.Op.String()
		if op != "==" && op != "!=" {
			continue
		}
		// check if there are 2 idents ("len" and variable name)
		idents := bin.FindByNodeType(astrav.NodeTypeIdent)
		if len(idents) < 2 {
			continue
		}
		// check if one of the idents is "len" and the other one is of type string
		var foundLen bool
		for _, ident := range idents {
			id := ident.(*astrav.Ident)
			if id.NodeName().String() == "len" {
				foundLen = true
			} else {
				if id.ValueType().String() != "string" {
					continue
				}
			}
		}
		if !foundLen {
			continue
		}
		// Check if a basicLit exists and it is 0
		basic := bin.FindByNodeType(astrav.NodeTypeBasicLit)
		if len(basic) != 1 || basic[0].(*astrav.BasicLit).Value != "0" {
			continue
		}

		suggs.AppendUnique(LenOfStringEqual)
	}
}
