package main

import (
	"fmt"
	"log"
)

func uSchoolSelection() {
	fmt.Println("Select your school's state:\n" +
		"1.  Alabama\n" +
		"2.  Arizona\n" +
		"3.  California\n" +
		"4.  Colorado\n" +
		"5.  Connecticut\n" +
		"6.  Florida\n" +
		"7.  Georgia\n" +
		"8.  Hawaii\n" +
		"9.  Illinois\n" +
		"10. Indiana\n" +
		"11. Iowa\n" +
		"12. Kansas\n" +
		"13. Kentucky\n" +
		"14. Maine\n" +
		"15. Maryland\n" +
		"16. Massachusetts\n" +
		"17. Michigan\n" +
		"18. Missouri\n" +
		"19. Nebraska\n" +
		"20. New Hampshire\n" +
		"21. New Jersey\n" +
		"22. New York\n" +
		"23. North Carolina\n" +
		"24. North Dakota\n" +
		"25. Ohio\n" +
		"26. Oregon\n" +
		"27. Pennsylvania\n" +
		"28. Rhode Island\n" +
		"29. South Carolina\n" +
		"30. Tennessee\n" +
		"31. Texas\n" +
		"32. Virginia\n" +
		"33. Wisconsin\n" +
		/*"34. \n" +
		"35. \n" +
		"36. \n" +
		"37. \n" +
		"38. \n" +
		"39. \n" +
		"40. \n" +*/
		"0.  Manual NID Entry")
	var input int
	_, err = fmt.Scanln(&input)
	if err != nil {
		log.Fatalln(err)
	}
	var input2 int
	switch input {
	case 1:
		fmt.Println("Select your school:\n" +
			"1.  North Alabama Lions\n" +
			"2.  Auburn Montgomery Warhawks\n" +
			"3.  Samford Bulldogs")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "61"
		case 2:
			nid = "598"
		case 3:
			nid = "348"
		}
	case 2:
		fmt.Println("Select your school:\n" +
			"1.  Arizona Wildcats")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "167"
		}
	case 3:
		fmt.Println("Select your school:\n" +
			"1.  California Golden Bears\n" +
			"2.  UCLA Bruins\n" +
			"3.  California Baptist Lancers\n" +
			"4.  Cal State Dominguez Hills Toros")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "89"
		case 2:
			nid = "45"
		case 3:
			nid = "441"
		case 4:
			nid = "382"
		}
	case 4:
		fmt.Println("Select your school:\n" +
			"1.  Colorado Christian Cougars")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "617"
		}
	case 5:
		fmt.Println("Select your school:\n" +
			"1.  Sacred Heart Pioneers\n" +
			"2.  Fairfield Stags\n" +
			"3.  New Haven Chargers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "543"
		case 2:
			nid = "20"
		case 3:
			nid = "207"
		}
	case 6:
		fmt.Println("Select your school:\n" +
			"1.  Florida Gators\n" +
			"2.  North Florida Ospreys\n" +
			"3.  South Florida Bulls\n" +
			"4.  Florida Gulf Coast Eagles")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "35"
		case 2:
			nid = "672"
		case 3:
			nid = "127"
		case 4:
			nid = "60"
		}
	case 7:
		fmt.Println("Select your school:\n" +
			"1.  Georgia Bulldogs" +
			"2.  Georgia Southern Eagles")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "90"
		case 2:
			nid = "773"
		}
	case 8:
		fmt.Println("Select your school:\n" +
			"1.  Hawaii Rainbow Warriors")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "547"
		}
	case 9:
		fmt.Println("Select your school:\n" +
			"1.  Western Illinois Leathernecks")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "520"
		}
	case 10:
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
	case 11:
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
	case 12:
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
	case 13:
		fmt.Println("Select your school:\n" +
			"1.  Western Kentucky Hilltoppers\n" +
			"2.  Centre College Colonels")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "625"
		case 2:
			nid = "440"
		}
	case 14:
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
	case 15:
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
	case 16: //mass
		fmt.Println("Select your school:\n" +
			"1.  UMass Minutemen\n" +
			"2.  UMass Lowell River Hawks\n" +
			"3.  Holy Cross Crusaders\n" +
			"4.  Boston College Eagles\n" +
			"5.  Boston University Terriers\n" +
			"6.  Northeastern University Huskies\n" +
			"7.  Harvard University Crimson\n" +
			"8.  Babson Beavers\n" +
			"9.  Springfield Squad\n" +
			"10. Brandeis Judges\n" +
			"11. Bentley Falcons\n" +
			"12. Assumption College")
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
		case 9:
			nid = "622"
		case 10:
			nid = "471"
		case 11:
			nid = "109"
		case 12:
			nid = "193"
		}
	case 17:
		fmt.Println("Select your school:\n" +
			"1.  Eastern Michigan Eagles\n" +
			"2.  Grand Valley State Lakers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "41"
		case 2:
			nid = "412"
		}
	case 18:
		fmt.Println("Select your school:\n" +
			"1.  UMSL Tritons\n" +
			"2.  Saint Louis Billikens\n" +
			"3.  Missouri Tigers\n" +
			"4.  Missouri Baptist Spartans")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "191"
		case 2:
			nid = "628"
		case 3:
			nid = "95"
		case 4:
			nid = "730"
		}
	case 19:
		fmt.Println("Select your school:\n" +
			"1.  Chadron State Eagles")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "609"
		}
	case 20:
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
			nid = "157"
		case 2:
			nid = "435"
		case 3:
			nid = "175"
		}
	case 21:
		fmt.Println("Select your school:\n" +
			"1.  Princeton Tigers\n" +
			"2.  NJIT Highlanders\n" +
			"3.  Seton Hall Pirates")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "149"
		case 2:
			nid = "501"
		case 3:
			nid = "17"
		}
	case 22:
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
	case 23:
		fmt.Println("Select your school:\n" +
			"1.  UNC Pembroke Braves\n" +
			"2.  UNC Charlotte 49ers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "185"
		case 2:
			nid = "124"
		}
	case 24:
		fmt.Println("Select your school:\n" +
			"1.  North Dakota Fighting Hawks")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "706"
		}
	case 25:
		fmt.Println("Select your school:\n" +
			"1.  Xavier Musketeers\n" +
			"2.  Cincinnati Bearcats\n" +
			"3.  Akron Zips")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "367"
		case 2:
			nid = "246"
		case 3:
			nid = "390"
		}
	case 26:
		fmt.Println("Select your school:\n" +
			"1.  George Fox Bruins\n" +
			"2.  Willamette Bearcats")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "605"
		case 2:
			nid = "405"
		}
	case 27:
		fmt.Println("Select your school:\n" +
			"1.  Penn Quakers\n" +
			"2.  Bucknell Bison\n" +
			"3.  Lafayette Leopards\n" +
			"4.  Edinboro Fighting Scots")
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
		case 4:
			nid = "487"
		}
	case 28:
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
	case 29:
		fmt.Println("Select your school:\n" +
			"1.  South Carolina Gamecocks")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "52"
		}
	case 30:
		fmt.Println("Select your school:\n" +
			"1.  Lipscomb Bisons\n" +
			"2.  Tusculum Pioneers\n" +
			"3.  Lee Flames")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "642"
		case 2:
			nid = "563"
		case 3:
			nid = "114"
		}
	case 31:
		fmt.Println("Select your school:\n" +
			"1.  St. Mary's Rattlers\n" +
			"2.  Texas A&M–Kingsville Javelinas\n" +
			"3.  Incarnate Word Cardinals")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "395"
		case 2:
			nid = "472"
		case 3:
			nid = "565"
		}
	case 32:
		fmt.Println("Select your school:\n" +
			"1.  Virginia Tech Hokies\n" +
			"2.  Liberty Flames")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "166"
		case 2:
			nid = "720"
		}
	case 33:
		fmt.Println("Select your school:\n" +
			"1.  Marquette Golden Eagles\n" +
			"2.  UW Milwaukee Panthers\n" +
			"3.  UW Platteville Pioneers")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "445"
		case 2:
			nid = "65"
		case 3:
			nid = "432"
		}
	}
}
