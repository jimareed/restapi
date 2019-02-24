package main

import (
	"os"
	"net/http"
	"strconv"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Fruit struct {
	ID		string	`json:"id,omitempty"`
	Type	string	`json:"type,omitempty"`
}

var fruit []Fruit

func nextID() string {
    next := 0
    for _, afruit := range fruit {
		id, err := strconv.Atoi(afruit.ID)
        if err == nil && id > next {
			next = id
        }
    }
	
	return strconv.Itoa(next+1)
}

func getFruit(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(fruit)
}

func getAFruit(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, afruit := range fruit {
        if afruit.ID == params["id"] {
            json.NewEncoder(w).Encode(afruit)
            return
        }
    }
    json.NewEncoder(w).Encode(&Fruit{})
}

func postFruit(w http.ResponseWriter, r *http.Request) {
    var afruit Fruit
    _ = json.NewDecoder(r.Body).Decode(&afruit)
    afruit.ID = nextID()
	fruit = append(fruit, afruit)
}

func deleteFruit(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, afruit := range fruit {
        if afruit.ID == params["id"] {
            fruit = append(fruit[:index], fruit[index+1:]...)
            break
        }
    }
}	

func main() {
	fruit = append(fruit, Fruit{ID: "0", Type: "apple"})
	fruit = append(fruit, Fruit{ID: "1", Type: "orange"})
	fruit = append(fruit, Fruit{ID: "2", Type: "pear"})
	
	port, ok := os.LookupEnv("PORT")
	
	if !ok {
		port = "8080"
	}
	
	router := mux.NewRouter()
    router.HandleFunc("/api/fruit", getFruit).Methods("GET")
    router.HandleFunc("/api/fruit/{id}", getAFruit).Methods("GET")
    router.HandleFunc("/api/fruit", postFruit).Methods("POST")
    router.HandleFunc("/api/fruit/{id}", deleteFruit).Methods("DELETE")

	http.ListenAndServe(":" + port, router)

}

