# go-logger
go-logger is wrapper from standard Go logger.

# Installtion
```
go get github.com/skanehira/go-logger
```

# Usage
## Basic
```go
package main

import (
	log "github.com/skanehira/go-logger"
)

func main() {
	log.SetMinLevel(log.DEBUG)
	log.Tracef("this is %s", log.TRACE) // not output
	log.Debugf("this is %s", log.DEBUG) // 2019/09/25 07:15:42 main.go:10: [DEBUG] this is [DEBUG]
	log.Infof("this is %s", log.INFO)   // 2019/09/25 07:15:42 main.go:11: [INFO] this is [INFO]
	log.Warnf("this is %s", log.WARN)   // 2019/09/25 07:15:42 main.go:12: [WARN] this is [WARN]
	log.Errorf("this is %s", log.ERROR) // 2019/09/25 07:15:42 main.go:13: [ERROR] this is [ERROR]
}
```

## New logger
```go
package main

import (
	"os"

	log "github.com/skanehira/go-logger"
)

func main() {
	logger := log.New(log.INFO, "[test]", os.Stdout, log.Lshortfile|log.LstdFlags)
	logger.Tracef("this is %s", log.TRACE) // not output
	logger.Debugf("this is %s", log.DEBUG) // not output
	logger.Infof("this is %s", log.INFO)   // [test]2019/09/25 08:12:02 main.go:13: [INFO] this is [INFO]
	logger.Warnf("this is %s", log.WARN)   // [test]2019/09/25 08:12:02 main.go:14: [WARN] this is [WARN]
	logger.Errorf("this is %s", log.ERROR) // [test]2019/09/25 08:12:02 main.go:15: [ERROR] this is [ERROR]
}
```
