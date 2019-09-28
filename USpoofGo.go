package main

import (
	"fmt"
)
var user, pass, nid, platform, uuid string
var loginKey = ""
func main() {
	fmt.Println("Welcome to USpoofGo 1.0\nBy Ian Anderson, 2019" +
		"\nEnter your username: ")
	fmt.Scanln(&user)
	fmt.Println("Enter your password")
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
	case 1: nid = "694"
	case 2: nid = "627"
	case 3: nid = "447"
	case 4: nid = "50"
	case 5: nid = "51"
	case 6: nid = "10"
	case 7: nid = "8"
	case 8: nid = "103"
	case 9: nid = "112"
	case 10: nid = "114"
	case 11: nid = "20"
	case 12: nid = "21"
	case 13: nid = "119"
	case 14: nid = "124"
	case 15: nid = "28"
	case 16: nid = "547"
	case 17: nid = "128"
	case 18: nid = "412"
	case 19: nid = "175"
	case 20: nid = "185"
	default: fmt.Println("Enter custom number: ")
		fmt.Scanln(&nid)
	}
}