package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

type DummyLevel struct {
	valid    error
	complete bool
	render   string
}

func (dl DummyLevel) ValidRequest(r *http.Request, ls *LevelState) error {
	return dl.valid
}

func (dl DummyLevel) IsComplete(r *http.Request,
	ls *LevelState) (bool, error) {
	return dl.complete, nil
}

func (dl DummyLevel) Render(w http.ResponseWriter, r *http.Request,
	ls *LevelState) {
	fmt.Fprint(w, dl.render)
}

type ResponseRecorder struct {
	response bytes.Buffer
	header   http.Header
}

func (rr *ResponseRecorder) Header() http.Header {
	return rr.header
}

func (rr *ResponseRecorder) Write(bytes []byte) (int, error) {
	return rr.response.Write(bytes)
}

func (rr *ResponseRecorder) WriteHeader(int) {}

func (rr ResponseRecorder) String() string {
	return rr.response.String()
}

func TestHandleLevel(t *testing.T) {
	var tests = []struct {
		input    DummyLevel
		expected string
	}{
		{DummyLevel{render: "hi"}, "hi"},
	}
	for _, test := range tests {
		var rr ResponseRecorder
		handleLevel(&rr, nil, test.input, LevelState(0))
		if test.expected != rr.String() {
			t.Errorf("handleLevel(%v) = %v, expected %v",
				test.input, rr, test.expected)
		}
	}

}
