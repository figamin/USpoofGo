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
var eventIDs, eventTitles, eventDescriptions, pointValues []string
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
	fmt.Println("Welcome to USpoofGo 1.1\nBy Ian Anderson, 2019" +
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
	fmt.Println("Select your school's state:\n" +
		"1.  Massachusetts\n" +
		"2.  California\n" +
		"3.  Connecticut\n" +
		"4.  New Hampshire\n" +
		"5.  Pennsylvania\n" +
		"6.  Maine\n" +
		"7.  Ohio\n" +
		"8.  Kansas\n" +
		"9.  Oregon\n" +
		"10. Iowa\n" +
		"11. Michigan\n" +
		"12. Texas\n" +
		"13. Rhode Island\n" +
		"14. Florida\n" +
		"15. Wisconsin\n" +
		"16. Indiana\n" +
		"17. Alabama\n" +
		"18. New York\n" +
		"19. New Jersey\n" +
		"20. Tennessee\n" +
		"21. Maryland\n" +
		"22. \n" +
		"23. \n" +
		"24. \n" +
		"25. \n" +
		"26. \n" +
		"27. \n" +
		"28. \n" +
		"29. \n" +
		"30. \n" +
		"0.  Manual NID Entry")
	var input int
	_, err = fmt.Scanln(&input)
	if err != nil {
		log.Fatalln(err)
	}
	var input2 int
	switch input {
	case 1: //mass
		fmt.Println("Select your school:\n" +
			"1.  UMass Minutemen\n" +
			"2.  UMass Lowell River Hawks\n" +
			"3.  Holy Cross Crusaders\n" +
			"4.  Boston College Eagles\n" +
			"5.  Boston University Terriers\n" +
			"6.  Northeastern University Huskies\n" +
			"7.  Harvard University Crimson\n" +
			"8.  Babson Beavers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "539"
		case 2:
			nid = "694"
		case 3:
			nid = "244"
		case 4:
			nid = "47"
		case 5:
			nid = "51"
		case 6:
			nid = "675"
		case 7:
			nid = "50"
		case 8:
			nid = "761"
		}
	case 2:
		fmt.Println("Select your school:\n" +
			"1.  California Golden Bears\n" +
			"2.  UCLA Bruins\n" +
			"3.  California Baptist Lancers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "539"
		case 2:
			nid = "694"
		case 3:
			nid = "244"
		}
	case 3:
		fmt.Println("Select your school:\n" +
			"1.  Sacred Heart Pioneers\n" +
			"2.  Fairfield Stags")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "543"
		case 2:
			nid = "20"
		}
	case 4:
		fmt.Println("Select your school:\n" +
			"1.  New Hampshire Wildcats\n" +
			"2.  Dartmouth Big Green\n" +
			"3.  Keene State Hooties")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			fmt.Println("School not working...")
		case 2:
			nid = "435"
		case 3:
			nid = "175"
		}
	case 5:
		fmt.Println("Select your school:\n" +
			"1.  Penn Quakers\n" +
			"2.  Bucknell Bison\n" +
			"3.  Lafayette Leopards")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "435"
		case 2:
			nid = "392"
		case 3:
			nid = "426"
		}
	case 6:
		fmt.Println("Select your school:\n" +
			"1.  Maine Black Bears")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "10"
		}
	case 7:
		fmt.Println("Select your school:\n" +
			"1.  Xavier Musketeers\n" +
			"2.  Cincinnati Bearcats")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			fmt.Println("School not working...")
		case 2:
			nid = "246"
		}
	case 8:
		fmt.Println("Select your school:\n" +
			"1.  Wichita State Shockers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "21"
		}
	case 9:
		fmt.Println("Select your school:\n" +
			"1.  George Fox Bruins")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "605"
		}
	case 10:
		fmt.Println("Select your school:\n" +
			"1.  Briar Cliff Chargers\n" +
			"2.  Drake Bulldogs\n" +
			"3.  Iowa State Cyclones")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "541"
		case 2:
			nid = "627"
		case 3:
			nid = "689"
		}
	case 11:
		fmt.Println("Select your school:\n" +
			"1.  Eastern Michigan Eagles\n" +
			"2.  Grand Valley State Lakers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			fmt.Println("School not working...")
		case 2:
			nid = "412"
		}
	case 12:
		fmt.Println("Select your school:\n" +
			"1.  St. Mary's Rattlers\n" +
			"2.  Texas A&M–Kingsville Javelinas")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "395"
		case 2:
			nid = "472"
		}
	case 13:
		fmt.Println("Select your school:\n" +
			"1.  Brown Bears")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "604"
		}
	case 14:
		fmt.Println("Select your school:\n" +
			"1.  Florida Gators\n" +
			"2.  North Florida Ospreys")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "35"
		case 2:
			nid = "672"
		}
	case 15:
		fmt.Println("Select your school:\n" +
			"1.  Marquette Golden Eagles")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "445"
		}
	case 16:
		fmt.Println("Select your school:\n" +
			"1.  Indiana State Sycamores")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "40"
		}
	case 17:
		fmt.Println("Select your school:\n" +
			"1.  North Alabama Lions")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			fmt.Println("School not working...")
		}
	case 18:
		fmt.Println("Select your school:\n" +
			"1.  Saint Rose Golden Knights")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "447"
		}
	case 19:
		fmt.Println("Select your school:\n" +
			"1.  Princeton Tigers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "149"
		}
	case 20:
		fmt.Println("Select your school:\n" +
			"1.  Lipscomb Bisons")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "642"
		}
	case 21:
		fmt.Println("Select your school:\n" +
			"1.  UMBC Retrievers\n" +
			"2.  Maryland Eastern Shore Hawks")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "310"
		case 2:
			nid = "392"
		}
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
				fmt.Println(" until " + eventTitles[i] + "\n")
				time.Sleep(time.Minute)
			}
			checkIn(i)
		}
		eventIDs = nil
		eventTitles = nil
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
			//fmt.Println(result5)
			var newStartTime, _ = time.Parse(time.RFC3339, strings.Replace(fmt.Sprint(result5["starttime"]), " ", "T", 1)+"Z")
			if newStartTime.Year() != 1 {
				eventIDs = append(eventIDs, fmt.Sprint(result5["eid"]))
				eventTitles = append(eventTitles, fmt.Sprint(result5["title"]))
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
		fmt.Println("EVENT TITLE = " + eventTitles[i])
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
		fmt.Println("No events currently. Sleeping for 1 hour.")
		//time.Sleep(time.Hour)
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
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Set("user-agent", "okhttp/3.11.0")
	request.Header.Set("content-length", "109")

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
