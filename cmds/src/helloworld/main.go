package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"bob.expert/cmds/src/bob.expert/lib/designer"
)

func main() {
	var buf bytes.Buffer

	b := &buf

	fmt.Fprintf(&b, "Server time at %s is %s\n",
		os.Getenv("SERVER_NAME"), time.Now().Format(time.RFC1123))
	b = designer.SimpleView(b)
	fmt.Println("Content-type: text/plain")
	fmt.Printf("Content-Length: %d\n\n", b.Len())
	b.WriteTo(os.Stdout)
}
