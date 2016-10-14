package main

import "testing"

const (
	pathInvalid = "somecrazypath"
)

func TestReadConfig(t *testing.T) {
	if s, err := readConfig(pathInvalid, "user1", 22); err == nil {
		t.Error("Reading of file '", pathInvalid, "' must return error. returned list:", s)
	}
}
