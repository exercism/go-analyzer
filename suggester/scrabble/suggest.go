package scrabble

import (
	"go/ast"
	"reflect"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/tehsphinx/astrav"
)

// Register registers all suggestion functions for this exercise.
var Register = sugg.Register{
	Funcs: []sugg.SuggestionFunc{
		examMainFunc,
		testGoRoutine,
		testRegex,
		testMapInFunc,
		testMultipleLoops,
		testMapRuneInt,
		testSliceRuneConv,
		sugg.ExamExtraVariable,
		testToLowerUpper("strings.ToLower"),
		testToLowerUpper("strings.ToUpper"),
		testIfElseToSwitch,
		testRuneLoop,
		testTrySwitch,
	},
	Severity: severity,
}

func examMainFunc(pkg *astrav.Package, suggs sugg.Suggester) {
	main := pkg.FuncDeclByName("Score")
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

func testSliceRuneConv(pkg *astrav.Package, suggs sugg.Suggester) {
	calls := pkg.FindByNodeType(astrav.NodeTypeCallExpr)
	for _, call := range calls {
		if reflect.TypeOf(call.(*astrav.CallExpr).Fun).String() != "*ast.ArrayType" {
			continue
		}

		if call.(*astrav.CallExpr).NodeName() == "[]rune" {
			suggs.AppendUnique(SliceRuneConv)
		}
	}
}

func testRegex(pkg *astrav.Package, suggs sugg.Suggester) {
	rgx := pkg.FindFirstByName("Score").FindFirstByName("regexp.MustCompile")
	if rgx == nil {
		rgx = pkg.FindFirstByName("Score").FindFirstByName("regexp.Compile")
	}

	if rgx != nil &&
		rgx.(*astrav.SelectorExpr).X != nil &&
		rgx.(*astrav.SelectorExpr).X.(*ast.Ident).Name == "regexp" {

		lit := rgx.Parent().ChildByNodeType(astrav.NodeTypeBasicLit)
		if lit != nil && lit.IsValueType("string") {
			// is a static regex
			suggs.AppendUnique(Regex)
			suggs.AppendUnique(RegexChallenge)
			suggs.AppendUnique(sugg.RegexComment)
			suggs.AppendUnique(sugg.BenchmarkComment)
		}
	}
}

func testToLowerUpper(fnName string) sugg.SuggestionFunc {
	return func(pkg *astrav.Package, suggs sugg.Suggester) {
		if suggs.HasSuggestion(Regex) {
			return
		}

		fns := pkg.FindByName(fnName)
		for _, fn := range fns {
			if _, ok := fn.(*astrav.SelectorExpr); !ok {
				continue
			}
			suggs.AppendUnique(sugg.BenchmarkComment)

			fName := fnName[8:]
			if fn.NextParentByType(astrav.NodeTypeBlockStmt).IsContainedByType(astrav.NodeTypeRangeStmt) {
				suggs.AppendUniquePH(UnicodeLoop, map[string]string{
					"function": fName,
				})
			} else {
				suggs.AppendUniquePH(Unicode, map[string]string{
					"function": fName,
				})
			}
		}
	}
}

func testMultipleLoops(pkg *astrav.Package, suggs sugg.Suggester) {
	loopCount := len(pkg.FindByNodeType(astrav.NodeTypeForStmt))
	loopCount += len(pkg.FindByNodeType(astrav.NodeTypeRangeStmt))
	if 1 < loopCount {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUnique(MultipleLoops)
	}
}

func testMapRuneInt(pkg *astrav.Package, suggs sugg.Suggester) {
	if suggs.HasSuggestion(MultipleLoops) {
		return
	}
	if len(pkg.FindByValueType("map[string]int")) != 0 {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUniquePH(MapRune, map[string]string{
			"type": "map[string]int",
		})
		return
	}
	if len(pkg.FindByValueType("map[int]string")) != 0 {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUniquePH(MapRune, map[string]string{
			"type": "map[int]string",
		})
		return
	}
}

func testTrySwitch(pkg *astrav.Package, suggs sugg.Suggester) {
	if suggs.HasSuggestion(MultipleLoops) {
		return
	}
	if len(pkg.FindByValueType("map[rune]int")) != 0 {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUnique(TrySwitch)
		return
	}
	if len(pkg.FindByValueType("map[string]int")) != 0 {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUnique(TrySwitch)
		return
	}
	if len(pkg.FindByValueType("map[int]string")) != 0 {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUnique(TrySwitch)
		return
	}
}

func testRuneLoop(pkg *astrav.Package, suggs sugg.Suggester) {
	ranges := pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		l := rng.(*astrav.RangeStmt)
		if l.Value() == nil {
			if l.Key() != nil {
				suggs.AppendUnique(LoopRuneNotByte)
			}
		} else {
			var isByte bool
			for _, ident := range rng.FindIdentByName(l.Value().(*astrav.Ident).Name) {
				if ident.IsValueType("byte") {
					isByte = true
				}
			}
			if isByte {
				suggs.AppendUnique(LoopRuneNotByte)
			}
		}

		if rng.FindByName("string") != nil &&
			!suggs.HasSuggestion(MapRune) &&
			!suggs.HasSuggestion(MultipleLoops) {

			suggs.AppendUnique(TypeConversion)
			return
		}
	}
}

func testGoRoutine(pkg *astrav.Package, suggs sugg.Suggester) {
	goStmts := pkg.FindByNodeType(astrav.NodeTypeGoStmt)
	if len(goStmts) != 0 {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUnique(GoRoutines)
	}
}

func testIfElseToSwitch(pkg *astrav.Package, suggs sugg.Suggester) {
	ranges := pkg.FindFirstByName("Score").FindByNodeType(astrav.NodeTypeRangeStmt)
	for _, rng := range ranges {
		ifs := rng.FindByNodeType(astrav.NodeTypeIfStmt)
		if 5 < len(ifs) {
			suggs.AppendUnique(IfsToSwitch)
			return
		}
	}
}

func testMapInFunc(pkg *astrav.Package, suggs sugg.Suggester) {
	fn := pkg.FindFirstByName("Score")
	maps := fn.FindMaps()

	var hasMapDef bool
	for _, m := range maps {
		if !m.IsNodeType(astrav.NodeTypeIdent) {
			hasMapDef = true
		}
	}
	if hasMapDef {
		suggs.AppendUnique(sugg.BenchmarkComment)
		suggs.AppendUnique(MoveMap)
	}
}
