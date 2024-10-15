# Snippets

![coverage](https://raw.githubusercontent.com/avnovoselov/snippets/badges/.badges/main/coverage.svg)

## ⬇️ Install
To install `snippets` use `go get`:
```shell
go get -u github.com/avnovoselov/snippets/issues
```

## ⌨️ Commands
### 🏗️ Git
#### Check if current directory is a git repository 
Run in checking directory
```shell
snippets git
```

#### Remove tracking branches no longer on remote
Run in project directory
```shell
snippets git remove-orphans
```
or shortest
```shell
snippets git ro
```

### 🐳 Docker
#### List of running containers
Run in any directory
```shell
snippets docker
```

#### Attach to running container
Run in any directory
```shell
snippets docker attach
```
and choose &darr;&uarr; container.