# Fano (Fast node_modules)

[![Go Report Card](https://goreportcard.com/badge/github.com/ductnn/fano)](https://goreportcard.com/report/github.com/ductnn/fano) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A simple tool that help you remove unneeded files from `node_modules` folder.

## Installation

From source:

```bash
➜  git clone https://github.com/ductnn/fano.git
➜  cd fano
➜  fano git:(master) ✗ go build -o fano main.go
➜  fano git:(master) ✗ ls
README.md go.mod    go.sum    internal  fano      main.go
```

Then, you can copy file `fano` to project nodejs and run:

```bash
➜  nodejs-boilerplate git:(master) ✗ ./fano node_modules
Files total:   78,206
Files removed: 20,069
Size removed:  55 MB
```

or check [Makefile](https://github.com/ductnn/fano/blob/master/Makefile) for building bin on your local.

## Contribution
All contributions are welcomed in this project.

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
