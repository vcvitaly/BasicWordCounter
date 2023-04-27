package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, map[FlagType]bool{CountLines: false, CountBytes: false})

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	exp := 3

	res := count(b, map[FlagType]bool{CountLines: true, CountBytes: false})

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1")

	exp := 5

	res := count(b, map[FlagType]bool{CountLines: false, CountBytes: true})

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestBuildFlagMap(t *testing.T) {
	countLines := true
	countBytes := true

	_, err := buildFlagMap(&countLines, &countBytes)
	if err == nil {
		t.Errorf("expected an error on receiving mutually exclusive flags")
	}
}
