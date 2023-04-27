package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	flags, err := buildFlagMap(lines, bytes)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println(count(os.Stdin, flags))
}

func count(r io.Reader, flags map[FlagType]bool) int {
	scanner := bufio.NewScanner(r)

	if countLines, ok := flags[CountLines]; ok && countLines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes, ok := flags[CountBytes]; ok && countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func buildFlagMap(lines *bool, bytes *bool) (map[FlagType]bool, error) {
	if *lines && *bytes {
		return nil, errors.New("only one of l/b options is supported")
	}

	return map[FlagType]bool{
		CountLines: *lines,
		CountBytes: *bytes,
	}, nil
}

type FlagType int

const (
	CountLines FlagType = iota
	CountBytes
)
