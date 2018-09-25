package e2e

import (
	"os"
	"strings"
	"testing"
)

type Env struct {
	Namespace string
}

func BuildEnv(t *testing.T) Env {
	env := Env{
		Namespace: os.Getenv("KWT_E2E_NAMESPACE"),
	}
	env.Validate(t)
	return env
}

func (e Env) Validate(t *testing.T) {
	errStrs := []string{}

	if len(e.Namespace) == 0 {
		errStrs = append(errStrs, "Expected Namespace to be non-empty")
	}

	if len(errStrs) > 0 {
		t.Fatalf("%s", strings.Join(errStrs, "\n"))
	}
}
