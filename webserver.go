package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var avaiableStreams []string

func main() {
	log.Println("Scanning video/ directory for mpd files.")
	if err := filepath.Walk("./video/", discover); err != nil {
		log.Fatalln("Could not read video/ directory.")
	}
	log.Println("Server is up and running at http://localhost:8080/")
	http.Handle("/", request())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func discover(path string, f os.FileInfo, err error) error {
	filename := f.Name()
	if strings.HasSuffix(filename, ".mpd") && err == nil {
		avaiableStreams = append(avaiableStreams, "/"+path)
	}
	return nil
}

//Router and logger
func request() http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		switch r.URL.Path {
		case "/":
			serveIndex(w, r)
		case "/streams":
			listStreams(w, r)
		default:
			serveVideo(w, r)
		}
		elapsed := time.Since(start) / 1000
		log.Printf("HTTP Request: %s %s, Total time: %dms\n", r.Method, r.URL.Path, elapsed)
	})
}

//Serve index file
func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./videoplayer.html")
}

//List available streams
func listStreams(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(avaiableStreams)
}

//Serve video files
func serveVideo(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}
