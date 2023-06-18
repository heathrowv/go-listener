package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func returnHost(w http.ResponseWriter, r *http.Request) {

	datetime := time.Now()
	w.Write([]byte(r.Host + "\n"))
	w.Write([]byte(datetime.String()))

}

func main() {

	http.HandleFunc("/", returnHost)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)

}
