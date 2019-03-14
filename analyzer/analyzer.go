package analyzer

// Result defines the result of the analyzer
type Result struct {
	Status   Status   `json:"status"`
	Comments []string `json:"comments"`
	Error    error    `json:"-"`
}

// Status defines the status of a solution to be acted upon by exercism.
type Status string

// status constants
const (
	ApproveAsOptimal   Status = "approve_as_optimal"
	ApproveWithComment Status = "approve_with_comment"
	ReferToMentor      Status = "refer_to_mentor"
	// DisapproveWithComment Status = "disapprove_with_comment"
)

// Analyze analyses a solution and returns the Result
func Analyze(exercise string, path string) Result {
	solution, err := LoadPackage(path)
	if err != nil {
		return Result{Error: err}
	}
	goodPattern, err := CheckPattern(exercise, solution)
	if err != nil {
		return Result{Error: err}
	}

	// 	TODO:
	// 	 - import and run suggest package on solution (can be 2nd step)

	return getResult(goodPattern, nil)
}

func getResult(goodPattern bool, comments []string) Result {
	if comments == nil {
		// make sure not to add nil to json
		comments = []string{}
	}
	return Result{
		Status:   getStatus(goodPattern, comments),
		Comments: comments,
	}
}

func getStatus(goodPattern bool, comments []string) Status {
	switch {
	case goodPattern && len(comments) == 0:
		return ApproveAsOptimal
	case goodPattern:
		return ApproveWithComment
	}

	return ReferToMentor
}
