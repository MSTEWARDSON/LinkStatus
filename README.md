# LinkStatus

LinkStatus reads a supplied file, finds, and checks links via a GET request. The program returns the status codes of the links and lists them for the user.
This program was developed using the GO programming language.

# Introduction
Hello to the person reading this. This is my release 0.1 project for the DPS909 class.

I have a blog going alongside the development of this project and future projects for this class which you can find [here](https://matthew-k-stewardson.blogspot.com/)

## Installation
Install GO on your machine and reboot. Then:

Use the package mananger to install LinkStatus.go
```bash
git clone https://github.com/MSTEWARDSON/LinkStatus.git
cd LinkStatus
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

