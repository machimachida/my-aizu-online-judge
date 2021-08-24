package itp1

import (
	"testing"
)

func TestGrading(t *testing.T) {
	s := []string{
		"40 42 -1\n20 30 -1\n0 2 -1\n-1 -1 -1",
		"A\nC\nF",
	}
	res, err := Grading(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestHowManyWays(t *testing.T) {
	s := []string{"5 9\n0 0", "2"}
	res, err := HowManyWays(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestSpreadsheet(t *testing.T) {
	s := []string{
		"4 5\n1 1 3 4 5\n2 2 2 4 5\n3 3 0 1 1\n2 3 4 4 6",
		"1 1 3 4 5 14\n2 2 2 4 5 15\n3 3 0 1 1 8\n2 3 4 4 6 19\n8 9 9 13 17 56",
	}
	res, err := Spreadsheet(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestMatrixMultiplication(t *testing.T) {
	s := []string{
		"3 2 3\n1 2\n0 3\n4 5\n1 2 1\n0 3 2",
		"1 8 5\n0 9 6\n4 23 14",
	}
	res, err := MatrixMultiplication(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}
