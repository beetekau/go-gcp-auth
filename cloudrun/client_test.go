package RUN

import (
	"os"
	"testing"
)

func TestAbc(t *testing.T) {
	var data interface{}
	runURL := os.Getenv("runURL")
	err := Get(runURL, data)
	if err != nil {
		t.Error(err) // to indicate test failed
	}

}
