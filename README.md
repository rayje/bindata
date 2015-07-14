# Bindata
## Overview
A small go project to embed text files as byte data

## Getting Started
The following steps will guide you through the process of using Bindata

### Installation
To install Bindata, just run `go get`

    $ go get github.com/rayje/bindata

### Running Bindata
To run the project, you will need a file to embed placed within a data directory.

* First create your data directory and file
```bash
$ mkdir data
$ echo "console.log('this is a test');" > data/test.js
```

* Next run Bindata, providing the data directory
```bash
$ bindata ./data
```

* This should then create a new file called bindata.go
```bash
$ ls 
bindata.go  data
```

* The contents of bindata.go should looks similar to this:
```bash
$ cat bindata.go
package main
    
import (
    "bytes"
    "compress/gzip"
    "io"
)
    
func asset(bs []byte) []byte {
    var b bytes.Buffer
    gz, _ := gzip.NewReader(bytes.NewReader(bs))
    io.Copy(&b, gz)
    gz.Close()
    return b.Bytes()
}
    
var TestJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xce\xcf\x2b\xce\xcf\x49\xd5\xcb\xc9\x4f\xd7\x50\x2f\xc9\xc8\x2c\x56\x00\xa2\x44\x85\x92\xd4\xe2\x12\x75\x4d\x6b\x2e\x40\x00\x00\x00\xff\xff\xe9\x00\x25\x7f\x1f\x00\x00\x00")
```

* You should then be able to access the contents of the byte data by calling the `asset` function
```bash
$ cat test.go
package main

import "fmt"

func main() {
    fmt.Println(string(asset(TestJs)))
}
```

* Running the step above should create the following output
```bash
$ go run test.go bindata.go
console.log('this is a test');
```
