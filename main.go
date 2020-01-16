package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tcolgate/mp3"
)

func main() {

	http.HandleFunc("/mp3", myHandler) //	设置访问路由
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		html, err := os.Open("test.html")
		if err != nil {
			w.Write([]byte{})
		}
		b, _ := ioutil.ReadAll(html)
		w.Header().Add("content-type", "text/html")
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandler(w http.ResponseWriter, rq *http.Request) {
	skipped := 0
	r, err := os.Open("test.mp3")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte{})
	}
	d := mp3.NewDecoder(r)
	f := &mp3.Frame{}
	for {

		if err := d.Decode(f, &skipped); err != nil {
			fmt.Println(err)
			return
		}

		mr := f.Reader()

		b, _ := ioutil.ReadAll(mr)

		w.Write(b)
	}
}
