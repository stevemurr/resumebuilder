package education

import (
	"resumebuilder/resumebuilder"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	e := &Education{
		From:        time.Now(),
		To:          time.Now(),
		Institution: "San Jose State",
		Degree:      "Bachleors in Compute Science",
	}
	result := resumebuilder.Parse(e)
	if result == "" {
		t.Error("result should not be empty")
	}
}
