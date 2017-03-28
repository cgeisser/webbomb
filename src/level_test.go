package main

import (
	"net/http"
	"testing"
)

func TestParseState(t *testing.T) {
	var tests = []struct {
		input     *http.Cookie
		wantState LevelState
		wantError bool
	}{
		{nil, LevelState(0), false},
		{&http.Cookie{Value: "7"}, LevelState(7), false},
		// errors
		{&http.Cookie{Value: "weasel"}, LevelState(0), true},
		{&http.Cookie{Value: ""}, LevelState(0), true},
	}
	for _, test := range tests {
		gotState, gotError := ParseStateFromCookie(test.input)
		if gotState != test.wantState ||
			(gotError == nil) == test.wantError {
			t.Errorf("ParseStateFromCookie(%v) = %v err: %v",
				*test.input,
				gotState, gotError)
		}
	}
}
