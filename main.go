package main

import (
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"fmt"
	"encoding/json"
)
type  Message  struct {
	Sender string	`json:"sender"`
	Data string `json:"data"`
}

type EndPoint  struct {
	Name string	`json:"name"`
	DCIn int	`json "dcIn"`
	DCOut int	`json "dcOut"`
}

type Scope struct {
	Name string `json:"name"`
	EndPoint EndPoint `json:"endPoint"`
}

type ServiceFromScope struct {
	ScopeProviderServices []string
	ScopeClientServices []string
}
func init(){
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetLevel(log.TraceLevel)
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		m := Message{ "Sender1", "Data"}
		if b, err := json.Marshal(m); err == nil{
			fmt.Fprintf(w,"%s",b)
			log.Trace("Sent")
		} else {
			log.Info("Failed")
		}
	})
	r.HandleFunc("/scopes/{key}", ScopeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func ScopeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	s := Scope{
		Name: vars["key"],
	}

	if b, err := json.Marshal(s); err == nil {
		fmt.Fprintf(w, "%s", b)
		log.Info("Scope Handled")
	} else {
		log.Error("Scope Handling Failed", err)
	}
}
