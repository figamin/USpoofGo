package main

import (
	"encoding/json"
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

var user, pass, nid, platform, uuid string
var loginKey = ""
var hclient = http.Client{}
var eventIDs, eventDescriptions, pointValues []string
var generatedLatitidues, generatedLongitudes []float64
var startTimes, endTimes []time.Time
var err error

func main() {
	logfile, err := os.OpenFile("USpoofLog", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to USpoofGo 1.0.2\nBy Ian Anderson, 2019" +
		"\nEnter your username: ")
	_, err = fmt.Scanln(&user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Enter your password: ")
	_, err = fmt.Scanln(&pass)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Select your school:\n" +
		"1.  UMass Lowell\n" +
		"2.  Drake University\n" +
		"3.  College of St. Rose\n" +
		"4.  Harvard University\n" +
		"5.  Boston University\n" +
		"6.  University of Maine\n" +
		"7.  Southeast Missouri State University\n" +
		"8.  California State University\n" +
		"9.  University of Toledo\n" +
		"10. Lee University\n" +
		"11. Fairfield University\n" +
		"12. Wichita State University\n" +
		"13. University of Kentucky\n" +
		"14. University of North Carolina at Charlotte\n" +
		"15. University of Pennsylvania\n" +
		"16. University of Hawaii at Manoa\n" +
		"17. Minot State University\n" +
		"18. Grand Valley State University\n" +
		"19. Keene State College\n" +
		"20. University of North Carolina at Pembroke\n" +
		"0.  Manual NID Entry")
	var input int
	_, err = fmt.Scanln(&input)
	if err != nil {
		log.Fatalln(err)
	}
	switch input {
	case 1:
		nid = "694"
	case 2:
		nid = "627"
	case 3:
		nid = "447"
	case 4:
		nid = "50"
	case 5:
		nid = "51"
	case 6:
		nid = "10"
	case 7:
		nid = "8"
	case 8:
		nid = "103"
	case 9:
		nid = "112"
	case 10:
		nid = "114"
	case 11:
		nid = "20"
	case 12:
		nid = "21"
	case 13:
		nid = "119"
	case 14:
		nid = "124"
	case 15:
		nid = "28"
	case 16:
		nid = "547"
	case 17:
		nid = "128"
	case 18:
		nid = "412"
	case 19:
		nid = "175"
	case 20:
		nid = "185"
	default:
		fmt.Println("Enter custom number: ")
		_, err = fmt.Scanln(&nid)
		if err != nil {
			log.Fatalln(err)
		}
	}
	possiblePlatforms := []string{"iOS", "Android"}
	platform = possiblePlatforms[rand.Intn(2)]
	uuid = generateUUID()
	logIn()
	getFeed()
	eventPrinter()
	for {
		for i := range eventIDs {
			var times = time.Second
			for times > time.Duration(0) {
				times = startTimes[i].Sub(time.Now())
				//times = time.Duration(0)
				printTime(times)
				fmt.Println(" until " + eventDescriptions[i] + "\n")
				time.Sleep(time.Minute)
			}
			checkIn(i)
		}
		eventIDs = nil
		eventDescriptions = nil
		pointValues = nil
		startTimes = nil
		endTimes = nil
		generatedLatitidues = nil
		generatedLongitudes = nil
		getFeed()
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

	request, err := http.PostForm("https://api.superfanu.com/7.0.1/login",
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

func getFeed() {
	fmt.Printf("Getting feed...")
	request, _ := http.NewRequest("GET", "https://api.superfanu.com/7.0.1/feed", nil)
	request.Header.Add("nid", nid)
	request.Header.Add("platform", platform)
	request.Header.Add("uuid", uuid)
	request.Header.Add("login_key", loginKey)
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
		for _, u := range result4 {
			var result5 = u.(map[string]interface{})
			var newStartTime, _ = time.Parse(time.RFC3339, strings.Replace(fmt.Sprint(result5["starttime"]), " ", "T", 1)+"Z")
			if newStartTime.Year() != 1 {
				eventIDs = append(eventIDs, fmt.Sprint(result5["eid"]))
				eventDescriptions = append(eventDescriptions, fmt.Sprint(result5["description"]))
				pointValues = append(pointValues, fmt.Sprint(result5["pointvalue"]))
				_, offset := time.Now().Zone()
				newStartTime = newStartTime.Add(time.Second * time.Duration(-offset))
				startTimes = append(startTimes, newStartTime.Add(time.Hour+time.Minute*time.Duration(rand.Intn(10))))
				var newEndTime, _ = time.Parse(time.RFC3339, strings.Replace(fmt.Sprint(result5["endtime"]), " ", "T", 1)+"Z")
				newEndTime = newEndTime.Add(time.Second * time.Duration(-offset))
				endTimes = append(endTimes, newEndTime.Add(time.Hour))
			}
		}
	}
}

func eventPrinter() {
	for i := range eventIDs {
		fmt.Println("EVENT ID = " + eventIDs[i])
		fmt.Println("EVENT DESCRIPTION = " + eventDescriptions[i])
		fmt.Println("POINT VALUE = " + pointValues[i])
		fmt.Println("START TIME = " + startTimes[i].Format(time.RFC1123))
		fmt.Println("END TIME = " + endTimes[i].Format(time.RFC1123))
		request, _ := http.NewRequest("GET", "https://api.superfanu.com/7.0.1/event/"+eventIDs[i]+"/details", nil)
		request.Header.Add("nid", nid)
		request.Header.Add("platform", platform)
		request.Header.Add("uuid", uuid)
		request.Header.Add("login_key", loginKey)
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
				var randCloseLatitude float64
				var randCloseLongitude float64
				randCloseLatitude, _ = strconv.ParseFloat(latitude, 64)
				randCloseLongitude, _ = strconv.ParseFloat(longitude, 64)
				randCloseLatitude = ((math.Round(randCloseLatitude*10000) * 100) + float64(rand.Intn(100))) / 1000000.0
				randCloseLongitude = ((math.Round(randCloseLongitude*10000) * 100) + float64(rand.Intn(100))) / 1000000.0
				fmt.Println("RANDOM CLOSE LATITUDE = " + fmt.Sprint(randCloseLatitude))
				fmt.Println("RANDOM CLOSE LONGITUDE = " + fmt.Sprint(randCloseLongitude))
				generatedLatitidues = append(generatedLatitidues, randCloseLatitude)
				generatedLongitudes = append(generatedLongitudes, randCloseLongitude)
				fmt.Println("EXACT LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=" + latitude + "," + longitude)
				fmt.Println("RANDOM CLOSE LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=" + fmt.Sprint(randCloseLatitude) + "," + fmt.Sprint(randCloseLongitude))
			}
		}
		fmt.Print("Event starts in ")
		printTime(startTimes[i].Sub(time.Now()))
		fmt.Println()
		fmt.Println()
	}

	if len(eventIDs) == 0 {
		fmt.Println("No events currently. Check again later!")
		os.Exit(0)
	}
}

func checkIn(index int) {
	request, _ := http.NewRequest("POST", "https://api.superfanu.com/7.0.1/event/check-in", nil)
	request.PostForm = url.Values{
		"eid":                 {eventIDs[index]},
		"latitude":            {fmt.Sprint(generatedLatitidues[index])},
		"longitude":           {fmt.Sprint(generatedLongitudes[index])},
		"altitude":            {fmt.Sprint(rand.Intn(100))},
		"horizontal_accuracy": {fmt.Sprint(rand.Intn(20))},
		"vertical_accuracy":   {fmt.Sprint(rand.Intn(20))},
	}
	request.Header.Add("nid", nid)
	request.Header.Add("platform", platform)
	request.Header.Add("uuid", uuid)
	request.Header.Add("login_key", loginKey)
	response, _ := hclient.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println(string(body) + "\n")
}

func printTime(d time.Duration) {
	var days = d.Hours() / 24
	_, hours := math.Modf(days)
	hours *= 24
	_, minutes := math.Modf(hours)
	minutes *= 60
	fmt.Print(fmt.Sprint(int(days))+" days, "+fmt.Sprint(int(hours))+" hours, and "+fmt.Sprint(int(minutes)), " minutes")
}
