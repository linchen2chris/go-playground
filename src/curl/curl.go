package curl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		s := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return
		}
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			return
		}
		secs := time.Since(s).Seconds()
		fmt.Printf("%.2fs %7d %s\n", secs, nbytes, url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
