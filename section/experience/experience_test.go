package experience

import (
	"resumebuilder/resumebuilder"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	e := &Experience{
		From:     time.Now(),
		To:       time.Now(),
		Title:    "Programmer",
		Company:  "Microsoft",
		Subtitle: []string{"Go", "Bash", "Python"},
		Points:   []string{"worked hard", "did stuff"},
	}
	result := resumebuilder.Parse(e)
	if result == "" {
		t.Error("result should not be empty")
	}
}
