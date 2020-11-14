## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Installation
Install GO on your machine and reboot. Then:

Use the package manager to install LinkStatus.go
```bash
git clone https://github.com/MSTEWARDSON/LinkStatus.git
cd LinkStatus
go build
```

## Formatting the code
Before making a pull request, please run gofmt on all files you changed, or on the whole directory itself.

Single File
```go
gofmt -w file.go
```
OR

Directory
```go
go fmt ./
```

## Using the linter
For this project we are using the golint linter. If you do not have it already installed use:
```bash
go get -u golang.org/x/lint/golint
```

Single File
```bash
golint file.go
```
OR

Directory
```bash
golint ./
```
