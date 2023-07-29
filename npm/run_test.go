package npm

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	err := Run("build", "../testdir")
	if err != nil && !errors.Is(err, RunError) {
		t.Error(err.Error())
	}
}
