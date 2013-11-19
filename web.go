package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", printIP)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func printIP(res http.ResponseWriter, req *http.Request) {

	realIP, ok := req.Header["X-Real-Ip"]
	if ok {
		fmt.Fprintln(res, cleanUp(realIP[0]))
		return
	}

	forwardedFor, ok := req.Header["X-Forwarded-For"]
	if ok {
		fmt.Fprintln(res, cleanUp(forwardedFor[0]))
		return
	}

	fmt.Fprintln(res, cleanUp(req.RemoteAddr))
}

func cleanUp(address string) string {
	address = strings.TrimPrefix(address, "::ffff:")
	return strings.Split(address, ":")[0]
}
