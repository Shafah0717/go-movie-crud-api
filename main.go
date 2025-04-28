package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}",deleteMovie).Methods("DELETE")

	
	fmt.Printf("starting server at port 8000")
	log.Fatal(http.ListenAndServe("8000:",r))
}