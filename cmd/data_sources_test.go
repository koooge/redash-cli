package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type dataSourcesClient struct {
	*redash.Client
}

const getDataSourceListResBody = `[{"something":"something"}]`

func (c *dataSourcesClient) ListDataSources(_ *redash.ListDataSourcesInput) *redash.ListDataSourcesOutput {
	return &redash.ListDataSourcesOutput{StatusCode: 200, Body: getDataSourceListResBody}
}

func TestNewCmdListDataSources(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDataSourceListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListDataSources(testClient)
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

func (c *dataSourcesClient) GetDataSource(_ *redash.GetDataSourceInput) *redash.GetDataSourceOutput {
	return &redash.GetDataSourceOutput{StatusCode: 200, Body: getDataSourceResBody}
}

func TestNewCmdGetDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "1"}, want: getDataSourceResBody},
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

func (c *dataSourcesClient) GetDataSourceSchema(_ *redash.GetDataSourceSchemaInput) *redash.GetDataSourceSchemaOutput {
	return &redash.GetDataSourceSchemaOutput{StatusCode: 200, Body: getDataSourceSchemaResBody}
}

func TestNewCmdGetDataSourceSchema(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "1"}, want: getDataSourceSchemaResBody},
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
