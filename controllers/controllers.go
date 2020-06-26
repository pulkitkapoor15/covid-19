package controllers

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Please stay safe and maintain distancing!")
}


func GetNationTimeSeries(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.covid19india.org/data.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	w.Write(body)
}
