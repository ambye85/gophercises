package caesar_test

import (
	"github.com/ambye85/gophercises/caesar"
	"testing"
)

func checkEncryption(unencrypted string, rotations int, encrypted string, t *testing.T) {
	if actual := caesar.Encrypt(unencrypted, rotations); actual != encrypted {
		t.Fatalf("Got %s, expected %s", actual, encrypted)
	}
}

func TestLowerARot0IsLowerA(t *testing.T) {
	checkEncryption("a", 0, "a", t)
}

func TestUpperARot0IsUpperA(t *testing.T) {
	checkEncryption("A", 0, "A", t)
}

func TestLowerARot1IsLowerB(t *testing.T) {
	checkEncryption("a", 1, "b", t)
}

func TestUpperARot1IsUpperB(t *testing.T) {
	checkEncryption("A", 1, "B", t)
}

func TestLowerZRot1IsLowerA(t *testing.T) {
	checkEncryption("z", 1, "a", t)
}

func TestUpperZRot1IsUpperA(t *testing.T) {
	checkEncryption("Z", 1, "A", t)
}

func TestLowerARot26IsLowerA(t *testing.T) {
	checkEncryption("a", 26, "a", t)
}

func TestUpperARot26IsUpperA(t *testing.T) {
	checkEncryption("A", 26, "A", t)
}

func TestWordRot2(t *testing.T) {
	checkEncryption("middle-Outz", 2, "okffng-Qwvb", t)
}
