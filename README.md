# YT Search
### A go (golang) library to search videos in YouTube.

## Installation
```bash
go get github.com/AnjanaMadu/YTSearch
```

## Usage
```go
package main

import (
    "fmt",
    "github.com/AnjanaMadu/YTSearch"
)

func main() {
    results, err := ytsearch.Search("faded")
    if err != nil {
        panic(err)
    }

    for _, result := range results {
        fmt.Printf("Title: %s\nVideo Id: %s\n\n", result.Title, result.VideoId)
    }
}
```

## Contributing
- Fork the repository
- Create a branch
- Commit your changes
- Push your branch to your fork
- Create a new issue or pull request

## License
**AGPL-3.0**, see [LICENSE](LICENSE) for more information.