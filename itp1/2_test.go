package itp1

import (
	"testing"
)

func TestSmallLargeEqual(t *testing.T) {
	samples := [][]string{{"1 2", "a < b"}, {"4 3", "a > b"}, {"5 5", "a == b"}}
	for _, s := range samples {
		res, err := SmallLargeEqual(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if res != s[1] {
			t.Fatalf("wrong result: %s", res)
		}
	}
}

func TestRange(t *testing.T) {
	samples := [][]string{{"1 3 8", "Yes"}, {"3 8 1", "No"}}
	for _, s := range samples {
		res, err := Range(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if res != s[1] {
			t.Fatalf("wrong result: %s", res)
		}
	}
}

func TestSortingThreeNumbers(t *testing.T) {
	s := []string{"3 8 1", "1 3 8"}
	res, err := SortingThreeNumbers(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}

func TestCircleInARectangle(t *testing.T) {
	samples := [][]string{{"5 4 2 2 1", "Yes"}, {"5 4 2 4 1", "No"}}
	for _, s := range samples {
		res, err := CircleInARectangle(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if res != s[1] {
			t.Fatalf("wrong result: %s", res)
		}
	}
}
