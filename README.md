# Fano (Fast node_modules)

[![Go Report Card](https://goreportcard.com/badge/github.com/ductnn/fano)](https://goreportcard.com/report/github.com/ductnn/fano) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A simple tool that help you remove unneeded files from `node_modules` folder.

## Installation

From source:

```bash
➜  git clone https://github.com/ductnn/fano.git
➜  cd fano
➜  fano git:(master) ✗ go build main.go
➜  fano git:(master) ✗ ls
README.md go.mod    go.sum    internal  main      main.go
```

Then, you can copy file `main` to project nodejs and run:

```bash
./main node_modules
➜  nodejs-boilerplate git:(master) ✗ ./main node_modules

         files total 78,206
       files removed 20,069
        size removed 55 MB
            duration 5.551s
```
## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
