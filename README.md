# MP4
Basic MP4 parser in Go!

*work-in-progress*

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
git clone github.com/alfg/mp4
go run example/example.go
```

Or build the CLI:
```
go build -o mp4info.exe cmd\mp4info\mp4info.go
mp4info -i test/tears-of-steel.mp4
```

## TODO
* Add more atoms
* Tests

## License
MIT
