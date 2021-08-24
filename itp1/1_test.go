package itp1

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
}

func TestXCubic(t *testing.T) {
	samples := [][]string{{"2", "8"}, {"3", "27"}}
	for _, s := range samples {
		result, err := XCubic(s[0])
		if err != nil {
			t.Fatalf("failed: %#v", err)
		}
		if result != s[1] {
			t.Fatalf("wrong result: %s", result)
		}
	}
}

func TestRectangle(t *testing.T) {
	sample := []string{"3 5", "15 16"}
	result, err := Rectangle(sample[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if result != sample[1] {
		t.Fatalf("wrong result: %s", result)
	}
}

func TestWatch(t *testing.T) {
	samples := []string{"46979", "13:2:59"}
	result, err := Watch(samples[0])
	if err != nil {
		t.Fatalf("failed: %#v", err)
	}
	if result != samples[1] {
		t.Fatalf("wrong result: %s", result)
	}
}
