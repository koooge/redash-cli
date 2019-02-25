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

const listDataSourcesResBody = `[{"something":"something"}]`

func (c *dataSourcesClient) ListDataSources(_ *redash.ListDataSourcesInput) *redash.ListDataSourcesOutput {
	return &redash.ListDataSourcesOutput{StatusCode: 200, Body: listDataSourcesResBody}
}

func TestNewCmdListDataSources(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listDataSourcesResBody},
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

const createDataSourceResBody = `{"something":"something"}`

func (c *dataSourcesClient) CreateDataSource(_ *redash.CreateDataSourceInput) *redash.CreateDataSourceOutput {
	return &redash.CreateDataSourceOutput{StatusCode: 200, Body: createDataSourceResBody}
}

func TestNewCmdCreateDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--dbname", "something", "--name", "something", "--type", "pg"}, want: createDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdCreateDataSource(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const listDataSourcesTypesResBody = `[{"something":"something"}]`

func (c *dataSourcesClient) ListDataSourcesTypes(_ *redash.ListDataSourcesTypesInput) *redash.ListDataSourcesTypesOutput {
	return &redash.ListDataSourcesTypesOutput{StatusCode: 200, Body: listDataSourcesTypesResBody}
}

func TestNewCmdListDataSourcesTypes(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listDataSourcesTypesResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListDataSourcesTypes(testClient)
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

const updateDataSourceResBody = `{"something":"something"}`

func (c *dataSourcesClient) UpdateDataSource(_ *redash.UpdateDataSourceInput) *redash.UpdateDataSourceOutput {
	return &redash.UpdateDataSourceOutput{StatusCode: 200, Body: updateDataSourceResBody}
}

func TestNewCmdUpdateDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "123", "--dbname", "something", "--name", "something", "--type", "pg"}, want: updateDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdUpdateDataSource(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const deleteDataSourceResBody = `{"something":"something"}`

func (c *dataSourcesClient) DeleteDataSource(_ *redash.DeleteDataSourceInput) *redash.DeleteDataSourceOutput {
	return &redash.DeleteDataSourceOutput{StatusCode: 200, Body: deleteDataSourceResBody}
}

func TestNewCmdDeleteDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "123"}, want: deleteDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdDeleteDataSource(testClient)
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

const pauseDataSourceResBody = `{"something":"something"}`

func (c *dataSourcesClient) PauseDataSource(_ *redash.PauseDataSourceInput) *redash.PauseDataSourceOutput {
	return &redash.PauseDataSourceOutput{StatusCode: 200, Body: pauseDataSourceResBody}
}

func TestNewCmdPauseDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "123"}, want: pauseDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdPauseDataSource(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const unpauseDataSourceResBody = `{"something":"something"}`

func (c *dataSourcesClient) UnpauseDataSource(_ *redash.UnpauseDataSourceInput) *redash.UnpauseDataSourceOutput {
	return &redash.UnpauseDataSourceOutput{StatusCode: 200, Body: unpauseDataSourceResBody}
}

func TestNewCmdUnpauseDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "123"}, want: unpauseDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdUnpauseDataSource(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const testDataSourceResBody = `{"something":"something"}`

func (c *dataSourcesClient) TestDataSource(_ *redash.TestDataSourceInput) *redash.TestDataSourceOutput {
	return &redash.TestDataSourceOutput{StatusCode: 200, Body: testDataSourceResBody}
}

func TestNewCmdTestDataSource(t *testing.T) {
	testClient := &dataSourcesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "123"}, want: testDataSourceResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdTestDataSource(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
