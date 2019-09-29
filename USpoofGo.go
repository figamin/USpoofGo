package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var user, pass, nid, platform, uuid string
var loginKey = ""
var hclient = http.Client{}
var eventIDs, eventDescriptions, pointValues []string
var generatedLatitidues, generatedLongitudes []float32
var startTimes, endTimes []time.Time
var err error

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to USpoofGo 1.0\nBy Ian Anderson, 2019" +
		"\nEnter your username: ")
	fmt.Scanln(&user)
	fmt.Println("Enter your password: ")
	fmt.Scanln(&pass)
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
	fmt.Scanln(&input)
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
		fmt.Scanln(&nid)
	}
	possiblePlatforms := []string{"iOS", "Android"}
	platform = possiblePlatforms[rand.Intn(2)]
	uuid = generateUUID()
	logIn()
	getFeed()
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
	fmt.Println(loginKey)
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
	fmt.Println(result3)
}
