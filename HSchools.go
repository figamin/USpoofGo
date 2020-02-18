package main

import (
	"fmt"
	"log"
)

func hSchoolSelection() {
	fmt.Println("Select your school's state:\n" +
		"1.  Colorado\n" +
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
		"22. North Carolina\n" +
		"23. Hawaii\n" +
		"24. Arizona\n" +
		"25. North Dakota\n" +
		"26. Kentucky\n" +
		"27. Missouri\n" +
		"28. Virginia\n" +
		"29. Nebraska\n" +
		"30. Georgia\n" +
		/*"31. \n" +
		"32. \n" +
		"33. \n" +
		"34. \n" +
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
			"1.  Littleton Lions")
		_, err = fmt.Scanln(&input2)
		if err != nil {
			log.Fatalln(err)
		}
		switch input2 {
		case 1:
			nid = "540"
		}
	}
}
