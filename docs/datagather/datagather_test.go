package datagather

import (
	"testing"
)

func TestGetData(t *testing.T) {
	// Only tests the valid API from the task
	// TO DO: Add more falsifiable tests
	testLink := "https://groupietrackers.herokuapp.com/api/artists"
	returnData, err := getData(testLink)
	if err != nil {
		t.Errorf("error returned when calling GetData function:\n%v", err)
	} else if returnData == nil {
		t.Errorf("no data retrieved when calling API")
	}
}

func TestGatherDataUp(t *testing.T) {
	// Only tests the valid API from the task
	// TO DO: Add more falsifiable tests
	testLink := "https://groupietrackers.herokuapp.com/api/artists"
	returnData, err := gatherDataUp(testLink)
	if err != nil {
		t.Errorf("error encountered while gathering and unmarshalling data:\n%v", err)
	} else if len(returnData) != 52 {
		// As of 16/11/2022, the number of artists listed on the API was 52
		t.Errorf("error encountered while gathering and unmarshalling data:\n"+
			"number of artists retrieved: %v, number of artists expected: 52", len(returnData))
	}
}

func TestSaveData(t *testing.T) {
	testLink := "https://groupietrackers.herokuapp.com/api/artists"
	err := SaveData(testLink)
	if err != nil {
		t.Errorf("error encountered while saving API data:\n%v", err)
	}
}

// func TestPadWhiteSpace(t *testing.T) {
// 	textToPad := "Move me 5 characters to the left"
// 	padding := 5
// 	result := padWhiteSpace(textToPad, padding)
// 	expected := "     Move me 5 characters to the left"
// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("\ngot:\n\"%v\"\nbut wanted:\n\"%v\"", result, expected)
// 	}
// }

// func TestCenterPad(t *testing.T) {
// 	testText1 := "A whole\nlotta rubbish\n"
// 	testText2 := "A whole lotta"
// 	testText3 := "A whole"
// 	center1 := centerPad(testText1, 8)
// 	fmt.Println(center1)
// 	center2 := centerPad(testText2, 8)
// 	fmt.Println(center2)
// 	center3 := centerPad(testText3, 20)
// 	fmt.Println(center3)
// 	if len(center1) != 2 {
// 		t.Errorf("\ngot:\n\"%v\" \"%v\" \"%v\"", center1, center2, center3)
// 	}
// }

/*
TO DO: Find a way to test getTermLength
Currently, the test package does not seem to access stdin / stdout in a regular way
*/
// func TestGetTermLength(t *testing.T) {
// 	currentTerminalLength := getTermLength()
// 	if currentTerminalLength <= 0 {
// 		t.Errorf("Current terminal length is incorrectly shown as:\n\"%v\"", currentTerminalLength)
// 	}
// }
