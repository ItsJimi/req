# req
A simple cli to send http requests.


## Navigate
- [Why](#why)
- [Demo](#demo)
- [Installl](#install)
- [Usage](#usage)
  - [Config file](#config-file)
  - [List](#list)
  - [Run](#run)
- [Reference](#reference)
  - [Config file](#config-file-1)
  - [Commands](#commands)
    - [help](#help)
    - [version](#version)
    - [list](#list-1)
    - [run](#run-1)
  - [Rofi](#rofi)
- [Contribute](#contribute)
- [License](#license)

## Why
I created req mainly to use it with [Rofi](https://github.com/davatorium/rofi). At home, I can control domotics directly from Rofi.
> Check the [Rofi](#Rofi) section below.

## Demo
[![asciicast](https://asciinema.org/a/rpqbwVyj1f4KE2IapyG3r0QCW.svg)](https://asciinema.org/a/rpqbwVyj1f4KE2IapyG3r0QCW)

## Install
Use precompiled versions in [releases page](https://github.com/ItsJimi/req/releases)

or

```shell
go get -u github.com/itsjimi/req
```

## Usage
First of all, you need to create a `.req.json` in your home directory.
#### Config file
This is the most common `.req.json`.
```json
[
  {
    "name": "Get first post",
    "url": "https://jsonplaceholder.typicode.com/posts/1",
    "method": "GET"
  }
]
```
#### List
You can list all of available requests by running.
```shell
req list
```
In our case we have one element displayed.
```
Get first post
```
#### Run
You can run one or multiple requests by using run command.
```shell
req run "Get first post"
```
The result is displayed on terminal.
```json
{
  "userId": 1,
  "id": 1,
  "title": "...",
  "body": "..."
}
```

## Reference
### Config file
By default, [.req.json](https://github.com/ItsJimi/req/blob/master/.req.json) must be in your home directory.
```json
[
  {
    "name": "Get first post",
    "url": "https://jsonplaceholder.typicode.com/posts/1",
    "method": "GET",
    "output": "echo \"Title: {{.title}}\nBody: {{.body}}\""
  },
  {
    "name": "Create post",
    "url": "https://jsonplaceholder.typicode.com/posts",
    "method": "POST",
    "headers": [
      "Content-type: application/json; charset=UTF-8"
    ],
    "body": {
      "title": "foo",
      "body": "bar",
      "userId": 1
    }
  }
]
```
> `output` exec commands with bash and use go template to replace variables.

### Commands
- All commands can use `--help` or `-h` to display its specific help.
- All commands can use `--config` or `-c` to use a custom `.json` path. (By default req use a `.req.json` in your home directory)

#### help
Display helper
```shell
req help
```

#### version
Display version req using SemVer.
```shell
req version
```

#### list
Display list of `.req.json` requests names (each separated by `\n`).
```shell
req list
```

### run
Send one or multiple requests and display results.
> Note: You can use --silent or -s to display nothing.
```shell
req run <request name> [other request name] [...]
```

### Rofi
Req can be used in rofi with the [rofi-req.sh](https://github.com/ItsJimi/req/blob/master/rofi-req.sh) script.
![rofi-req](https://i.imgur.com/kjSOnSB.png)

## Contribute
Feel free to fork and make pull requests.


## License
[MIT](https://github.com/ItsJimi/req/blob/master/LICENSE)
