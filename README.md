# go-log-logr-adapter
A logr adapter for the go default logger

This module creates a standard go logger thet logs messages via a given logr logger.

## Usage 
```go
import (
	la "github.com/bakito/go-log-logr-adapter/adapter"
	"github.com/go-logr/logr"
)

func main() {
  var log logr.Logger
  
  stdLog := la.ToStd(log)
}

```
