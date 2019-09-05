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

---

## Contents
- [Install or Download](#install-or-download)
  - [Mac OSX](#mac-osx)
  - [Download release from GitHub](#download-release-from-github)
- [Setup](#setup)
  - [1. Get `API_TOKEN`](#1-get-api_token)
  - [2. Make a `.env` config file](#2-make-a-env-config-file)
  - [3. Ready to go](#3-ready-to-go)
- [Configurations](#configurations)
  - [config by `.env` file](#config-by-env-file)
  - [config by Environment Variables](#config-by-environment-variables)
  - [Support `TOKEN_NAME`](#support-token_name)
  - [Support configuration source](#support-configuration-source)
- [Support Cloud(IaaS)](#support-cloudiaas)
- [Development and Build](#development-and-build)

## Install or Download
### Mac OSX
```
$ brew install jyny/tap/outliner
```

### Arch linux (wip)
```
# pacman -S outliner
```

### Ubuntu (wip)
```
# apt isntall outliner
```

### Download release from GitHub
Visit the [latest releases page](https://github.com/Jyny/outliner/releases/latest)

## Setup
#### 1. Get `API_TOKEN`
get `API_TOKEN` from cloud providers you want, like `Linode`, `DigitalOcean` and etc.

#### 2. Make a `.env` config file
write the `API_TOKEN` to `.env` file with [Support `TOKEN_NAME`](#support-token_name).
* `~/.outliner/.env`
```
TOKEN_NAME_1 = TOKEN_VALUE_1
TOKEN_NAME_2 = TOKEN_VALUE_2
...
```

#### 3. Ready to go
* install by package manager like `homeberw` `pacman` `apt` etc.
  1. `$ outliner`

* install by download
  1. open terminal, go directory whrere you download outliner
  2. maybe should add execute permission to binary (linux or mac)
  3. `$ ./outliner_{OS}`

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
#### 1. install package and build binary
```
$ make
```
#### 2. run
```
$ ./outliner
```
