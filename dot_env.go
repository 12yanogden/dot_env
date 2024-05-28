package dot_env

import (
	"os/exec"
	"strconv"
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
	bytes, err := exec.Command("grep", "-c", key, file).Output()

	if err != nil {
		return err
	}

	keyCount := int(bytes[0] - '0')

	if keyCount == 0 {
		panic("no instances of " + key + " in " + file)
	} else if keyCount > 1 {
		panic("only 1 instance of " + key + " allowed in " + file + ". Found " + strconv.Itoa(int(keyCount)))
	} else {
		_, err := exec.Command("sed", "-i", "'s%"+key+" *= *[a-zA-Z0-9]*%"+key+" = "+value+"%'", file).Output()

		if err != nil {
			return err
		}
	}

	return nil
}
