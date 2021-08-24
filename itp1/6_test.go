package itp1

import (
	"testing"
)

func TestReversingNumbers(t *testing.T) {
	samples := [][]string{
		{"5\n1 2 3 4 5", "5 4 3 2 1"},
		{"8\n3 3 4 4 5 8 7 9", "9 7 8 5 4 4 3 3"},
	}
	for _, s := range samples {
		res, err := ReversingNumbers(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if res != s[1] {
			t.Fatalf("ans: %s\nresult: %s", s[1], res)
		}
	}
}

func TestFindingMissingCards(t *testing.T) {
	s := []string{
		"47\nS 10\nS 11\nS 12\nS 13\nH 1\nH 2\nS 6\nS 7\nS 8\nS 9\nH 6\nH 8\nH 9\nH 10\nH 11\nH 4\nH 5\nS 2\nS 3\nS 4\nS 5\nH 12\nH 13\nC 1\nC 2\nD 1\nD 2\nD 3\nD 4\nD 5\nD 6\nD 7\nC 3\nC 4\nC 5\nC 6\nC 7\nC 8\nC 9\nC 10\nC 11\nC 13\nD 9\nD 10\nD 11\nD 12\nD 13",
		"S 1\nH 3\nH 7\nC 12\nD 8\n",
	}
	res, err := FindingMissingCards(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestOfficialHouse(t *testing.T) {
	s := []string{
		"3\n1 1 3 8\n3 2 2 7\n4 3 8 1",
		" 0 0 8 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 0 0 0\n####################\n 0 0 0 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 0 0 0\n####################\n 0 0 0 0 0 0 0 0 0 0\n 0 7 0 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 0 0 0\n####################\n 0 0 0 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 0 0 0\n 0 0 0 0 0 0 0 1 0 0",
	}
	res, err := OfficialHouse(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestMatrixVectorMultiplication(t *testing.T) {
	s := []string{
		"3 4\n1 2 0 1\n0 3 0 1\n4 1 1 0\n1\n2\n3\n0",
		"5\n6\n9",
	}
	res, err := MatrixVectorMultiplication(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}
