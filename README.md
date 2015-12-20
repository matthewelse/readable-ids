# readable-ids

Human Readable Identifiers in Go. Generates a random, memorable, human-readable string.

## Usage

```go
package main

import (
    "fmt"
    "crypto/rand"
    "github.com/matthewelse/readable-ids"
)

func main() {
    fmt.Println(readableid.GenID(rand.Reader, 100))
}
```

## Improvements

- [ ] Increase the number of possible combinations by adding verbs and adverbs.


