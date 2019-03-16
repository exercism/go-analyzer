package analyzer

import (
	"github.com/exercism/go-analyzer/suggester/sugg"
)

func getResult(goodPattern bool, suggReporter *sugg.SuggestionReport) Result {
	comments, severity := suggReporter.GetComments()
	if comments == nil {
		// make sure not to add nil to json
		comments = []string{}
	}
	errs := suggReporter.GetErrors()
	return Result{
		Status:   getStatus(goodPattern, len(comments), severity, len(errs)),
		Comments: comments,
		Errors:   errs,
		Severity: severity,
	}
}

func getStatus(goodPattern bool, comments, severity int, errors int) Status {
	switch {
	case errors != 0:
		return ReferToMentor
	case goodPattern && comments == 0:
		return ApproveAsOptimal
	case goodPattern && severity < 5:
		return ApproveWithComment
	case goodPattern:
		return DisapproveWithComment
	case 5 <= severity:
		return DisapproveWithComment
	}

	return ReferToMentor
}
