package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dummy2, Hello!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":9998", nil))
}

//docker rmi [image id]

// #before delete the image deleting the container
//docker rmi -f [image id]

// kubectl apply -f <deployment yaml path>
// kubectl delete -f <deployment yaml path>
