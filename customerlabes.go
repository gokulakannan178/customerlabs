package main

import (
	"customerlabes/models"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var (
	requestChannel   chan map[string]string
	convertedChannel chan models.Converted
)

func main() {
	requestChannel = make(chan map[string]string)
	convertedChannel = make(chan models.Converted)
	go worker()

	router := http.NewServeMux()
	router.HandleFunc("/https://webhook.site/183f3bce-44bb-474f-b94d-dcfdec51f1f3", ProcessHandler)
	server := &http.Server{
		Addr:    ":8010",
		Handler: router,
	}
	fmt.Println("Server listening on :8010")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string

	decoder := json.NewDecoder(r.Body)
	r.Header.Add("Content-Type", "application/json")
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	requestChannel <- req
	json.NewEncoder(w).Encode(<-convertedChannel)

}
func Convert(m map[string]string) {
	convertRequest := new(models.Converted)
	convertRequest.Event = m["ev"]
	convertRequest.EventType = m["et"]
	convertRequest.AppID = m["id"]
	convertRequest.UserID = m["uid"]
	convertRequest.MessageID = m["mid"]
	convertRequest.PageTitle = m["t"]
	convertRequest.PageURL = m["p"]
	convertRequest.BrowserLanguage = m["l"]
	convertRequest.ScreenSize = m["cs"]
	convertRequest.Attributes = make(map[string]models.Attribute)
	convertRequest.UserTraits = make(map[string]models.Attribute)
	pattern := "^atrk.*"
	pattern1 := "^uatrk.*"
	search, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	search1, err := regexp.Compile(pattern1)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	for key, value := range m {
		if search.MatchString(key) {
			str := strings.Split(key, "atrk")
			v := "atrv" + str[1]
			t := "atrt" + str[1]
			var atr models.Attribute
			atr.Value = m[v]
			atr.Type = m[t]
			convertRequest.Attributes[value] = atr
		}
		if search1.MatchString(key) {
			str := strings.Split(key, "uatrk")
			v := "uatrv" + str[1]
			t := "uatrt" + str[1]
			var atr models.Attribute
			atr.Value = m[v]
			atr.Type = m[t]
			convertRequest.UserTraits[value] = atr
		}
	}
	convertedChannel <- *convertRequest
}

func worker() {
	for req := range requestChannel {
		Convert(req)
	}
}
