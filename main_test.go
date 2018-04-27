package main

import (
	"testing"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"test/iq/iq"
	"log"
	"os/exec"
)

func TestGetCacheAll(t *testing.T) {
	cache = iq.NewCache(30*time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, 3*time.Second)

	router := mux.NewRouter()
	router.HandleFunc("/all", GetCacheAll).Methods("GET")
	router.HandleFunc("/cache/{key}", GetCache).Methods("GET")
	router.HandleFunc("/keys", GetKeys).Methods("GET")
	router.HandleFunc("/cache_create/{time}", CreateCacheTop).Methods("POST")
	router.HandleFunc("/cache_create/{key}/{val}/{time}", CreateCache).Methods("POST")
	router.HandleFunc("/cache/{key}", DeleteCache).Methods("DELETE")
	go http.ListenAndServe(":8080", router)
	time.Sleep(time.Second*3)
	response, err := http.Get("http://localhost:8080/all")
	if err != nil {
		log.Println(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Method GetCacheAll doesn't work. Status %d", response.StatusCode)
	}

}

func TestGetCache(t *testing.T) {
	cache = iq.NewCache(3*time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, 3*time.Second)
	response, err := http.Get("http://localhost:8080/cache/one")
		if err != nil {
			log.Println(err)
		}
	if response.StatusCode != 200 {
		t.Errorf("Method GetCache doesn't work. Status %d", response.StatusCode)
	}
}

func TestGetKeys(t *testing.T) {
	cache = iq.NewCache(3*time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, 3*time.Second)
	response, err := http.Get("http://localhost:8080/keys")
	if err != nil {
		log.Println(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Method GetKeys doesn't work. Status %d", response.StatusCode)
	}
}

func TestCreateCacheTop(t *testing.T) {
	response, err := http.PostForm("http://localhost:8080/cache_create/3s", nil)
	if err != nil {
		log.Println(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Method CreateCacheTop doesn't work. Status %d", response.StatusCode)
	}
}

func TestCreateCache(t *testing.T) {
	response, err := http.PostForm("http://localhost:8080/cache_create/one/two/3s", nil)
	if err != nil {
		log.Println(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Method CreateCache doesn't work. Status %d", response.StatusCode)
	}
}

func TestDeleteCache(t *testing.T) {
	_, err := http.PostForm("http://localhost:8080/cache_create/three/two/6s", nil)
		if err != nil {
			log.Println(err)
		}

	exec.Command("curl", "-i http://localhost:8080/cache/three -XDELETE")

	c, _ := cache.Get("three")
	if c != nil {
		t.Errorf("Method DeleteCache isn't work")
	}
}