package itp1

import (
	"testing"
)

func TestPrintManyHelloWorld(t *testing.T) {
}

func TestPrintTestCases(t *testing.T) {
	s := []string{
		"3\n5\n11\n7\n8\n19\n0",
		"Case 1: 3\nCase 2: 5\nCase 3: 11\nCase 4: 7\nCase 5: 8\nCase 6: 19",
	}
	res, err := PrintTestCases(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}

func TestSwappingTwoNumbers(t *testing.T) {
	s := []string{
		"3 2\n2 2\n5 3\n0 0",
		"2 3\n2 2\n3 5",
	}
	res, err := SwappingTwoNumbers(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}

func TestHowManyDivisors(t *testing.T) {
	s := []string{"5 14 80", "3"}
	res, err := HowManyDivisors(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}
