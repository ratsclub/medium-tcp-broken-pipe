package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)

		// creating a large data size
		// that will take a long time to be written
		size := 900 * 1000 * 1000
		tpl := make([]byte, size)
		t, err := template.New("page").Parse(string(tpl))
		if err != nil {
			log.Printf("error parsing template: %s", err)
			return
		}

		if err := t.Execute(w, nil); err != nil {
			log.Printf("error writing: %s", err)
			return
		}
	})

	srv := &http.Server{
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		Addr:         ":8080",
		Handler:      mux,
	}

	log.Println("server is running!")
	log.Println(srv.ListenAndServe())
}
