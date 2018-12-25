# CSV Reader

 Simple Golang interface to extract csv data.

# Usage

Get the library:

    $ go get -v github.com/Illuminasy/CSVReader

```go
package somepackge

import "github.com/Illuminasy/CSVReader"

func someFunction() {
	lines, err := CSVReader.ExtractContents("some_file_path.csv")
}
```