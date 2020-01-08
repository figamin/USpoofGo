# USpoofGo
USpoofGo is [USpoof](https://github.com/figman57/USpoof), rewritten in Go, and is a GPS spoofing bot for SuperFanU based programs.

  - Automatically fetches events and checks into them
  - Live event countdown
  - Google Maps preview of event location
  - Randomized platform, device UUID, check in time, and GPS Location
  - Support for many schools (suggestions are welcome)
### Supported School List
[a relative link](SCHOOLS.md)
## Why rewrite in Go?
- Learning (good to get to know a new language)
- Better suited language (strange to write a single class CLI Java program)
- Reduced overhead (lack of JVM), useful for running multiple concurrent instances
- Easier to run (user does not need the JRE)
- No external libraries (used builtin json parsing, even if it is a pain)
### Preview
```
Logging in with username IanA57 at school ID 627...
Login Successful!
Getting feed...Feed successfully fetched!
EVENT ID = 2NXKF
EVENT DESCRIPTION = Drake Football vs. Valparaiso (Homecoming and Parent & Family Weekend)
Check-in begins at 1 p.m.
FREE ADMISSION for students!

STUDENT PROMOTIONS
* Student tailgate begins at 10am with music, games and free prizes for most spirited tailgater
* Reserve your spot in the student tailgate lot here: https://tinyurl.com/DogPoundTailgate 
*Intramural Bags Tournament - sign up via the Drake Rec App
*Gift Card giveaways throughout the game with $$ increasing each quarter
POINT VALUE = 10
START TIME = Sat, 05 Oct 2019 18:08:00 UTC
END TIME = Sat, 05 Oct 2019 21:00:00 UTC
LATITUDE = 41.604994118508856
LONGITUDE = -93.65506312209016
RANDOM CLOSE LATITUDE = 41.605094
RANDOM CLOSE LONGITUDE = -93.655001
EXACT LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=41.604994118508856,-93.65506312209016
RANDOM CLOSE LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=41.605094,-93.655001
Event starts in 140h59m34.702165268s
```
