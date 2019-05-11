# Set
[![CircleCI branch](https://img.shields.io/circleci/project/github/trivigy/set/master.svg?label=master&logo=circleci)](https://circleci.com/gh/trivigy/workflows/set)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE.md)
[![](https://godoc.org/github.com/trivigy/set?status.svg&style=flat)](http://godoc.org/github.com/trivigy/set)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/trivigy/set.svg?style=flat&color=e36397&label=release)](https://github.com/trivigy/set/releases/latest)

Set is a threadsafe abstract data structure library for representing collection 
of distinct values, without any particular order.

### Examples
```go
package main

import (
    "fmt"
	
    "github.com/trivigy/set"
)

func main() {
    m := set.New("b", "b", "c", "d")
    fmt.Println(m.Contains("b", "c", "d"))
	
    m1 := set.New("b", "b", "c", "d")
    m2 := set.New("c", "b", "d", "b")
    fmt.Println(m1.Equals(m2))
}
```
