# 🐤 Compose from anywhere

Small utility to run `docker-compose` commands on your projects without having to `cd` into them.

It is heavily inspired from [Captain](https://github.com/jenssegers/captain) but aims at being as transparent as possible by using the default `docker-compose` [CLI commands](https://docs.docker.com/compose/reference/) for all commands, so that all the commad you already use and know keep working out of the box.

```bash
# Replace
$ cd path/to/my/project1
$ docker-compose down

$ cd other/path/project2
$ docker-compose up -d

# By
$ cfa project1 down
$ cfa project2 up -d
```

**cfa** scan your directories looking for a `docker-compose.yml` file. Every folder meeting the condition are marked as candidate projects that you can operate on from anywhere.

Note that the project names are fuzzy searched. If several projects have a similar name, `cfa` will give you a list to chose from

## Installing

Using Compose from anywhere is easy. First, use `go get` to install the latest version of the library. This command will install the `cfa` executable along with the library and its dependencies:

```bash
go get -u github.com/riron/cfa
```

If you don't have Go installed on your machine you can download the executables.

### MacOSX

```
curl -L https://github.com/riron/cfa/releases/download/0.0.1/cfa-osx > /usr/local/bin/cfa && chmod +x /usr/local/bin/cfa
```

### Linux

```
curl -L https://github.com/riron/cfa/releases/download/0.0.1/cfa-linux > /usr/local/bin/cfa && chmod +x /usr/local/bin/cfa
```

### Windows

Download cfa.exe via https://github.com/riron/cfa/releases/download/0.0.1/cfa.exe and add it to your path

## Usage

```bash
$ cfa [my_project_name] [compose_command --compose-flags]
```

## Options

```bash
$ cfa
Manage your compose projects from anywhere

cfa allows you to use the same compose CLI you already know
but without the need to cd into your directories.
Just pass your project name as the first argument
and run your compose command on it.

Usage:
  cfa [flags] [project] [compose command]

Examples:
cfa my_project up -d
cfa -u=dev my_project exec my_container sh
cfa -f=my_pro
cfa -s

Flags:
  -D, --dev             Shorthand for -u=dev
  -f, --find string     List projects corresponding to search
  -h, --help            help for cfa
  -l, --list            List all available projects
  -s, --stop            Stop all running containers
  -u, --suffix string   Use a suffixed compose file (ex: -s=dev will use the docker-compose.dev.yml file)
```

## Config

Two environment variables allow you ton configure **cfa**

| ENV       | Description                                      | Default             |
| --------- | ------------------------------------------------ | ------------------- |
| CFA_ROOT  | Root folder from which to scan for compose files | User home directory |
| CFA_DEPTH | Maximum scan depth                               | 5                   |