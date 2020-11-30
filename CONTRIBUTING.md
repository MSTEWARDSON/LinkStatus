## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Installation
Install GO on your machine and reboot. Then:

Use the package manager to install LinkStatus.go
```go
git clone https://github.com/MSTEWARDSON/LinkStatus.git
cd LinkStatus
go install
```

If you are making changes to the code, before you test run:
```go
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

### Testing
If you are making a pull request, make sure to run:
```bash
go test
```
This will run your code through some checks to make sure it doesn't break the program. If you want to add your own tests just make a new file with a clear name to show its a test. For example "A_test.go".

### Code Coverage Check
To check the code coverage run:
```bash
go test [filename] -cover
```
