package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ParseInt(s string) int {
	k, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return k
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Index[T comparable](x T, s []T) (int, bool) {
	for i, t := range s {
		if x == t {
			return i, true
		}
	}
	return -1, false
}

func IndexKey[K comparable, T interface{}](key K, s []T, f func(T) K) (int, bool) {
	for i, t := range s {
		if key == f(t) {
			return i, true
		}
	}
	return -1, false
}

func Delete[T comparable](t T, s []T) ([]T, bool) {
	for i, t2 := range s {
		if t == t2 {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return s, false
}

func CheckErrorP(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadWholeFileForDay(day string, test bool) string {
	path := "./day" + day + "/input.txt"
	if test {
		path = "./day" + day + "/test.txt"
	}
	return ReadWholeFile(path)
}

func ReadWholeFile(name string) string {
	bytes, err := os.ReadFile(name)
	CheckErrorP(err)
	return string(bytes)
}

func ReadFileLinesForDay(day string, test bool) []string {
	path := "./day" + day + "/input.txt"
	if test {
		path = "./day" + day + "/test.txt"
	}
	return ReadFileLines(path)
}

func ReadFileLines(name string) []string {
	readFile, err := os.Open(name)
	defer readFile.Close()
	CheckErrorP(err)
	fileScanner := bufio.NewScanner(readFile)
	lines := make([]string, 0)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
