package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*Index function to handle the / path*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

/*TodoIndex function to handle the /todos path*/
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host Meetup"},
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

/*TodoShow function to handle the /todos path*/
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
