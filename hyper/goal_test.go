package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestGoal_Reached(t *testing.T) {
	goal := hyper.Goal{hyper.Red, hyper.Point{10, 5}}

	testcases := []struct {
		Name string
		*hyper.Actor
		Expected bool
	}{
		{"Reached", &hyper.Actor{hyper.Red, hyper.Point{10, 5}}, true},
		{"Different color", &hyper.Actor{hyper.Blue, hyper.Point{10, 5}}, false},
		{"Different row", &hyper.Actor{hyper.Red, hyper.Point{10, 4}}, false},
		{"Different column", &hyper.Actor{hyper.Red, hyper.Point{9, 5}}, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			if goal.Reached(testcase.Actor) != testcase.Expected {
				t.Errorf("unexpected value in case %s", testcase.Name)
			}
		})
	}
}
