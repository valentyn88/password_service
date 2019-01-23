package password_service

import (
	"reflect"
	"testing"
)

func TestRequestParam_DefVal(t *testing.T) {
	expectedRP := RequestParam{LettersLen: 0, SpecChLen: 0, NumbersLen: 0, Count: 1}

	rp := RequestParam{LettersLen: -1, SpecChLen: -2, NumbersLen: -3, Count: -5}
	rp.DefVal()

	if !reflect.DeepEqual(expectedRP, rp) {
		t.Errorf("expected %v and got %v are unequal", expectedRP, rp)
	}
}

func TestIsValidReqParamsLen(t *testing.T) {
	testCases := []struct {
		name     string
		rp       RequestParam
		expected bool
	}{
		{name: "total length less than 5", rp: RequestParam{LettersLen: 0, SpecChLen: 2, NumbersLen: 2}, expected: false},
		{name: "total length more than 15", rp: RequestParam{LettersLen: 10, SpecChLen: 2, NumbersLen: 4}, expected: false},
		{name: "total length is 10", rp: RequestParam{LettersLen: 3, SpecChLen: 4, NumbersLen: 3}, expected: true},
	}

	for _, tc := range testCases {
		got := IsValidReqParamsLen(tc.rp)
		if tc.expected != got {
			t.Errorf("expected %v and got %v are unequal", tc.expected, got)
		}
	}
}

func TestIsValidReqParamsCount(t *testing.T) {
	testCases := []struct {
		name     string
		rp       RequestParam
		expected bool
	}{
		{name: "count less than 50", rp: RequestParam{Count: 20}, expected: true},
		{name: "count more than 50", rp: RequestParam{Count: 51}, expected: false},
	}

	for _, tc := range testCases {
		got := IsValidReqParamsCount(tc.rp)
		if tc.expected != got {
			t.Errorf("expected %v and got %v are unequal", tc.expected, got)
		}
	}
}
