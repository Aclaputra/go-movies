package main;

import(
  "fmt"
  "log"
  "encoding/json"
  "math/rand"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"

  L "./lib"
)

type Movie struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}

type Director struct {
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

var movies []Movie

func main() {
  r := mux.NewRouter();
  
  movies = append(movies, Movie{
    ID: "1", 
    Isbn: "428392", 
    Title: "Spiderman: No Way Home", 
    Director: &Director{FirstName: "John", LastName: "Doe"},
  })
  movies = append(movies, Movie{
    ID: "2", 
    Isbn: "447923", 
    Title: "Thanos: Without Stones", 
    Director: &Director{FirstName: "Steve", LastName: "Smith"},
  })

  r.HandleFunc("/movies", L.getMovies).Methods("GET")
  r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  r.HandleFunc("/movies", createMovie).Methods("POST")
  r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
  r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

  fmt.Printf("Starting server at port 8000\n")
  log.Fatal(http.ListenAndServe(":8000", r))
}
