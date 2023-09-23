# shuffler

Go package with 6 popular shuffling algorithms]

## Installation

```
go get github.com/nemzyxt/shuffler
```

## Usage

Example:

```
package main

import (
    "fmt"
    "github.com/nemzyxt/shuffler"
)

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    shuffler.FisherYatesShuffle[int](arr[:]) // use the Fisher-Yates algorithm
    fmt.Println(arr)
}
```

## Supported Algorithms

- Fisher-Yates Shuffle
- Knuth Shuffle
- Sattolo's Shuffle
- Durstenfeld Shuffle
- Cohen's Shuffle
- Riffle/Perfect Shuffle
