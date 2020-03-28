package camelcase_test

import (
	camel "github.com/ambye85/gophercises/camelcase"
	"testing"
)

func checkWordCount(actual int, expected int, t *testing.T) {
	if wordCount := actual; wordCount != expected {
		t.Fatalf("Expected wordCount to be %d, got %d", expected, wordCount)
	}
}

func TestSingleLetter(t *testing.T) {
	checkWordCount(camel.CountWords("a"), 1, t)
}

func TestSingleWord(t *testing.T) {
	checkWordCount(camel.CountWords("hello"), 1, t)
}

func TestTwoLetters(t *testing.T) {
	checkWordCount(camel.CountWords("aB"), 2, t)
}

func TestMultipleWords(t *testing.T) {
	checkWordCount(camel.CountWords("testingMultipleWords"), 3, t)
}
