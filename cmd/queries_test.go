package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewCmdGetQueryList(t *testing.T) {
	const getQueryListResBody = `{"something":"something","results":[{"something":"something"}]}`

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQueryListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQueryList()
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
