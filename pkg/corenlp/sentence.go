package corenlp

type Sentence struct {
	Index                        int          `json:"index"`
	Parse                        string       `json:"parse"`
	BasicDependencies            []Dependency `json:"basicDependencies"`
	EnhancedDependencies         []Dependency `json:"enhancedDependencies"`
	EnhancedPlusPlusDependencies []Dependency `json:"enhancedPlusPlusDependencies"`
	Tokens                       []Token      `json:"tokens"`
}
