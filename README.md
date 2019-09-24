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
	log.Tracef("this is %s", log.TRACE) //  not display
	log.Debugf("this is %s", log.DEBUG) //  2019/09/25 07:15:42 main.go:10: [DEBUG] this is [DEBUG]
	log.Infof("this is %s", log.INFO)   //  2019/09/25 07:15:42 main.go:11: [INFO] this is [INFO]
	log.Warnf("this is %s", log.WARN)   //  2019/09/25 07:15:42 main.go:12: [WARN] this is [WARN]
	log.Errorf("this is %s", log.ERROR) //  2019/09/25 07:15:42 main.go:13: [ERROR] this is [ERROR]
}
```
