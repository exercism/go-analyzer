package analyzer

type limits struct {
	OptimalLimit float64
	ApproveLimit float64
}

var patternLimits = map[string]limits{
	"two-fer": {
		OptimalLimit: 0.99,
		ApproveLimit: 0.9,
	},
	"hamming": {
		OptimalLimit: 0.99,
		ApproveLimit: 0.9,
	},
	"raindrops": {
		OptimalLimit: 0.95,
		ApproveLimit: 0.85,
	},
	"concepts-basicstrings": {
		OptimalLimit: 0.99,
		ApproveLimit: 0.9,
	},
}
