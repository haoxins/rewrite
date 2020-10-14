
[![Build status][travis-img]][travis-url]
[![PkgGoDev](https://pkg.go.dev/badge/pkg4go/rewrite)](https://pkg.go.dev/github.com/pkg4go/rewrite)

### rewrite

golang URL rewriting

### Usage

```go
import "github.com/pkg4go/rewrite"

// ...

handler := rewrite.NewHandler(map[string]string{
  "/a": "/b",
  "/api/(.*)", "/api/v1/$1",
  "/api/(.*)/actions/(.*)", "/api/v1/$1/actions/$2",
  "/from/:one/to/:two", "/from/:two/to/:one",
})

// ...
```

### License
MIT

[travis-img]: https://img.shields.io/travis/pkg4go/rewrite.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/rewrite
