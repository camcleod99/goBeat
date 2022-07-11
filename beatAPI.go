package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
	Global Variables
	Outputs - an array of data to output to Json
*/
type Output struct {
	RLTime   string `json:"RLTime"`
	BeatTime string `json:"BeatTime"`
}

var Outputs []Output

/* Request Handler */
func handleRequests() {
	// Instance Router
	myRouter := mux.NewRouter().StrictSlash(true)
	// Paths and Webpage Functions
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/beat", returnBeatTime)
	// myRouter replaces nil as the second argument when logging Fatal
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

/*
	Webpage Functions
	HomePage : Index Page
	returnBeatTime : Outputs the data as JSON
*/
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage!")
	fmt.Println("Endpoint Hit: HP")
}

func returnBeatTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Beats")
	setOutputs()
	json.NewEncoder(w).Encode(Outputs)
}

/*
	Basic Functions
	getBeat : Gets the time in Swatch Internet Time
	setOutputs : Sets the output for displaying in the JSON
*/
func getBeat() string {
	ct := time.Now()
	ctHoursStr := ct.Format("15")
	ctMinutesStr := ct.Format("04")
	ctSecondsStr := ct.Format("05")

	ctHours, _ := strconv.Atoi(ctHoursStr)
	ctMinutes, _ := strconv.Atoi(ctMinutesStr)
	ctSeconds, _ := strconv.Atoi(ctSecondsStr)

	/* Cast into float32 to allow division by 86.4 */
	totalSeconds := float32(((ctHours * 60) * 60) + (ctMinutes * 60) + ctSeconds)

	/* Cast into Int to format into proper format */
	beatTime := int(totalSeconds / 86.4)

	beatTimeString := strconv.Itoa(beatTime)

	/* Return value of function */
	return beatTimeString
}

func setOutputs() {
	t := time.Now()
	tOut := t.Format("2006-01-02 : 15:04")
	beat := getBeat()
	Outputs = []Output{
		Output{RLTime: tOut, BeatTime: beat},
	}
}

func main() {
	handleRequests()
}
