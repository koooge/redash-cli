package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type organizationsettingsClient struct {
	*redash.Client
}

const getOrganizationSettingsResBody = `{"something":"something"}`

func (c *organizationsettingsClient) GetOrganizationSettings(_ *redash.GetOrganizationSettingsInput) *redash.GetOrganizationSettingsOutput {
	return &redash.GetOrganizationSettingsOutput{StatusCode: 200, Body: getOrganizationSettingsResBody}
}

func TestNewCmdGetOrganizationSettings(t *testing.T) {
	testClient := &organizationsettingsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getOrganizationSettingsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetOrganizationSettings(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
