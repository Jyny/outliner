# Outliner
**CLI tool for Auto setup and deploy outline VPN**

[![asciicast](https://asciinema.org/a/265622.svg)](https://asciinema.org/a/265622)
```
$ outliner --help
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
- [Download](#download)
- [Setup](#setup)
  - [1. Get `API_TOKEN`](#1-get-api_token)
  - [2. Make a `.env` config file](#2-make-a-env-config-file)
  - [3. Generate ssh key](#3-generate-ssh-key)
  - [4. Ready to go](#4-ready-to-go)
- [Usage](#usage)
- [Configurations](#configurations)
  - [config by `.env` file](#config-by-env-file)
  - [config by Environment Variables](#config-by-environment-variables)
  - [Support `TOKEN_NAME`](#support-token_name)
  - [Support configuration source](#support-configuration-source)
- [Support Cloud(IaaS)](#support-cloudiaas)
- [Development and Build](#development-and-build)

## Download
download from latest [release](https://github.com/Jyny/outliner/releases/latest)

## Setup
##### 1. Get `API_TOKEN`
get `API_TOKEN` from cloud providers you want, like `Linode`, `DigitalOcean` and etc.

##### 2. Make a `.env` config file
write the `API_TOKEN` to `.env` file whith `TOKEN_NAME` [like this](#config-by-env-file).
Reference the  [Support `TOKEN_NAME`](#supported-token_name) below.

##### 3. Generate SSH key
outliner support auto generate ssh key, if Not found `id_rsa` and `id_rsa.pub` in `$HOME/.ssh/`.
[ssh library package](https://godoc.org/golang.org/x/crypto/ssh) outliner use, is implement by golang.org
run any command in outliner will do this after asking.
```
$ outliner help
Continue to Generate New ssh key? (y/n) [y]:
```
To generate key by yourself, see [Generating a new SSH key and adding it to the ssh-agent](https://help.github.com/en/enterprise/2.16/user/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

##### 4. Ready to go
* open terminal, go directory whrere you download outliner
* maybe should add execute permission to binary (linux or mac)
* `$ ./build/outliner_{OS}`

## Configurations
### config by `.env` file
* `~/.outliner/.env`
```
TOKEN_NAME_1 = TOKEN_VALUE_1
TOKEN_NAME_2 = TOKEN_VALUE_2
...
```

### config by Environment Variables
```
$ {TOKEN_NAME_1}={TOKEN_VALUE_1} outliner [command]
```

### Support `TOKEN_NAME`
| Provider     | TOKEN_NAME         | 
| -------------|--------------------|
| Linode       | `LINODE_TOKEN`     |
| Linode       | `LINODE_CLI_TOKEN` |
| Linode       | `LINODE_API_TOKEN` |
| Digitalocean |                    |
| Vultr        |                    |

find full list in `pkg/cloud/{ProviderNmae}/activator.go` as below
```
var tokenNames = []string{
    "SUPPORTED_TOKEN_NAME_1",
    "SUPPORTED_TOKEN_NAME_2",
    ...
}
```

### Support configuration source
outliner Support the following configuration source (list by Precedence order)

1. with flag `-F, --file {FILE_PATH}`
2. Environment variables
3. `.env` file at `~/.outliner/`
4. `.env` file at `~/`
5. `.env` file at `./`

## Support Cloud(IaaS)
* Linode
* DigitalOcean (WIP)
* Vultr (WIP)
...

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
