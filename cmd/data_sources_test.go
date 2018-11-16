package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type datasourcesClient struct {
	*redash.Client
}

const getDataSourceListResBody = `[{"something":"something"}]`

func (c *datasourcesClient) GetDataSourceList() *redash.GetDataSourceListOutput {
	return &redash.GetDataSourceListOutput{StatusCode: 200, Body: getDataSourceListResBody}
}

func TestNewCmdGetDataSourceList(t *testing.T) {
	testClient := &datasourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDataSourceListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDataSourceList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDataSourceResBody = `{"something":"something"}`

func (c *datasourcesClient) GetDataSource(_ *redash.GetDataSourceInput) *redash.GetDataSourceOutput {
	return &redash.GetDataSourceOutput{StatusCode: 200, Body: getDataSourceResBody}
}

func TestNewCmdGetDataSource(t *testing.T) {
	testClient := &datasourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--datasource-id", "1"}, want: getDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDataSource(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDataSourceSchemaResBody = `{"something":"something"}`

func (c *datasourcesClient) GetDataSourceSchema(_ *redash.GetDataSourceSchemaInput) *redash.GetDataSourceSchemaOutput {
	return &redash.GetDataSourceSchemaOutput{StatusCode: 200, Body: getDataSourceSchemaResBody}
}

func TestNewCmdGetDataSourceSchema(t *testing.T) {
	testClient := &datasourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--datasource-id", "1"}, want: getDataSourceSchemaResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDataSourceSchema(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
