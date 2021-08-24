package itp1

import (
	"testing"
)

func TestPrintARectangle(t *testing.T) {
	s := []string{
		"3 4\n5 6\n2 2\n0 0",
		"####\n####\n####\n\n######\n######\n######\n######\n######\n\n##\n##\n\n",
	}
	res, err := PrintARectangle(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("wrong result: %s", res)
	}
}

func TestPrintAFrame(t *testing.T) {
	s := []string{
		"3 4\n5 6\n3 3\n0 0",
		"####\n#..#\n####\n\n######\n#....#\n#....#\n#....#\n######\n\n###\n#.#\n###\n\n",
	}
	res, err := PrintAFrame(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestPrintAChessboard(t *testing.T) {
	s := []string{
		"3 4\n5 6\n3 3\n2 2\n1 1\n0 0",
		"#.#.\n.#.#\n#.#.\n\n#.#.#.\n.#.#.#\n#.#.#.\n.#.#.#\n#.#.#.\n\n#.#\n.#.\n#.#\n\n#.\n.#\n\n#\n\n",
	}
	res, err := PrintAChessboard(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}

func TestStructuredProgramming(t *testing.T) {
	s := []string{"30", " 3 6 9 12 13 15 18 21 23 24 27 30"}
	res, err := StructuredProgramming(s[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if res != s[1] {
		t.Fatalf("ans: %s\nresult: %s", s[1], res)
	}
}
