package hamming

import (
	"strings"
	"unicode"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/tehsphinx/astrav"
)

// Register registers all suggestion functions for this exercise.
var Register = sugg.Register{
	Funcs: []sugg.SuggestionFunc{
		examMainFunc,
		examReturns,
		examNoErrorMsg,
		examInvertIf,
		examDeclareWhenNeeded,
		examRuneToByte,
		examMultipleStringConversions,
		examErrorMessage,
		examIncrease,
	},
	Severity: severity,
}

func examMainFunc(pkg *astrav.Package, suggs sugg.Suggester) {
	main := pkg.FuncDeclByName("Distance")
	if main == nil {
		suggs.AppendUnique(MissingMainFunc)
		return
	}

	if len(main.Params().Children()) != 2 {
		suggs.AppendUnique(FuncSignatureChanged)
	}
	if len(main.Results().Children()) != 2 {
		suggs.AppendUnique(FuncSignatureChanged)
	}
}

func examReturns(pkg *astrav.Package, suggs sugg.Suggester) {
	main := pkg.FindFirstByName("Distance")
	if main == nil {
		suggs.AppendUnique(MissingMainFunc)
		return
	}

	returns := main.FindByNodeType(astrav.NodeTypeReturnStmt)
	for _, ret := range returns {
		if len(ret.Children()) == 0 {
			suggs.AppendUnique(NakedReturns)
			continue
		}
		if len(ret.Children()) < 2 {
			continue
		}

		if ret.Children()[1].ValueType().String() == "nil" {
			continue
		}
		switch node := ret.Children()[0].(type) {
		case *astrav.BasicLit:
			if node.Value != "0" {
				suggs.AppendUnique(ZeroValueOnErr)
			}
		case *astrav.UnaryExpr:
			lit := node.FindFirstByNodeType(astrav.NodeTypeBasicLit)
			if lit == nil {
				continue
			}
			if lit.(*astrav.BasicLit).Value != "0" {
				suggs.AppendUnique(ZeroValueOnErr)
			}
		}
	}
}

func examNoErrorMsg(pkg *astrav.Package, suggs sugg.Suggester) {
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
			suggs.AppendUnique(ErrorMsgFormat)
		}
	}
}

func examDeclareWhenNeeded(pkg *astrav.Package, suggs sugg.Suggester) {
	if suggs.HasSuggestion(InvertIf) {
		return
	}

	main := pkg.FindFirstByName("Distance")
	if main == nil {
		return
	}
	returns := main.FindByNodeType(astrav.NodeTypeReturnStmt)
	for _, ret := range returns {
		for _, child := range ret.Children() {
			if !child.IsNodeType(astrav.NodeTypeIdent) {
				continue
			}
			returnVar := child.(*astrav.Ident)
			if returnVar.Obj == nil {
				continue
			}

			varDecl := main.FindFirstByName(returnVar.Name).Parent()

			// variable not declared in the same block as the return statement
			if varDecl.IsNodeType(astrav.NodeTypeAssignStmt) {
				if !returnVar.NextParentByType(astrav.NodeTypeBlockStmt).Contains(varDecl) {
					suggs.AppendUniquePH(DeclareWhenNeeded, map[string]string{
						"returnVar": returnVar.Name,
					})
				}
			}

			// there is another return inbetween
			for _, rt := range returns {
				if rt == ret {
					continue
				}
				if varDecl.Pos() <= rt.Pos() && rt.Pos() <= returnVar.Pos() {
					suggs.AppendUniquePH(DeclareWhenNeeded, map[string]string{
						"returnVar": returnVar.Name,
					})
				}
			}
		}
	}
}

func examInvertIf(pkg *astrav.Package, suggs sugg.Suggester) {
	for _, ifNode := range pkg.FindByNodeType(astrav.NodeTypeIfStmt) {
		loop := ifNode.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
		if loop == nil {
			loop = ifNode.FindFirstByNodeType(astrav.NodeTypeForStmt)
		}
		binExpr := ifNode.ChildByNodeType(astrav.NodeTypeBinaryExpr)
		if binExpr == nil {
			continue
		}
		condition := binExpr.(*astrav.BinaryExpr)
		if loop != nil && condition != nil && condition.Op.String() == "==" {
			suggs.AppendUnique(InvertIf)
		}
	}
}

func examRuneToByte(pkg *astrav.Package, suggs sugg.Suggester) {
	nodes := pkg.FindByName("byte")
	for _, node := range nodes {
		parentType := node.Parent().NodeType()
		if parentType != astrav.NodeTypeCallExpr {
			continue
		}
		for _, n := range node.Siblings() {
			if n.ValueType().String() == "rune" {
				suggs.AppendUnique(RuneToByte)
			}
		}
	}
}

func examMultipleStringConversions(pkg *astrav.Package, suggs sugg.Suggester) {
	rngNode := pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
	if rngNode == nil {
		return
	}

	count := 0
	for _, node := range rngNode.FindByName("string") {
		if node.Parent().IsNodeType(astrav.NodeTypeCallExpr) {
			count++
		}
	}
	if 1 < count {
		suggs.AppendUnique(ToStringConversion)
	}
}

func examIncrease(pkg *astrav.Package, suggs sugg.Suggester) {
	loop := pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
	if loop == nil {
		loop = pkg.FindFirstByNodeType(astrav.NodeTypeForStmt)
	}
	if loop == nil {
		return
	}
	for _, node := range loop.FindByNodeType(astrav.NodeTypeBinaryExpr) {
		if node.(*astrav.BinaryExpr).Op.String() == "+" {
			suggs.AppendUnique(IncreaseOperator)
		}
	}
}

func examErrorMessage(pkg *astrav.Package, suggs sugg.Suggester) {
	checkErrMessage(pkg.FindByName("New"), suggs)
	checkErrMessage(pkg.FindByName("Errorf"), suggs)
}

func checkErrMessage(nodes []astrav.Node, suggs sugg.Suggester) {
	for _, node := range nodes {
		if node.NodeType() == astrav.NodeTypeIdent {
			continue
		}
		errMsgNode := node.(*astrav.SelectorExpr).Parent().FindFirstByNodeType(astrav.NodeTypeBasicLit)
		if errMsgNode == nil {
			continue
		}

		errText := errMsgNode.(*astrav.BasicLit).Value

		// check punctuation
		if strings.HasSuffix(errText, ".") {
			suggs.AppendUnique(ErrorMsgFormat)
			continue
		}

		// check if first letter is capitalized and second not.
		var isUpper bool
		for i, rn := range strings.Split(errText, " ")[0] {
			// first letter is " or `
			if i == 2 {
				if isUpper && !unicode.IsUpper(rn) {
					suggs.AppendUnique(ErrorMsgFormat)
				}
				break
			}
			isUpper = unicode.IsUpper(rn)
		}
	}
}
