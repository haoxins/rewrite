[![Build Status Badge]][Build Status]
[![Go Docs Badge]][Go Docs]

### Rewrite

Golang URL rewriting

### Usage

```go
import "github.com/haoxins/rewrite"

// ...

handler := rewrite.NewHandler(map[string]string{
  "/a": "/b",
  "/api/(.*)", "/api/v1/$1",
  "/api/(.*)/actions/(.*)", "/api/v1/$1/actions/$2",
  "/from/:one/to/:two", "/from/:two/to/:one",
})

// ...
```

[Build Status Badge]: https://github.com/haoxins/rewrite/actions/workflows/test.yaml/badge.svg
[Build Status]: https://github.com/haoxins/rewrite/actions/workflows/test.yaml
[Go Docs Badge]: https://pkg.go.dev/badge/github.com/haoxins/rewrite
[Go Docs]: https://pkg.go.dev/github.com/haoxins/rewrite
