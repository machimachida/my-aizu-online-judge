package itp1

import (
	"testing"
)

func TestTogglingCases(t *testing.T) {
	s := []string{
		"fAIR, LATER, OCCASIONALLY CLOUDY.",
		"Fair, later, occasionally cloudy.",
	}
	res, err := TogglingCases(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestSumOfNumbers(t *testing.T) {
	s := []string{
		"123\n55\n1000\n0",
		"6\n10\n1",
	}
	res, err := SumOfNumbers(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestCountingCharacters(t *testing.T) {
	s := []string{
		"This is a pen.",
		"a : 1\nb : 0\nc : 0\nd : 0\ne : 1\nf : 0\ng : 0\nh : 1\ni : 2\nj : 0\nk : 0\nl : 0\nm : 0\nn : 1\no : 0\np : 1\nq : 0\nr : 0\ns : 2\nt : 1\nu : 0\nv : 0\nw : 0\nx : 0\ny : 0\nz : 0",
	}
	res, err := CountingCharacters(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestRing(t *testing.T) {
	samples := [][]string{
		{"vanceknowledgetoad\nadvance", "Yes"},
		{"vanceknowledgetoad\nadvanced", "No"},
	}
	for _, s := range samples {
		res, err := Ring(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if res != s[1] {
			t.Fatalf("ans: %s\nresult: %s", s[1], res)
		}
	}
}
