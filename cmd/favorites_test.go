package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type favoritesClient struct {
	*redash.Client
}

const getQueryFavoriteListResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *favoritesClient) GetQueryFavoriteList() *redash.GetQueryFavoriteListOutput {
	return &redash.GetQueryFavoriteListOutput{StatusCode: 200, Body: getQueryFavoriteListResBody}
}

func TestNewCmdGetQueryFavoriteList(t *testing.T) {
	testClient := &favoritesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQueryFavoriteListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQueryFavoriteList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDashboardFavoriteListResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *favoritesClient) GetDashboardFavoriteList() *redash.GetDashboardFavoriteListOutput {
	return &redash.GetDashboardFavoriteListOutput{StatusCode: 200, Body: getDashboardFavoriteListResBody}
}

func TestNewCmdGetDashboardFavoriteList(t *testing.T) {
	testClient := &favoritesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDashboardFavoriteListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDashboardFavoriteList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
