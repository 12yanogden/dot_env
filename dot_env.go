package dot_env

import (
	"strings"

	"github.com/12yanogden/cat"
)

func parse(file string) (map[string]string, error) {
	contents, err := cat.Cat(file)
	variables := map[string]string{}

	if err != nil {
		return variables, err
	}

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		pair := strings.Split(
			strings.ReplaceAll(line, " ", ""),
			"=",
		)
		variables[pair[0]] = pair[1]
	}

	return variables, nil
}

func Get(file string, key string) (string, error) {
	variables, err := parse(file)

	if err != nil {
		return "", err
	}

	return variables[key], nil
}

func Set(file string, key string, value string) error {
	variables, err := parse(file)

	if err != nil {
		return err
	}

	variables[key] = value

	return nil
}
