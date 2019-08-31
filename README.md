# Outliner
Auto setup & deploy tool for outline VPN server

## Install

## Usage
```
Auto setup & deploy tool for outline VPN server

Usage:
  outliner [command]

Available Commands:
  create      create a server
  destroy     destroy a server
  help        Help about any command
  inspect     inspect server
  list        list following [command]

Flags:
  -F, --file string   config file (default is $HOME/.outliner/.env)
  -h, --help          help for outliner

Use "outliner [command] --help" for more information about a command.
```

## Configuration
### Example
#### 1. by config file (`.env` by default)
* `~/.outliner/.env`
```
{TOKEN_NAME_1} = {TOKEN_VALUE_1}
{TOKEN_NAME_2} = {TOKEN_VALUE_2}
```

#### 2. by Environment Variables
```
$ {TOKEN_NAME_1}={TOKEN_VALUE_1} outliner [command]
```

### support following config methods (by Precedence order)
1. with flag `-F, --file {FILE_PATH}`
2. Environment variables
3. `.env` file at `~/.outliner/`
4. `.env` file at `~/`
5. `.env` file at `./`

### supported `{TOKEN_NAME}`
find in `pkg/cloud/{ProviderNmae}/activator.go` 
```
var tokenNames = []string{
    "SUPPORTED_TOKEN_NAME_1",
    "SUPPORTED_TOKEN_NAME_2",
    ...
}
```

## Development & Build
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
$ ./outliner
```