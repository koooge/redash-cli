[![CircleCI](https://circleci.com/gh/koooge/redash-cli/tree/master.svg?style=svg)](https://circleci.com/gh/koooge/redash-cli/tree/master)
[![Build status](https://ci.appveyor.com/api/projects/status/l41n2c8ini09ppuv/branch/master?svg=true)](https://ci.appveyor.com/project/koooge/redash-cli/branch/master)
[![Maintainability](https://api.codeclimate.com/v1/badges/0cc6f5c87721c65e9f2a/maintainability)](https://codeclimate.com/github/koooge/redash-cli/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/0cc6f5c87721c65e9f2a/test_coverage)](https://codeclimate.com/github/koooge/redash-cli/test_coverage)
[![GoDoc](https://godoc.org/github.com/koooge/redash-cli?status.svg)](https://godoc.org/github.com/koooge/redash-cli)

# redash-cli
A single binary client tool for Redash.

## Install
Download from [Release page](https://github.com/koooge/redash-cli/releases)

```
$ wget -qO- https://github.com/koooge/redash-cli/releases/download/0.1.0/redash_$(uname -s)_amd64_0.1.0.tar.gz | tar zx
$ sudo mv redash /usr/local/bin/

$ redash version
0.1.0
```

## Usage
```
$ redash config [--profile <value>]
endpoint_url: <Redash's url>
api_key: <Your API key>

$ redash list-queries
```

[More about redash-cli](/doc/redash.md)
