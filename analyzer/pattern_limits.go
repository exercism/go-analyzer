package analyzer

var patternLimits = map[string]patternReport{
	"two-fer": {
		OptimalLimit: 0.99,
		ApproveLimit: 0.9,
	},
	"hamming": {
		OptimalLimit: 0.99,
		ApproveLimit: 0.9,
	},
}
