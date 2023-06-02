package datagather

import (
	"reflect"
	"testing"
)

func TestTitleCase(t *testing.T) {
	oneWord := "super-fast"
	multipleWords := "multiple Test caSes"

	oneWord = TitleCase(oneWord)
	multipleWords = TitleCase(multipleWords)

	if oneWord != "Super-Fast" {
		t.Errorf("\ngot \"%s\", but wanted \"Super-fast\"", oneWord)
	} else if multipleWords != "Multiple Test Cases" {
		t.Errorf("\ngot \"%s\", but wanted \"Multiple Test Cases\"", multipleWords)
	}
}

func TestLocationCleaner(t *testing.T) {
	testSlice := []string{"california-usa", "rio_de_janeiro-brazil", "los_angeles-usa",
		"alabama-usa", "massachusetts-usa", "athens-greece", "florence-italy",
		"landgraaf-netherlands"}
	correctSlice := []string{"USA: California", "BRAZIL: Rio De Janeiro", "USA: Los Angeles",
		"USA: Alabama", "USA: Massachusetts", "GREECE: Athens", "ITALY: Florence",
		"NETHERLANDS: Landgraaf"}
	testSlice = LocationCleaner(testSlice)
	if !reflect.DeepEqual(testSlice, correctSlice) {
		t.Errorf("\ngot:\n\"%v\"\nbut wanted: \n\"%v\"", testSlice, correctSlice)
	}
}

func TestCleanLocation(t *testing.T) {
	testLocation := "buenos_aires-argentina"
	correctLocation := "ARGENTINA: Buenos Aires"
	testLocation = CleanLocation(testLocation)
	if !reflect.DeepEqual(testLocation, correctLocation) {
		t.Errorf("\ngot \"%v\", but wanted \"%v\"", testLocation, correctLocation)
	}
}

func TestDateFormat(t *testing.T) {
	dateSlice := []string{"14-12-2019", "15-12-2019", "*15-12-2019"}
	fixedDates := []string{DateFormat(dateSlice[0]), DateFormat(dateSlice[1]), DateFormat(dateSlice[2])}
	correctDates := []string{"14 December 2019", "15 December 2019", "15 December 2019"}
	if !reflect.DeepEqual(fixedDates, correctDates) {
		t.Errorf("\ngot:\n\"%v\", \nbut wanted:\n\"%v\"", fixedDates, correctDates)
	}
}

func TestRelationFormat(t *testing.T) {
	testMap := make(map[string][]string)
	dateSlice := []string{"29-02-2020", "01-03-2020", "*06-07-2021", "08-11-2021"}
	correctDateSlice := []string{"29 February 2020", "1 March 2020", "6 July 2021", "8 November 2021"}
	testMap["dallas-usa"] = dateSlice
	editedMap := RelationFormat(testMap)
	if _, exists := editedMap["USA: Dallas"]; !exists {
		t.Errorf("error in map key:\n%v", editedMap)
	} else if !reflect.DeepEqual(editedMap["USA: Dallas"], correctDateSlice) {
		t.Errorf("error in map values, wanted:\n%v \ngot:\n%v", correctDateSlice, editedMap["Usa: Dallas"])
	}
}
