package main;

import (
    "net/http"
    "os"
    "log"
)



func main() {
    bufferSize := os.Args[1]
    if bufferSize != "1" && bufferSize != "4" && bufferSize != "15"{
        log.Fatalln("Buffersize needs to be ither 1, 4 or 15.")
    }
    http.Handle("/", request(bufferSize))
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

//Router and logger
func request(bufferSize string) http.HandlerFunc {
    return (func(w http.ResponseWriter, r *http.Request){
        log.Println("Request: " + r.URL.Path)
        switch r.URL.Path {
        case "/":
            serveIndex(w, r)
        case "/video.mpd":
            serveMeta(w, r, bufferSize)
        default:
            serveVideo(w, r, bufferSize)
        }
    });
}

//Serve index file 
func serveIndex(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "./index.html")
}

//Serve meta file 
func serveMeta(w http.ResponseWriter, r *http.Request, bufferSize string){
     http.ServeFile(w, r, "./video/" + bufferSize + "sec/BigBuckBunny_" + bufferSize + "s_simple_2014_05_09.mpd")
}

//Serve video files
func serveVideo(w http.ResponseWriter, r *http.Request, bufferSize string) {
     http.ServeFile(w, r, "./video/" + bufferSize + "sec/")
}