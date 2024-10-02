package main

import (
	"net/http"
	// "html/template"
	"encoding/json"
	"fmt"
	"strconv"
)

type Data struct {
	ID 				int 	`json:"id"`
	Image 			string 	`json:"image"`
	Name 			string 	`json:"name"`
	Members 		[]string `json:"members"`
	CreationDate 	int 	`json:"creationDate"`
	FirstAlbum 		string 	`json:"firstAlbum"`
	Locations 		string 	`json:"locations"`
	ConcertDates 	string 	`json:"concertDates"`
	Relations 		string	`json:"relations"`
}

var data []Data

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}



func getResult(w http.ResponseWriter, r *http.Request) {
	text := []byte(`[{"id":1,"image":"https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
	"name":"Queen","members":["Freddie Mercury","Brian May","John Daecon",
	"Roger Meddows-Taylor","Mike Grose","Barry Mitchell","Doug Fogie"],
	"creationDate":1970,"firstAlbum":"14-12-1973","locations":"https://groupietrackers.herokuapp.com/api/locations/1",
	"concertDates":"https://groupietrackers.herokuapp.com/api/dates/1",
	"relations":"https://groupietrackers.herokuapp.com/api/relation/1"},
	{"id":2,"image":"https://groupietrackers.herokuapp.com/api/images/soja.jpeg",
	"name":"SOJA","members":["Jacob Hemphill","Bob Jefferson","Ryan \"Byrd\" Berty",
	"Ken Brownell","Patrick O'Shea","Hellman Escorcia","Rafael Rodriguez","Trevor Young"],
	"creationDate":1997,"firstAlbum":"05-06-2002","locations":"https://groupietrackers.herokuapp.com/api/locations/2",
	"concertDates":"https://groupietrackers.herokuapp.com/api/dates/2",
	"relations":"https://groupietrackers.herokuapp.com/api/relation/2"}]`)

	err := json.Unmarshal(text, &data)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
	// t, err := template.ParseFiles("home.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// w.Write((text))
	// t.Execute(w, text)
}
func displayData(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "data.html")
	
}
func a(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, item := range data {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}
func main() {
	http.HandleFunc("/data.html", displayData)
	http.HandleFunc("/ad", a)
	http.HandleFunc("/data", getResult)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}