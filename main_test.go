package main

import (
	"testing"
	"time"
)

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
