package itp1

import (
	"testing"
)

func TestABProblem(t *testing.T) {
	s := []string{"3 2", "1 1 1.50000"}
	res, err := ABProblem(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}

func TestCircle(t *testing.T) {
	samples := [][]string{{"2", "12.566371 12.566371"}, {"3", "28.274334 18.849556"}}
	for _, s := range samples {
		res, err := Circle(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if res != s[1] {
			t.Fatalf("wrong result: %s", res)
		}
	}
}

func TestSimpleCalculator(t *testing.T) {
	s := []string{
		"1 + 2\n56 - 18\n13 * 2\n100 / 10\n27 + 81\n0 ? 0",
		"3\n38\n26\n10\n108",
	}
	res, err := SimpleCalculator(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}

func TestMinMaxAndSum(t *testing.T) {
	s := []string{
		"5\n10 1 5 4 17",
		"1 17 37",
	}
	res, err := MinMaxAndSum(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}
