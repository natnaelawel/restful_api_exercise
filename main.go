package main
//import (
//	mux2 "github.com/gorilla/mux"
//	"net/http"
//	"encoding/json"
//)
//type Person struct {
//	ID string `json:"id"`
//	FirstName string `json:"first_name"`
//	LastName string `json:"last_name"`
//	Address *Address `json:"address"`
//}
//type Address struct {
//	City string `json:"city,omitempty"`
//	Country string `json:"country"`
//}
//
//var People []Person
//func readPeople(wr http.ResponseWriter, req *http.Request) {
//	//json.NewEncoder(wr).Encode(People)
//	wr.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(wr).Encode(People)
//
//}
//func createPeople(wr http.ResponseWriter, req *http.Request){
//	wr.Header().Set("Content-Type", "application/json")
//
//	var person Person
//	params := mux2.Vars(req)
//	id := params["id"]
//	_ = json.NewDecoder(req.Body).Decode(&person)
//	person.ID = id
//	People = append(People, person)
//	json.NewEncoder(wr).Encode(People)
//
//}
//func updatePeople(wr http.ResponseWriter, r *http.Request){
//	wr.Header().Set("Content-Type", "application/json")
//	params := mux2.Vars(r)
//	for index, item := range People{
//		if item.ID == params["id"] {
//			var person Person
//			_ = json.NewDecoder(r.Body).Decode(&person)
//			person.ID = params["id"]
//
//			People = append(People[:index], person)

//			json.NewEncoder(wr).Encode(People)
//
//	}
//	}
//}
//func deletePeople(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//
//	params := mux2.Vars(request)
//	for index, item := range People{
//		if item.ID == params["id"] {
//			People = append(People[:index], People[:index]...)
//			break
//		}
//	}
//	json.NewEncoder(writer).Encode(People)
//
//}
//func readSinglePeople(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//
//	params := mux2.Vars(request)
//		for _, item := range People{
//			if item.ID == params["id"] {
//				json.NewEncoder(writer).Encode(item)
//			}
//		}
//}
//
//func main(){
//	mux := mux2.NewRouter()
//
//	People = append(People, Person{ID:"1",FirstName:"Nathaniel", LastName:"Awel",Address:&Address{"Hawassa", "Ethiopia"}})
//	People = append(People, Person{ID:"2",FirstName:"Matiwos", LastName:"Bekele",Address:&Address{"Addis Ababa", "Ethiopia"}})
//	People = append(People, Person{ID:"3",FirstName:"Abebe", LastName:"Kebede",Address:&Address{"Hawassa", "Ethiopia"}})
//
//	mux.HandleFunc("/people", readPeople).Methods("GET")
//	mux.HandleFunc("/people/{id}", readSinglePeople).Methods("GET")
//	mux.HandleFunc("/people/{id}", createPeople).Methods("POST")
//	mux.HandleFunc("/people/{id}", deletePeople).Methods("DELETE")
//	mux.HandleFunc("/people/{id}", updatePeople).Methods("PATCH")
//
//	http.ListenAndServe(":12345", mux)
//
//}
//
