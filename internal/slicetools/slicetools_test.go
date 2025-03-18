package slicetools_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/internal/slicetools"
)

func TestEvery(t *testing.T) {
	testdata := []int{1, 2, 3, 4, 5}

	isLessThanTen := func(_ int, n int) bool {
		return n < 10
	}

	isEven := func(_ int, n int) bool {
		return n%2 == 0
	}

	if slicetools.Every(testdata, isEven) {
		t.Errorf("all numbers are even")
	}

	if !slicetools.Every(testdata, isLessThanTen) {
		t.Errorf("not all numbers are less than 10")
	}
}

func TestSome(t *testing.T) {
	testdata := []int{1, 2, 3, 4, 5}

	isLessThanZero := func(_ int, n int) bool {
		return n < 0
	}

	isEven := func(_ int, n int) bool {
		return n%2 == 0
	}

	if !slicetools.Some(testdata, isEven) {
		t.Errorf("all numbers are not even")
	}

	if slicetools.Some(testdata, isLessThanZero) {
		t.Errorf("some numbers are less than 0")
	}
}
