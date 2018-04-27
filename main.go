package main

import (
	"time"
	"test/iq/iq"
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

var cache *iq.Cache

func GetCacheAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, cache)
}

func GetCache(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if val, ok := cache.Get(params["key"]); ok {
		fmt.Fprint(w,val)
	}
}

func DeleteCache(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cache.Remove(params["key"])
}

func CreateCache(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	t, err := time.ParseDuration(params["time"])
	if err != nil {
		fmt.Fprint(w, err)
	}
	cache.Set(params["key"], params["val"], t*time.Minute)
	fmt.Fprint(w, "created")
}

func CreateCacheTop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	t, err := time.ParseDuration(params["time"])
	if err != nil {
		fmt.Fprint(w, err)
	}
	cache = iq.NewCache(t * time.Minute)
	fmt.Fprint(w, cache, "top cache created")
}

func GetKeys(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, cache.Keys())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Works server")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/all", GetCacheAll).Methods("GET")
	router.HandleFunc("/cache/{key}", GetCache).Methods("GET")
	router.HandleFunc("/keys", GetKeys).Methods("GET")
	router.HandleFunc("/cache_create/{time}", CreateCacheTop).Methods("POST")
	router.HandleFunc("/cache_create/{key}/{val}/{time}", CreateCache).Methods("POST")
	router.HandleFunc("/cache/{key}", DeleteCache).Methods("DELETE")


	log.Println("Server started at http://localhost:8080")
	log.Println("GET methods: /all; /cache/{key}; /keys ")
	log.Println("POST methods: /cache_create/{timeDuration}; /cache_create/{key}/{val}/{timeDuration} ")
	log.Println("DELETE methods: /cache/{key}")
	log.Fatal(http.ListenAndServe(":8080", router))
}