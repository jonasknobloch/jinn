package corenlp

type Dependency struct {
	Dep            string `json:"dep"`
	Governor       int    `json:"governor"`
	GovernorGloss  string `json:"governorGloss"`
	Dependent      int    `json:"dependent"`
	DependentGloss string `json:"dependentGloss"`
}
