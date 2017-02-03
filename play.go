// Command play shares a file on the Go playground.
//
// Usage:
//
//	$ play file.go
//	$ play < file.go
package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var body []byte

	if len(os.Args) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		body = b
	} else {
		b, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		body = b
	}

	req, _ := http.NewRequest("POST", "https://golang.org/share", bytes.NewReader(body))
	req.Header.Set("User-Agent", "github-com-broady-play")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("POST: ", err)
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stderr, resp.Body)
		log.Fatalf("%s", http.StatusText(resp.StatusCode))
	}
	os.Stdout.WriteString("https://play.golang.org/p/")
	io.Copy(os.Stdout, resp.Body)
	os.Stdout.WriteString("\n")
}
