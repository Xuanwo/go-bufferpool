# go-bufferpool

## Usage

```go
import "github.com/Xuanwo/go-bufferpool"

var pool = bufferpool.New(1024)

func main() {
    buf := pool.Get()
    defer buf.Free()

    buf.AppendString("Hello, World!")

    buf.String() // "Hello, World!"
}
```

## References

Inspired by following projects:

- [uber/zap/buffer](https://go.uber.org/zap/buffer)
- [valyala/bytebufferpool](https://github.com/valyala/bytebufferpool)
