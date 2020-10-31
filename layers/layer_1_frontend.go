package layers

// i tried to keep things minimal in this file, here only
// contains logic that is necessary to make a http frontend.

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func RunFrontend() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// global variable no good, but i use as it is easy for demo purpose
var notepadService NotepadService = NotepadImpl{
	dbService: getRedisDatabase(),
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		key := r.URL.Path
		val, err := notepadService.ReadNotepad(key)
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, val)
	} else if r.Method == "POST" {
		key := r.URL.Path
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		val := string(b)

		notepad := Notepad{
			ID:      key,
			Content: val,
		}

		if err := notepadService.WriteNotepad(notepad); err != nil {
			panic(err)
		}
		println("Thanks")
	}
}
