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
	"github.com/alfg/mp4/mp4"
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
## TODO
* Add more atoms
* Tests

## License
MIT
