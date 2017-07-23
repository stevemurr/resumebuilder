package skills

import (
	"resumebuilder/resumebuilder"
	"testing"
)

func TestParse(t *testing.T) {
	s := &Skills{
		Proficient: []string{"Go", "Python"},
		Familiar:   []string{"Python", "Go"},
		Frameworks: []string{"bootstrap", "jQuery"},
		Libraries:  []string{"bootstrap"},
		Devices:    []string{"Raspberry PI"},
	}
	if result := resumebuilder.Parse(s); result == "" {
		t.Error("result should not be empty")
	}
}
