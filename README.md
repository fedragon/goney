# goney

[![Build Status](https://travis-ci.org/fedragon/goney.svg?branch=master)](https://travis-ci.org/fedragon/goney)

Represents monetary amounts in Go with arbitrary-precision arithmetic on their amounts.

## Usage

```
package main

import (
  "fmt"
  "github.com/fedragon/goney"
)

func main () {
	a := goney.FromFloat(goney.EUR, 1.23)

	// Ignoring errors to keep the example concise:
	// don't do this at home!
	b, _ := goney.FromString("EUR 3.45")
	c, _ := a.Add(b)

	fmt.Println(c)
}
```