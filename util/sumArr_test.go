package util_test

import (
	"go-intro/util"
	"testing"
)

func TestSumArr(t *testing.T) {
	cases := []int{0}
	for _, c := range cases {
		got := util.SumArray()
		if got != c {
			t.Fail()
		}
	}
}
