package dot_env

import (
	"reflect"
	"testing"
)

func TestCatError(t *testing.T) {
	does_not_exist_file := "does_not_exist.txt"
	expected := false
	_, err := Get(does_not_exist_file, "SOME_VAR")
	actual := err == nil

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestParse(t *testing.T) {
	dot_env_file := ".env"
	expected := map[string]string{
		"KEY1": "VALUE1",
		"KEY2": "VALUE2",
		"KEY3": "VALUE3",
	}
	actual, _ := parse(dot_env_file)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestGet(t *testing.T) {
	dot_env_file := ".env"
	expected := "VALUE2"
	actual, _ := Get(dot_env_file, "KEY2")

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestSet(t *testing.T) {
	dot_env_file := ".env"
	expected := "CHANGED"
	actual := ""

	Set(dot_env_file, "KEY2", "CHANGED")
	actual, _ = Get(dot_env_file, "KEY2")

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
