package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Exploration inspired by: https://eli.thegreenplace.net/2019/on-concurrency-in-go-http-servers/

type Counter struct {
	counters map[string]int
}

func (cs Counter) get(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	name := req.URL.Query().Get("name")
	if val, ok := cs.counters[name]; ok {
		fmt.Fprintf(w, "%d\n", val)
	} else {
		fmt.Fprintf(w, "%s Not Found\n", name)
	}
}

func (cs Counter) set(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	name := req.URL.Query().Get("name")
	val := req.URL.Query().Get("val")
	intval, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(w, "%s\n", name)
	} else {
		cs.counters[name] = intval
		fmt.Fprintf(w, "ok\n")
	}
}

func (cs Counter) inc(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	name := req.URL.Query().Get("name")
	if _, ok := cs.counters[name]; ok {
		cs.counters[name]++
		fmt.Fprintf(w, "ok\n")
	} else {
		fmt.Fprintf(w, "%s Not Found\n", name)
	}
}

func main() {
	store := Counter{counters: map[string]int{"i": 0, "j": 0}}
	http.HandleFunc("/get", store.get)
	http.HandleFunc("/set", store.set)
	http.HandleFunc("/inc", store.inc)

	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Listening on port %d", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}
