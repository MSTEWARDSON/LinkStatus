# LinkStatus

LinkStatus reads a supplied file, finds, and checks links via a GET request. The program returns the status codes of the links and lists them for the user.
This program was developed using the GO programming language.

## Installation

Use the package mananger to install LinkStatus.go
```bash
go install LinkStatus.go
```
## Usage

To recieve help information run the program with no arguments:
```bash
go run LinkStatus.go
```
To see the current version run:
```bash
go run LinkStatus.go -v or --version
```
To run the program with a file and check the links use:
```bash
go run LinkStatus.go filename
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)

