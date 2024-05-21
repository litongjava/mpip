# mpip
`mpip` is a command line tool written in Go for installing and uninstalling Python packages, and automatically updating the `requirements.txt` file.

## Features

- Install Python packages and add them to `requirements.txt`
- Uninstall Python packages and remove them from `requirements.txt`

## Installation

First, clone this repository:

```sh
git clone https://github.com/litongjava/mpip.git
cd mpip
```

Then, compile the `mpip` tool:

```sh
go build -o mpip main.go
```
This will generate an executable named `mpip` in the current directory. Please add `mpip` to your path:

```
go install
```
It will be automatically added to the go-path.

## Usage

### Installing Packages

Use the `mpip install <package>` command to install a Python package and add the package name to the `requirements.txt` file:

```sh
mpip install <package>
```

For example, to install the `requests` package:

```sh
mpip install requests
```

### Uninstalling Packages

Use the `mpip uninstall <package>` command to uninstall a Python package and remove the package name from the `requirements.txt` file:

```sh
mpip uninstall <package>
```

For example, to uninstall the `requests` package:

```sh
mpip uninstall requests
```

## Notes

- Ensure that `pip` is installed in your environment and accessible from the command line.
- The `requirements.txt` file will be automatically created or updated in the current directory.

## Contribution

If you have any suggestions for improvements or find any issues, please submit a pull request or an issue.

## License

This project uses the MIT license. For more information, see the [LICENSE](LICENSE) file.