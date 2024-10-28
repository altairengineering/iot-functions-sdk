package variables

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const (
	variablesSecret = "function-variables"
)

var secretsPath = os.Getenv("_SECRETS_PATH")

func getSecretsPath() string {
	a := strings.TrimRight(secretsPath, "/") + "/" + variablesSecret
	a, _ = filepath.Abs(a)
	return a
}

type Variable struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

// Get returns the variable with the given name. It returns an empty string if
// the variable is not found.
func Get(variableName string) interface{} {
	variables, err := loadVariables()
	if err != nil {
		return ""
	}
	variable, ok := variables[variableName]
	if !ok {
		return ""
	}
	return variable.Value
}

// Exists returns whether the variable exists or not.
func Exists(variableName string) bool {
	variables, err := loadVariables()
	if err != nil {
		return false
	}
	_, ok := variables[variableName]
	return ok
}

func loadVariables() (map[string]Variable, error) {
	b, err := os.ReadFile(getSecretsPath())
	if err != nil {
		return nil, err
	}

	variables := make(map[string]Variable)
	err = json.Unmarshal(b, &variables)
	if err != nil {
		return variables, err
	}
	return variables, nil
}
