package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	myRelation  realtionship
	myLocation  locaco
	myArtist    fnan
	myDates     taree5
	all_Artists []fnan
)

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/moreinfo", hatInfo)
	http.HandleFunc("/search", search)
	fmt.Println("Server running on http://localhost:405")
	http.ListenAndServe(":405", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		fmt.Fprintln(w, "fi werror ya shater", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "jeep post 7abibi", http.StatusBadRequest)
		return
	}

	NameRes := []fnan{}
	LocRes := []fnan{}
	_1stAlbRes := []fnan{}
	CreRes := []fnan{}
	memRes := []fnan{}

	searchtxt := r.FormValue("searchfor")
	searchtxt = strings.ToLower(searchtxt)

	for _, theArtist := range all_Artists {
		if strings.Contains(strings.ToLower(theArtist.Name), searchtxt) {
			NameRes = append(NameRes, theArtist)
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(theArtist.Creation)), searchtxt) {
			CreRes = append(CreRes, theArtist)
		}
		if strings.Contains(strings.ToLower(theArtist.FirstAlbum), searchtxt) {
			_1stAlbRes = append(_1stAlbRes, theArtist)
		}

		for _, mem := range theArtist.Members {
			if strings.Contains(strings.ToLower(mem), searchtxt) {
				memRes = append(memRes, theArtist)
				break
			}
		}
	}

	var myAllloc AllLocations
	k := getData("https://groupietrackers.herokuapp.com/api/locations")
	err := json.Unmarshal(k, &myAllloc)
	checkForErrors(err)

	for artInd, allart := range myAllloc.Index {
		for _, theLoc := range allart.Locations {
			if strings.Contains(strings.ToLower(theLoc), searchtxt) {
				LocRes = append(LocRes, all_Artists[artInd])
				break
			}
		}
	}
	temp, err := template.ParseFiles("./searchResult.html")
	checkForErrors(err)

	allFoundData := map[string][]fnan{
		"byName":       NameRes,
		"byFirstAlbum": _1stAlbRes,
		"byMembers":    memRes,
		"byCreation":   CreRes,
		"byLocation":   LocRes,
		"all_Artists":  all_Artists,
	}

	err = temp.Execute(w, allFoundData)
	checkForErrors(err)

}

func getData(link string) []byte {
	theLink, err := http.Get(link)
	checkForErrors(err)
	data, err := ioutil.ReadAll(theLink.Body)
	checkForErrors(err)
	defer theLink.Body.Close()
	return data
}

func checkForErrors(err error) {
	if err != nil {
		fmt.Println("Oops, an error happened: ", err)
		return
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "Oops, an error happened: ", http.StatusNotFound)
		return
	}

	temp, err := template.ParseFiles("./index.html")
	checkForErrors(err)

	bodyart := getData("https://groupietrackers.herokuapp.com/api/artists")
	err = json.Unmarshal(bodyart, &all_Artists)
	checkForErrors(err)

	temp.Execute(w, all_Artists)
}

func hatInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/moreinfo" {
		fmt.Fprintln(w, "fi werror ya 7marr", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "jeep post 7abibi", http.StatusBadRequest)
		return
	}

	num := r.FormValue("artistNumber")
	getMyArtistData(num)

	temp, err := template.ParseFiles("./information.html")
	checkForErrors(err)

	artistDetails := Details{
		Artists:           myArtist,
		Locations:         myLocation,
		Concertdates:      myDates,
		DatesAndLocations: myRelation,
	}
	temp.Execute(w, artistDetails)
}

func getMyArtistData(num string) {
	qira2a := getData("https://groupietrackers.herokuapp.com/api/artists/" + num)
	err := json.Unmarshal(qira2a, &myArtist)
	checkForErrors(err)
	qira2aito := getData("https://groupietrackers.herokuapp.com/api/locations/" + num)
	erririri := json.Unmarshal(qira2aito, &myLocation)
	checkForErrors(erririri)
	qira2aotion := getData("https://groupietrackers.herokuapp.com/api/dates/" + num)
	erririri = json.Unmarshal(qira2aotion, &myDates)
	checkForErrors(erririri)
	qira2ariri := getData("https://groupietrackers.herokuapp.com/api/relation/" + num)
	erririri = json.Unmarshal(qira2ariri, &myRelation)
	checkForErrors(erririri)
}
