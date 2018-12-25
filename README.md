# CSV Reader [![Build Status](https://travis-ci.org/Illuminasy/csvreader.svg?branch=master)](https://travis-ci.org/Illuminasy/csvreader) [![GoDoc](https://godoc.org/github.com/Illuminasy/csvreader?status.svg)](https://godoc.org/github.com/Illuminasy/csvreader) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Illuminasy/csvreader/blob/master/LICENSE.md)

 Simple Golang interface to extract csv data.
 
# Usage

Get the library:

    $ go get -v github.com/Illuminasy/csvreader

```go
package somepackge

import "github.com/Illuminasy/csvreader"

func someFunction() {
	lines, err := csvreader.ExtractContents("some_file_path.csv")
}
```