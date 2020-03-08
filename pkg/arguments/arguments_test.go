package arguments

import (
	"errors"
	"testing"
)

func TestArgs(t *testing.T) {
	t.Run("Validate", testValidate)
}

func testValidate(t *testing.T) {
	cases := []struct {
		name   string
		url    string
		format string
		size   string
		want   bool
	}{
		{name: "testName", url: "https://www.google.com/", format: "png", size: "large", want: false},
		{name: "", url: "https://www.google.com/", want: true},
		{name: "testName", url: "", want: true},
	}

	for _, c := range cases {
		a := Args{
			Name:   c.name,
			URL:    c.url,
			Format: c.format,
			Size:   c.size,
		}
		got := a.Validate()
		isError := isCausedErrors(t, got)
		if isError != c.want {
			t.Errorf("invalid result.\n%s\ntest case: %#v \nerror expected: %#v \nactual: %#v ", got, c, c.want, isError)
		}
	}

}

func isCausedErrors(t *testing.T, err error) bool {
	t.Helper()
	return err != nil
}

const errorCase = "errorExpectded"

func TestParseFlags(t *testing.T) {
	cases := []struct {
		args      []string
		wantError bool
		want      Args
	}{
		{args: []string{"hoge.png", "https://www.google.com/"}, want: Args{Name: "hoge.png", URL: "https://www.google.com/"}},
		{args: []string{"", "https://www.google.com/"}, wantError: true},
		{args: []string{errorCase, "https://www.google.com/"}, wantError: true},
	}

	for _, c := range cases {
		flagSet = &FlagSetMock{_args: c.args}
		got, err := ParseFlags()

		if c.wantError && err == nil {
			t.Errorf("want error.\nParseFlags() = \n%#v", c)
			continue
		}

		if got.Name != c.want.Name {
			t.Errorf("got.Name = %s, want %s", got.Name, c.want.Name)
		}
		if got.URL != c.want.URL {
			t.Errorf("got.URL = %s, want %s", got.URL, c.want.URL)
		}
	}
}

type FlagSetMock struct {
	_args []string
	args  []string
}

func (f *FlagSetMock) Parse(arguments []string) error {
	// check error case.
	if f._args[0] == errorCase {
		err := errors.New("parse error")
		return err
	}

	f.args = f._args
	return nil
}

func (f *FlagSetMock) Arg(i int) string {
	if i >= len(f.args) {
		return ""
	}
	return f.args[i]
}
