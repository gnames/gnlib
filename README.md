# gnlib

[![Go Reference](https://pkg.go.dev/badge/github.com/gnames/gnlib.svg)](https://pkg.go.dev/github.com/gnames/gnlib)

A collection of shared utilities and entities for Global Names Architecture Go projects.

## Features

- **User-Friendly Error Handling**: Terminal-colorized error messages with clean output
- **Generic Utilities**: Type-safe Map, Filter, and other collection operations
- **Channel Operations**: Chunk channel data into manageable batches
- **Version Comparison**: Semantic version comparison utilities
- **UTF-8 Handling**: Fix and normalize UTF-8 strings
- **Language Utilities**: Convert between language codes and names
- **Domain Entities**: Shared types for taxonomic name verification, reconciliation, and matching

## Installation

```bash
go get github.com/gnames/gnlib
```

## Usage

### User-Friendly Error Messages

The `gnfmt` error handling system allows you to create errors that produce clean, colorized output for the terminal, while preserving the underlying error details for logging.

```go
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gnames/gnfmt"
)

func main() {
	// Create a new error with a format string and variables.
	// The message includes tags for colorization.
	err := gnfmt.NewError(
		"<warning>Could not process file '%s'</warning>",
		[]any{"important.txt"},
	)

	// The UserMessage() method returns the formatted, colorized string.
	fmt.Fprintln(os.Stdout, err.UserMessage())

	// Example of wrapping the error.
	wrappedErr := fmt.Errorf("operation failed: %w", err)

	// You can inspect the error chain to get the user-friendly message.
	var gnErr gnfmt.Error
	if errors.As(wrappedErr, &gnErr) {
		fmt.Fprintln(os.Stdout, "---")
		fmt.Fprintln(os.Stdout, "Message from wrapped error:")
		fmt.Fprintln(os.Stdout, gnErr.UserMessage())
	}
}
```

This will produce the following output (with colors in the terminal):

```
Could not process file 'important.txt'
---
Message from wrapped error:
Could not process file 'important.txt'
```

#### Colorization Tags

The `UserMessage()` method recognizes the following tags for styling terminal output:

-   `<title>...</title>`: Renders text in **green**.
-   `<warning>...</warning>`: Renders text in **red**.
-   `<em>...</em>`: Renders text in **yellow**.

### Generic Utilities

The library provides type-safe generic functions for common operations:

```go
package main

import (
    "fmt"
    "github.com/gnames/gnlib"
)

func main() {
    // Map transforms a slice
    numbers := []int{1, 2, 3, 4, 5}
    doubled := gnlib.Map(numbers, func(n int) int { return n * 2 })
    fmt.Println(doubled) // [2 4 6 8 10]

    // Filter returns elements matching a condition
    evens := gnlib.FilterFunc(numbers, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens) // [2 4]

    // SliceMap creates a lookup map from a slice
    fruits := []string{"apple", "banana", "cherry"}
    fruitMap := gnlib.SliceMap(fruits)
    fmt.Println(fruitMap["banana"]) // 1
}
```

### Channel Operations

Process channel data in chunks:

```go
package main

import (
    "context"
    "fmt"
    "github.com/gnames/gnlib"
)

func main() {
    input := make(chan int)
    go func() {
        for i := 1; i <= 10; i++ {
            input <- i
        }
        close(input)
    }()

    chunked := gnlib.ChunkChannel(context.Background(), input, 3)
    for chunk := range chunked {
        fmt.Println(chunk)
    }
    // Output:
    // [1 2 3]
    // [4 5 6]
    // [7 8 9]
    // [10]
}
```

### Version Comparison

Compare semantic versions:

```go
package main

import (
    "fmt"
    "github.com/gnames/gnlib"
)

func main() {
    result := gnlib.CmpVersion("v1.2.3", "v1.2.4")
    fmt.Println(result) // -1 (first is less than second)

    result = gnlib.CmpVersion("v2.0.0", "v1.9.9")
    fmt.Println(result) // 1 (first is greater than second)

    result = gnlib.CmpVersion("v1.0.0", "v1.0.0")
    fmt.Println(result) // 0 (versions are equal)
}
```

### UTF-8 String Handling

Fix invalid UTF-8 sequences and normalize strings:

```go
package main

import (
    "fmt"
    "github.com/gnames/gnlib"
)

func main() {
    // Replaces invalid UTF-8 with U+FFFD and normalizes to NFC
    fixed := gnlib.FixUtf8("invalid\xc3\x28utf8")
    fmt.Println(fixed)
}
```

### Language Utilities

Convert between language codes and names:

```go
package main

import (
    "fmt"
    "github.com/gnames/gnlib"
)

func main() {
    // Get ISO 639-3 code from language name
    code := gnlib.LangCode("English")
    fmt.Println(code) // "eng"

    code = gnlib.LangCode("en")
    fmt.Println(code) // "eng"

    // Get language name from code
    name := gnlib.LangName("fra")
    fmt.Println(name) // "French"
}
```

## Domain Entities

The library includes shared entity types for taxonomic name processing:

- **`ent/verifier`**: Types for taxonomic name verification results
- **`ent/reconciler`**: Types for name reconciliation and manifests
- **`ent/matcher`**: Types for name matching operations
- **`ent/nomcode`**: Nomenclatural code enumerations
- **`ent/gnml`**: Global Names Markup Language types
- **`ent/gnvers`**: Version information types

See the [API documentation](https://pkg.go.dev/github.com/gnames/gnlib) for details on these packages.

## License

Released under the MIT License. See LICENSE file for details.

