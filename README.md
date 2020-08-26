# MP4
Basic MP4 reader in Go!

CLI and library for ISO/IEC 14496-12 - ISO Base Media File Format

https://godoc.org/github.com/alfg/mp4

[![Build Status](https://travis-ci.org/alfg/mp4.svg?branch=master)](https://travis-ci.org/alfg/mp4) 
[![Build status](https://ci.appveyor.com/api/projects/status/63ky9q869j8xetst?svg=true)](https://ci.appveyor.com/project/alfg/mp4)
[![GoDoc](https://godoc.org/github.com/alfg/mp4?status.svg)](https://godoc.org/github.com/alfg/mp4)
[![Go Report Card](https://goreportcard.com/badge/github.com/alfg/mp4)](https://goreportcard.com/report/github.com/alfg/mp4)

## Usage

```
go get -u github.com/alfg/mp4
```

```go
package main

import (
    "fmt"
    "github.com/alfg/mp4"
)

func main() {
    file, _ := mp4.Open("test/tears-of-steel.mp4")
    file.Close()
    fmt.Println(file.Ftyp.Name)
    fmt.Println(file.Ftyp.MajorBrand)
}
```

See [example.go](/example/example.go) for a full example.

## Develop 

```
git clone https://github.com/alfg/mp4.git
go run example\example.go
```

Or build the CLI:
```
go build -o mp4info cmd\mp4info\mp4info.go
mp4info -i test\tears-of-steel.mp4
```

## Documentation
* [GoDocs](https://godoc.org/github.com/alfg/mp4)

## License
MIT
