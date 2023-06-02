package datagather

import (
	"strings"
	"time"
)

/*
TitleCase converts an input string to its title case and returns it
e.g. big bob --> Big Bob, little-head --> Little-Head
*/
func TitleCase(s string) string {
	runeSlice := []rune(s)
	wordStart := false
	for i := 0; i < len(runeSlice); i++ {
		if !wordStart && runeSlice[i] >= 'a' && runeSlice[i] <= 'z' {
			// Correct to uppercase for start of word
			runeSlice[i] = runeSlice[i] - 32
			wordStart = true
		} else if !wordStart && runeSlice[i] >= 'A' && runeSlice[i] <= 'Z' {
			// Detect start of word which is correctly capitalised
			wordStart = true
		} else if wordStart && runeSlice[i] >= 'A' && runeSlice[i] <= 'Z' {
			// Correct to lowercase within word
			runeSlice[i] = runeSlice[i] + 32
		} else if (runeSlice[i] == ' ' || runeSlice[i] == '-') && wordStart {
			// Detect end of word
			wordStart = false
		}
	}
	return string(runeSlice)
}

/*
LocationCleaner reformats the location strings as per the following example:
eg. "north_carolina-usa" --> "USA: North Carolina".
It takes a slice of strings as input, and calls the CleanLocation function
on each of the slices elements
*/
func LocationCleaner(s []string) []string {
	var CleanedString []string
	for i := 0; i < len(s); i++ {
		CleanedString = append(CleanedString, CleanLocation(s[i]))
	}
	return CleanedString
}

/*
CleanLocation reformats an input location string so that its original JSON formatting
is corrected. Regions / countries are changed to upper case, underscore characters are
replaced by white space, and cities / locations are changed to their title case.
*/
func CleanLocation(s string) string {
	// A Location string should contain two sub elements (1. city / place; 2. country / region)
	// separated by a hyphen ('-')
	set := strings.Split(s, "-")
	for i := 0; i < len(set); i++ {
		// '_' is replaced by white space
		set[i] = strings.ReplaceAll(set[i], "_", " ")
		if i == 0 {
			// Every 1st / odd term is a city / place and all words changed to their title case
			set[i] = TitleCase(set[i])
			continue
		} else {
			// Every 2nd / even term is a country / region, which should be capitalised
			set[i] = strings.ToUpper(set[i])
			continue
		}
	}
	// Swap location elements to 1. country / region; 2. city / place
	return set[1] + ": " + set[0]
}

/*
DateFormat takes an input date string and reformats it to the desired format
of DD MONTH YEAR, e.g. 2 January 2006
*/
func DateFormat(s string) string {
	layoutExpected := "02-01-2006"
	layoutDesired := "2 January 2006"
	// Remove the asterisk that exists for some date entries from the API
	if s[0] == '*' {
		s = s[1:]
	}
	date, _ := time.Parse(layoutExpected, s)
	return date.Format(layoutDesired)
}

/*
RelationFormat takes relational data for a single artist as a map structure of
concert locations & dates, and a returns a cleaned and reformatted map
*/
func RelationFormat(s map[string][]string) map[string][]string {
	relationMap := make(map[string][]string)
	for concLoc, concDate := range s {
		concLoc = CleanLocation(concLoc)
		var concertDates []string
		for i := 0; i < len(concDate); i++ {
			concertDates = append(concertDates, DateFormat(concDate[i]))
		}
		relationMap[concLoc] = concertDates
	}
	return relationMap
}
