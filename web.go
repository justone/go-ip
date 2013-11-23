package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", printIP)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "5000"
	}

	fmt.Printf("listening on %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func printIP(res http.ResponseWriter, req *http.Request) {

	// try proxy friendly headers
	for _, header := range []string{"X-Real-Ip", "X-Forwarded-For"} {
		realIP, ok := req.Header[header]
		if ok {
			fmt.Fprintln(res, cleanUp(realIP[0]))
			return
		}
	}

	// fall back to the remote address
	fmt.Fprintln(res, cleanUp(req.RemoteAddr))
}

func cleanUp(address string) string {
	address = strings.TrimPrefix(address, "::ffff:")
	return strings.Split(address, ":")[0]
}
