package main

import (
	"testing"
	"time"
)

func TestArgs_Validate(t *testing.T) {
	cases := []struct {
		name string
		url  string
		want bool
	}{
		{name: "testName", url: "https://www.google.com/", want: false},
		{name: "", url: "https://www.google.com/", want: true},
		{name: "testName", url: "", want: true},
	}

	for _, c := range cases {
		a := Args{
			Name: c.name,
			URL:  c.url,
		}
		got := a.Validate()
		isError := isCausedErrors(t, got)
		if isError != c.want {
			t.Errorf("invalid result.\ntest case: %#v \nerror expected: %#v \nactual: %#v ", c, c.want, isError)
		}
	}

}

func TestMakeBasePath(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	name := "testfile"
	date := time.Date(2020, time.January, 1, 0, 0, 0, 0, jst)
	expect := "1577804400_testfile"

	got := MakeBasePath(name, date)
	if got != expect {
		t.Errorf("invalid value.\nexpect: %s\nactual: %s", expect, got)
	}
}

func isCausedErrors(t *testing.T, err error) bool {
	t.Helper()
	return err != nil
}
