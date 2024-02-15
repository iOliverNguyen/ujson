package ujson

import "testing"

// https://github.com/iOliverNguyen/ujson/pull/5
func TestEmptyJson(t *testing.T) {
	input := []byte(`{

	}`)
	callbacks := []*MatchCallback{}
	opt := MatchOptions{
		IgnoreCase:       false,
		QuitIfNoCallback: false,
	}
	res := MatchResult{
		Count: 0,
	}
	err := Match(input, &opt, &res, callbacks...)
	if err != nil {
		t.Fatal(err)
	}
	// print count
	if res.Count != 0 {
		t.Fatal("expected count to be 0")
	}
}
