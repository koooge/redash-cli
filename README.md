[![CircleCI](https://circleci.com/gh/koooge/redash-cli/tree/master.svg?style=svg)](https://circleci.com/gh/koooge/redash-cli/tree/master)

# redash-cli
A single binary client tool for Redash.

## Install
Download from [Release page](https://github.com/koooge/redash-cli/releases)

```
$ wget -qO- https://github.com/koooge/redash-cli/releases/download/0.0.2/redash_$(uname -s)_amd64_0.0.2.tar.gz | tar zx
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
