package main;

import (
    "net/http"
    "log"
)


var bufferSize = "4"

func main() {
    http.Handle("/", request())
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

//Router and logger
func request() http.HandlerFunc {
    return (func(w http.ResponseWriter, r *http.Request){
        log.Println("Request: " + r.URL.Path)
        switch r.URL.Path {
        case "/":
            serveIndex(w, r)
        case "/video.mpd":
            serveMeta(w, r)
        default:
            serveVideo(w, r)
        }
    });
}

//Serve index file 
func serveIndex(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "./index.html")
}

//Serve meta file 
func serveMeta(w http.ResponseWriter, r *http.Request){
     http.ServeFile(w, r, "./video/" + bufferSize + "sec/BigBuckBunny_" + bufferSize + "s_simple_2014_05_09.mpd")
}

//Serve video files
func serveVideo(w http.ResponseWriter, r *http.Request) {
     http.ServeFile(w, r, "./video/" + bufferSize + "sec/")
}