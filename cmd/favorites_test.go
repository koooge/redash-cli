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

const listQueryFavoritesResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *favoritesClient) ListQueryFavorites(_ *redash.ListQueryFavoritesInput) *redash.ListQueryFavoritesOutput {
	return &redash.ListQueryFavoritesOutput{StatusCode: 200, Body: listQueryFavoritesResBody}
}

func TestNewCmdListQueryFavorites(t *testing.T) {
	testClient := &favoritesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listQueryFavoritesResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListQueryFavorites(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const listDashboardFavoritesResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *favoritesClient) ListDashboardFavorites(_ *redash.ListDashboardFavoritesInput) *redash.ListDashboardFavoritesOutput {
	return &redash.ListDashboardFavoritesOutput{StatusCode: 200, Body: listDashboardFavoritesResBody}
}

func TestNewCmdListDashboardFavorites(t *testing.T) {
	testClient := &favoritesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listDashboardFavoritesResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListDashboardFavorites(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
