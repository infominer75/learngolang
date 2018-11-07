package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Usage: netfetcher <urls separated by space>")
		os.Exit(-1)
	}

	//loop through the specified URLs. check if the https prefix is present; if not add it
	start := time.Now()
	status := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		go writeURLToFile(url, status)

	}
	for _, url := range os.Args[1:] {
		fmt.Println("Status For url : ", url, <-status)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readAllInMemory(body io.ReadCloser) bool {
	b, err1 := ioutil.ReadAll(body)
	body.Close()
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Error %v reading response", err1)
		return false
	}
	fmt.Printf("%s\n\n", b)
	return true
}

func writeToStdout(resp *http.Response) {
	buffer := bytes.NewBufferString("")
	io.Copy(buffer, resp.Body)
	fmt.Println(buffer)
}

func writeURLToFile(url string, status chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v fetching url %s.", err, url)
		status <- fmt.Sprintf("Error %v fetching url %s.\n", err, url)
		return
	}
	//the below call is bad since it attempts to load the entire
	//response content within memory in one gulp. recommended is to use
	//io.Copy() to perform a streamed copy.

	//also tee the content to a file. the name of the file is the name of the URL
	f, err := os.Create("/tmp/" + strings.Replace(url, "https://", "", -1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v fetching url %s.", err, url)
		status <- fmt.Sprintf("Error %v fetching url %s.\n", err, url)
		return
	}
	io.Copy(f, resp.Body)
	f.Close()
	status <- fmt.Sprintf("Wrote URL:%s to file %s.\n", url, f.Name())
}
