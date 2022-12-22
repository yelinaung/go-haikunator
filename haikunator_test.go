package haikunator

import (
	"testing"
	"time"
)

func TestAreGeneratedNamesRandomOrNot(t *testing.T) {
	names := make([]string, 50)

	for i, _ := range names {
		haikunator := New(time.Now().UTC().UnixNano())
		names[i] = haikunator.Haikunate()
	}

	for i, name1 := range names {
		for j, name2 := range names {
			if i != j && name1 == name2 {
				t.Fatalf("not unique: %v : %v and %v :%v", i, j, name1, name2)
			}
		}
	}
}

func TestTotalCombinations(t *testing.T) {
	// This test should be kept up-to-date with the README
	expected := 8645
	found := len(ADJECTIVES) * len(NOUNS)
	if expected != found {
		t.Fatalf("mismatched total combinations: expected %v but found %v", expected, found)
	}
}
