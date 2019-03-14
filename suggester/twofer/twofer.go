package twofer

import (
	"bytes"
	"errors"

	"github.com/exercism/go-analyzer/suggester/suggTypes"
	"github.com/tehsphinx/astrav"
)

// FuncRegister registers all suggestion functions for this exercise.
var FuncRegister = []suggTypes.SuggestionFunc{
	examPlusUsed,
	// examGeneralizeNames,
	examFmt,
	// examComments,
	// examConditional,
	examStringsJoin,
}

func examStringsJoin(pkg *astrav.Package) (suggs suggTypes.Suggestions, err error) {
	node := pkg.FindFirstByName("Join")
	if node != nil {
		suggs.AppendUnique(StringsJoin, severity)
	}
	return suggs, nil
}

func examPlusUsed(pkg *astrav.Package) (suggs suggTypes.Suggestions, err error) {
	main := pkg.FindFirstByName("ShareWith")
	if main == nil {
		return nil, errors.New("main function 'ShareWith' not found")
	}
	nodes := main.FindByNodeType(astrav.NodeTypeBinaryExpr)

	var plusUsed bool
	for _, node := range nodes {
		expr, ok := node.(*astrav.BinaryExpr)
		if !ok {
			continue
		}
		if expr.Op.String() == "+" {
			plusUsed = true
		}
	}
	if plusUsed {
		suggs.AppendUnique(PlusUsed, severity)
	}
	return suggs, nil
}

func examFmt(pkg *astrav.Package) (suggs suggTypes.Suggestions, err error) {
	nodes := pkg.FindByName("Sprintf")

	var spfCount int
	for _, fmtSprintf := range nodes {
		if !fmtSprintf.IsNodeType(astrav.NodeTypeSelectorExpr) {
			continue
		}

		spfCount++
		if 1 < spfCount {
			suggs.AppendUnique(MinimalConditional, severity)
		}
	}

	nodes = pkg.FindByNodeType(astrav.NodeTypeBasicLit)
	for _, node := range nodes {
		bLit := node.(*astrav.BasicLit)
		if bytes.Contains(bLit.GetSource(), []byte("%v")) {
			suggs.AppendUnique(UseStringPH, severity)
		}
	}
	return suggs, nil
}

// func examComments(pkg *astrav.Package, r *extypes.Response) {
// 	if bytes.Contains(pkg.GetSource(), []byte("stub")) {
// 		addStub(r)
// 	}
//
// 	cGroup := pkg.ChildByNodeType(astrav.NodeTypeFile).ChildByNodeType(astrav.NodeTypeCommentGroup)
// 	checkComment(cGroup, r, "package", "twofer")
//
// 	cGroup = pkg.FindFirstByName("ShareWith").ChildByNodeType(astrav.NodeTypeCommentGroup)
// 	checkComment(cGroup, r, "function", "ShareWith")
// }
//
// var outputPart = regexp.MustCompile(`, one for me\.`)
//
// func examConditional(pkg *astrav.Package, r *extypes.Response) {
// 	matches := outputPart.FindAllIndex(pkg.FindFirstByName("ShareWith").GetSource(), -1)
// 	if 1 < len(matches) {
// 		r.AppendImprovement(tpl.MinimalConditional)
// 	}
// }
//
// func examGeneralizeNames(pkg *astrav.Package, r *extypes.Response) {
// 	contains := bytes.Contains(pkg.FindFirstByName("ShareWith").GetSource(), []byte("Alice"))
// 	if !contains {
// 		contains = bytes.Contains(pkg.FindFirstByName("ShareWith").GetSource(), []byte("Bob"))
// 	}
// 	if contains {
// 		r.AppendTodo(tpl.GeneralizeName)
// 	}
// }
//
// var commentStrings = map[string]struct {
// 	typeString       string
// 	stubString       string
// 	prefixString     string
// 	wrongCommentName string
// }{
// 	"package": {
// 		typeString:       "Packages",
// 		stubString:       "should have a package comment",
// 		prefixString:     "Package %s ",
// 		wrongCommentName: "package `%s`",
// 	},
// 	"function": {
// 		typeString:       "Exported functions",
// 		stubString:       "should have a comment",
// 		prefixString:     "%s ",
// 		wrongCommentName: "function `%s`",
// 	},
// }
//
// // we only do this on the first exercise. Later we ask them to use golint.
// func checkComment(cGroup astrav.Node, r *extypes.Response, commentType, name string) {
// 	strPack := commentStrings[commentType]
// 	if cGroup == nil {
// 		r.AppendImprovement(tpl.MissingComment.Format(strPack.typeString))
// 		addCommentFormat(r)
// 	} else {
// 		text := cGroup.Children()[0].(*astrav.Suggestion).Text
// 		c := strings.TrimSpace(strings.Replace(strings.Replace(text, "/*", "", 1), "//", "", 1))
//
// 		if strings.Contains(c, strPack.stubString) {
// 			addStub(r)
// 		} else if !strings.HasPrefix(c, fmt.Sprintf(strPack.prefixString, name)) {
// 			r.AppendImprovement(tpl.WrongCommentFormat.Format(fmt.Sprintf(strPack.wrongCommentName, name)))
// 			addCommentFormat(r)
// 		}
// 	}
// }
//
// var (
// 	addStub          func(r *extypes.Response)
// 	addCommentFormat func(r *extypes.Response)
// )
//
// func getAddStub() func(r *extypes.Response) {
// 	var added bool
// 	return func(r *extypes.Response) {
// 		if added {
// 			return
// 		}
// 		added = true
// 		r.AppendImprovement(tpl.Stub)
// 	}
// }
// func getAddCommentFormat() func(r *extypes.Response) {
// 	var added bool
// 	return func(r *extypes.Response) {
// 		if added {
// 			return
// 		}
// 		added = true
// 		r.AppendOutro(tpl.CommentFormat)
// 	}
// }
