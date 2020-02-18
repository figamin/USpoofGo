package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var platform, uuid, ustring string
var user = "0"
var pass = "0"
var nid = "0"
var loginKey = ""
var oldEventID = ""
var hclient = http.Client{}
var eventID, eventTitle, eventDescription, pointValue string
var generatedLatitude, generatedLongitude float64
var startTime, endTime time.Time
var err error
var retry = false
var debug = true

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.StringVar(&user, "user", "0", "SuperFanU Username")
	flag.StringVar(&pass, "pass", "0", "SuperFanU Password")
	flag.StringVar(&nid, "nid", "0", "SuperFanU School ID")
	flag.Parse()
	fmt.Println("Welcome to USpoofGo 1.3\nBy Ian Anderson, 2020")
	if user == "0" {
		fmt.Println("Enter your username:")
		_, err = fmt.Scanln(&user)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if pass == "0" {
		fmt.Println("Enter your password:")
		_, err = fmt.Scanln(&pass)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if nid == "0" {
		fmt.Println("Select your school list:\n" +
			"1.  Colleges/Universities\n" +
			"2.  High Schools\n" +
			"0.  Manual NID Entry")
		var input int
		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Fatalln(err)
		}
		switch input {
		case 1:
			uSchoolSelection()
		case 2:
			hSchoolSelection()
		default:
			fmt.Println("Enter custom number: ")
			_, err = fmt.Scanln(&nid)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	logfile, err := os.OpenFile("USpoofLog-School-"+nid, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	possiblePlatforms := []string{"iOS", "Android"}
	platform = possiblePlatforms[rand.Intn(2)]
	if platform == "iOS" {
		ustring = "SuperFan/1 CFNetwork/1121.2.2 Darwin/19.2.0"
	} else {
		ustring = "okhttp/3.11.0"
	}
	uuid = generateUUID()
	logIn()
	getNextEvent()
	eventPrinter()
	for {
		var times = time.Second
		for times > time.Duration(0) {
			if debug {
				times = time.Duration(0)
				time.Sleep(time.Second * 5)
			} else {
				times = startTime.Sub(time.Now())
				time.Sleep(time.Minute)
			}

			printTime(times)
			fmt.Println(" until " + eventTitle + "\n")

		}
		checkIn()
		times = time.Second
		for times > time.Duration(0) {
			if debug {
				times = time.Duration(0)
				time.Sleep(time.Second * 5)
			} else {
				times = endTime.Sub(time.Now())
				times -= time.Minute * 15
				time.Sleep(time.Minute)
			}
			printTime(times)
			fmt.Println(" until fetching the next event.\n")

		}
		getNextEvent()
		eventPrinter()
	}
}

func generateUUID() string {
	validChars := "0123456789abcdefghijklmnopqrstuvxyz"
	var uuidBuilder strings.Builder
	for i := 0; i < 16; i++ {
		uuidBuilder.WriteString(string(validChars[rand.Intn(len(validChars))]))
	}
	return uuidBuilder.String()
}

func logIn() {
	fmt.Println("Logging in with username " + user + " at school ID " + nid + "...")
	request, err := http.PostForm("https://api.superfanu.com/8.0/login",
		url.Values{
			"user":     {user},
			"pass":     {pass},
			"nid":      {nid},
			"platform": {platform},
			"uuid":     {uuid},
		})
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(request.Body)
	var results interface{}
	if request.StatusCode == 200 {
		fmt.Println("Login Successful!")
	}
	err = json.Unmarshal([]byte(string(body)), &results)
	if err != nil {
		log.Fatalln(err)
	}
	result3 := results.(map[string]interface{})["data"].([]interface{})
	for _, v := range result3 {
		loginKey = fmt.Sprint(v.(map[string]interface{})["login_key"])
	}
}

func getNextEvent() {
	fmt.Printf("Getting feed...")
	request, _ := http.NewRequest("GET", "https://api.superfanu.com/8.0/feed", nil)
	request.Header.Add("nid", nid)
	request.Header.Add("platform", platform)
	request.Header.Add("uuid", uuid)
	request.Header.Add("login_key", loginKey)
	request.Header.Set("user-agent", ustring)
	response, _ := hclient.Do(request)
	if response.StatusCode == 200 {
		fmt.Println("Feed successfully fetched!")
	}
	body, _ := ioutil.ReadAll(response.Body)
	var results interface{}
	err = json.Unmarshal([]byte(string(body)), &results)
	if err != nil {
		log.Fatalln(err)
	}
	result3 := results.(map[string]interface{})["data"].([]interface{})
	for _, v := range result3 {
		var result4 = v.([]interface{})
		resultCount := 0
		for eventID == oldEventID {
			var result5 = result4[resultCount].(map[string]interface{})
			var newStartTime, _ = time.Parse(time.RFC3339, strings.Replace(fmt.Sprint(result5["starttime"]), " ", "T", 1)+"Z")
			if newStartTime.Year() != 1 {
				eventID = fmt.Sprint(result5["eid"])
				eventTitle = fmt.Sprint(result5["title"])
				eventDescription = fmt.Sprint(result5["description"])
				pointValue = fmt.Sprint(result5["pointvalue"])
				_, offset := time.Now().Zone()
				startTime = newStartTime.Add((time.Second * time.Duration(-offset)) + (time.Hour + time.Minute*time.Duration(rand.Intn(10))))
				var newEndTime, _ = time.Parse(time.RFC3339, strings.Replace(fmt.Sprint(result5["endtime"]), " ", "T", 1)+"Z")
				endTime = newEndTime.Add((time.Second * time.Duration(-offset)) + (time.Hour + time.Minute))
			}
			resultCount++
		}

	}
}

func eventPrinter() {
	fmt.Println("EVENT ID = " + eventID)
	fmt.Println("EVENT TITLE = " + eventTitle)
	fmt.Println("EVENT DESCRIPTION = " + eventDescription)
	fmt.Println("POINT VALUE = " + pointValue)
	fmt.Println("START TIME = " + startTime.Format(time.RFC1123))
	fmt.Println("END TIME = " + endTime.Format(time.RFC1123))
	log.Println("EVENT ID = " + eventID)
	log.Println("EVENT TITLE = " + eventTitle)
	log.Println("EVENT DESCRIPTION = " + eventDescription)
	log.Println("POINT VALUE = " + pointValue)
	log.Println("START TIME = " + startTime.Format(time.RFC1123))
	log.Println("END TIME = " + endTime.Format(time.RFC1123))
	request, _ := http.NewRequest("GET", "https://api.superfanu.com/8.0/event/"+eventID+"/details", nil)
	request.Header.Add("nid", nid)
	request.Header.Add("platform", platform)
	request.Header.Add("uuid", uuid)
	request.Header.Add("login_key", loginKey)
	request.Header.Set("user-agent", ustring)
	response, _ := hclient.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	var results interface{}
	err = json.Unmarshal([]byte(string(body)), &results)
	if err != nil {
		log.Fatalln(err)
	}
	result3 := results.(map[string]interface{})["data"].([]interface{})
	for _, v := range result3 {
		result4 := v.(map[string]interface{})["event"].(map[string]interface{})["venues"].([]interface{})
		for _, u := range result4 {
			var result5 = u.(map[string]interface{})
			var latitude = fmt.Sprint(result5["latitude"])
			var longitude = fmt.Sprint(result5["longitude"])
			fmt.Println("LATITUDE = " + latitude)
			fmt.Println("LONGITUDE = " + longitude)
			generatedLatitude, _ = strconv.ParseFloat(latitude, 64)
			generatedLongitude, _ = strconv.ParseFloat(longitude, 64)
			generatedLatitude = ((math.Round(generatedLatitude*10000) * 100) + float64(rand.Intn(100))) / 1000000.0
			generatedLongitude = ((math.Round(generatedLongitude*10000) * 100) + float64(rand.Intn(100))) / 1000000.0
			fmt.Println("RANDOM CLOSE LATITUDE = " + fmt.Sprint(generatedLatitude))
			fmt.Println("RANDOM CLOSE LONGITUDE = " + fmt.Sprint(generatedLongitude))
			fmt.Println("EXACT LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=" + latitude + "," + longitude)
			fmt.Println("RANDOM CLOSE LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=" + fmt.Sprint(generatedLatitude) + "," + fmt.Sprint(generatedLongitude))
		}
	}
	fmt.Print("Event starts in ")
	printTime(startTime.Sub(time.Now()))
	fmt.Println()
	fmt.Println()

	if eventID == oldEventID {
		fmt.Println("No events currently. Sleeping for 1 hour.")
		time.Sleep(time.Hour)
	}
}

func checkIn() {
	data := url.Values{
		"eid":                 {eventID},
		"latitude":            {fmt.Sprint(generatedLatitude)},
		"longitude":           {fmt.Sprint(generatedLongitude)},
		"altitude":            {fmt.Sprint(rand.Intn(100))},
		"horizontal_accuracy": {fmt.Sprint(rand.Intn(20))},
		"vertical_accuracy":   {fmt.Sprint(rand.Intn(20))},
	}
	request, _ := http.NewRequest("POST", "https://api.superfanu.com/8.0/event/check-in", strings.NewReader(data.Encode()))
	request.Header.Add("nid", nid)
	request.Header.Add("platform", platform)
	request.Header.Add("uuid", uuid)
	request.Header.Add("login_key", loginKey)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Set("content-length", strconv.Itoa(len(data.Encode())))
	request.Header.Set("user-agent", ustring)
	response, _ := hclient.Do(request)
	if response.StatusCode == 200 {
		fmt.Println("Checked In Successfully!")
	} else {
		fmt.Println("Failed to check in. Retrying...")
		logIn()
		checkIn()
	}
	oldEventID = eventID

}

func printTime(d time.Duration) {
	var days = d.Hours() / 24
	_, hours := math.Modf(days)
	hours *= 24
	_, minutes := math.Modf(hours)
	minutes *= 60
	fmt.Print(fmt.Sprint(int(days))+" days, "+fmt.Sprint(int(hours))+" hours, and "+fmt.Sprint(int(minutes)), " minutes")
}
