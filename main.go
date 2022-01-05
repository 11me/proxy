package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	fetch "github.com/11me/simplefetch"
)

var (
	PROXY_ADDR  = "http://127.0.0.1:8888"
	LISTEN_ADDR = "0.0.0.0:8888"
)

func proxy(origin string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// get the remote address from URI
		remote := strings.Trim(r.RequestURI, "/?")
		remoteAddr := fmt.Sprintf("http://%s", remote)

		if r.Method == http.MethodGet {
			// fetch the remote resource
			res, err := fetch.Get(fetch.Options{
				URL: remoteAddr,
			})
			if err != nil {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.Header().Set("Access-Control-Allow-Methods", "*")
				w.Write([]byte(`Bad request`))
				return
			}

			// read the fetched body of the resource
			fbody, _ := io.ReadAll(res.Body)
			// replace remote URLs to origin
			body, err := matchAndReplaceURL(remote, fbody)

			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Write(body)

		} else {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.WriteHeader(http.StatusOK)
		}
	}
}

func matchAndReplaceURL(remote string, body []byte) ([]byte, error) {

	input := fmt.Sprintf(`https?:\/\/%s`, remote)
	repl := fmt.Sprintf("%s/%s", PROXY_ADDR, remote)

	rgx, err := regexp.Compile(input)
	if err != nil {
		return nil, err
	}

	result := rgx.ReplaceAll(body, []byte(repl))

	return []byte(result), nil

}

func init() {
	// get server IP
	IP, ok := os.LookupEnv("IP")
	if ok {
		PROXY_ADDR = fmt.Sprintf("http://%s:8888", IP)
	}
}

func main() {

	http.HandleFunc("/", proxy(PROXY_ADDR))

	log.Printf("PROXY_ADDR is set to: %s\nListening for connections on %s\n", PROXY_ADDR, LISTEN_ADDR)

	log.Fatal(http.ListenAndServe(LISTEN_ADDR, nil))
}
