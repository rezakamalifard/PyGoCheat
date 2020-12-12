package list_test

import (
	"github.com/rezakamalifard/PyGoCheat/list"
	"testing"
)

func TestReverse(t *testing.T) {
	got := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list.Reverse(got)
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want []int) {
	if (got == nil) != (want == nil) {
		t.Fatalf("want %q got %q", want, got)
	}

	if len(got) != len(want) {
		t.Fatalf("want %q got %q", want, got)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("want %q got %q", want, got)
		}
	}
}
