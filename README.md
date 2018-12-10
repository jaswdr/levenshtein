# Levenshtein

### Getting started

Download the package

```sh
$ go get github.com/jaswdr/levenshtein
```

### Usage example

```go
package main

import (
  "fmt"
  
  "github.com/jaswdr/levenshtein"
)

func main() {
  str1 := "book"
  str2 := "back"
  
  diff := levenshtein.Compare(str1, str2)
  fmt.Println(diff) // 2
}
```
