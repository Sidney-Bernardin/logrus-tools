# Logrus Tools

## Middleware example.
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
