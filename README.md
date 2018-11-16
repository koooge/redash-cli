[![CircleCI](https://circleci.com/gh/koooge/redash-cli/tree/master.svg?style=svg)](https://circleci.com/gh/koooge/redash-cli/tree/master)

# redash-cli
A single binary client tool for Redash.

## Install
Download from [Release page](https://github.com/koooge/redash-cli/releases)

```
$ wget -qO- https://github.com/koooge/redash-cli/releases/download/0.0.2/redash_$(uname -s)_amd64_0.0.2.tar.gz | tar zx
$ sudo mv redash /usr/local/bin/

$ redash version
0.0.2
```

## Usage
```
$ redash config [--profile <value>]
endpoint_url: <Redash's url>
api_key: <Your API key>

$ redash get-querylist
```

more info

```
$ redash -h
redash

Usage:
  redash [flags]
  redash [command]

Available Commands:
  config                   config
  delete-query             delete-query
  get-alert                get-alert
  get-alertlist            get-alertlist
  get-datasource           get-datasource
  get-datasourcelist       get-datasourcelist
  get-datasourceschema     get-datasourceschema
  get-destination          get-destination
  get-destinationlist      get-destinationlist
  get-destinationtypelist  get-destinationtypelist
  get-events               get-events
  get-group                get-group
  get-grouplist            get-grouplist
  get-job                  get-job
  get-myqueries            get-myqueries
  get-organizationsettings get-organizationsettings
  get-query                get-query
  get-querylist            get-querylist
  get-queryrecent          get-queryrecent
  get-queryresult          get-queryresult
  get-querysearch          get-querysearch
  get-querysnippet         get-querysnippet
  get-querysnippetlist     get-querysnippetlist
  get-querytags            get-querytags
  get-user                 get-user
  get-userlist             get-userlist
  help                     Help about any command
  post-query               post-query
  post-querylist           post-querylist
  version                  version

Flags:
      --apikey string         api key
      --config string         config file (default "$HOME/.config/redash-cli/config.yaml")
      --endpoint-url string   endpoint url
  -h, --help                  help for redash
      --profile string        profile (default "default")

Use "redash [command] --help" for more information about a command.
```
