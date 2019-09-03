# Outliner
**Auto setup & deploy tool for outline VPN server**

[![asciicast](https://asciinema.org/a/265622.svg)](https://asciinema.org/a/265622)
```
Auto setup & deploy tool for outline VPN server

Usage:
  outliner [command]

Available Commands:
  create      create a Server
  deploy      deploy outliner to Server
  destroy     destroy a Server
  help        Help about any command
  inspect     inspect Server
  list        list following [command]
  version     show outliner version

Flags:
  -F, --file string   config file (default is $HOME/.outliner/.env)
  -h, --help          help for outliner

Use "outliner [command] --help" for more information about a command.
```

## Contents
- [Download and Install](#download-and-install)
- [Configuration](#configuration)
  - [Steps](#steps)
  - [Basic](#basic)
    - [config by `.env` file](#config-by-env-file)
    - [config by Environment Variables](#config-by-environment-variables)
  - [Other configuration source](#other-configuration-source)
  - [supported `TOKEN_NAME`](#supported-token_name)
- [Development and Build](#development-and-build)

## Download
download from latest [release](https://github.com/Jyny/outliner/releases/latest)

## Install

## Configuration
### Steps
1.  you will need `API-TOKEN` from cloud providers, like `Linode`, `DigitalOcean` and etc.

2. write `.env` config file, [refer to the following](config-by-env-file)

### Basic
following is two of config method in easiest way

#### config by `.env` file
* `~/.outliner/.env`
```
{TOKEN_NAME_1} = {TOKEN_VALUE_1}
{TOKEN_NAME_2} = {TOKEN_VALUE_2}
```

#### config by Environment Variables
```
$ {TOKEN_NAME_1}={TOKEN_VALUE_1} outliner [command]
```

### Other configuration source
Support the following configuration source (list by Precedence order)

1. with flag `-F, --file {FILE_PATH}`
2. Environment variables
3. `.env` file at `~/.outliner/`
4. `.env` file at `~/`
5. `.env` file at `./`

### supported `TOKEN_NAME`
find in `pkg/cloud/{ProviderNmae}/activator.go` 
```
var tokenNames = []string{
    "SUPPORTED_TOKEN_NAME_1",
    "SUPPORTED_TOKEN_NAME_2",
    ...
}
```

## Development and Build
1. install depend package
```
$ make mod
```
2. build binary
```
$ make build
```
3. run
```
$ ./build/outliner_$(go env GOOS)
```
