# Logrus Tools

## Middleware.

#### Example:
```golang
package main

import (
	"fmt"
	"net/http"

	logrusTools "github.com/Sidney-Bernardin/logrus-tools"
)

func example(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "example")
}

func main() {
	http.HandleFunc("/example", logrusTools.Middleware(example))
	http.ListenAndServe(":8080", nil)
}
```

#### Output:
```
Aug  3 12:43:34.257 [INFO] [id:f31f6c9f-8cc5-46ee-ad44-3e57c1832292] [addr:[::1]:12345] [status:200] [method:GET] [url:/example] [time:702.321µs] [chars:22] [referer:] [user-agent:curl/7.64.0] Handled request!
```
