# USpoofGo
USpoofGo is [USpoof](https://github.com/figman57/USpoof), rewritten in Go, and is a GPS spoofing bot for SuperFanU based programs.

  - Automatically fetches events and checks into them
  - Live event countdown
  - Google Maps preview of event location
  - Randomized platform, device UUID, check in time, and GPS Location
  - Support for many schools (suggestions are welcome)
  - CLI flags for easier scripting
### Supported Schools List
[Here](SCHOOLS.md)
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
EVENT ID = 5YAIK
EVENT TITLE = Ice Hockey
EVENT DESCRIPTION = Hockey vs. Boston College | 7:00 p.m.
*Check-in to receive 2 points (starting at 6:00 p.m.)
POINT VALUE = 2
START TIME = Fri, 17 Jan 2020 23:04:00 UTC
END TIME = Sat, 18 Jan 2020 02:00:00 UTC
LATITUDE = 42.650307619810654
LONGITUDE = -71.3132348932541
RANDOM CLOSE LATITUDE = 42.650349
RANDOM CLOSE LONGITUDE = -71.31314
EXACT LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=42.650307619810654,-71.3132348932541
RANDOM CLOSE LOCATION PREVIEW = https://www.google.com/maps/search/?api=1&query=42.650349,-71.31314
Event starts in 2 days, 5 hours, and 47 minutes
```
