package datagather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ArtistsData struct {
	Id               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	CFirstAlbum      string
	FirstAlbum       string `json:"firstAlbum"`
	RelationAPI      string `json:"relations"`
	LocationAPI      string `json:"locations"`
	DatesAPI         string `json:"dates"`
	ConcertLocations []string
	Concerts         map[string][]string
}
type RelationData struct {
	ID       int                 `json:"id"`
	Concerts map[string][]string `json:"datesLocations"`
}
type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

/*
newData is a global struct where data gathered from the RESTful-API
is recorded
*/
var NewData []ArtistsData

/*
GetData takes as input an API link in the form of a string,
and returns all data retreived as a slice of bytes.
*/
func getData(link string) ([]byte, error) {
	dataResponse, getErr := http.Get(link)
	// fmt.Println(dataResponse.StatusCode)
	if getErr != nil {
		return nil, getErr
	}
	dataByte, readErr := io.ReadAll(dataResponse.Body)
	if readErr != nil {
		return nil, readErr
	}
	return dataByte, nil
}

/*
populateRelationData takes a relationData struct, slice of ArtistsData structs and an index integer
as inputs, and returns the slice of ArtistsData structs with populated relation data retrieved from the
Relation API, along with an error value, which is non-nil if errors are encountered.
*/
func populateRelationData(r RelationData, artists []ArtistsData, index int) ([]ArtistsData, error) {
	artistRelation, relationErr := getData(artists[index].RelationAPI)
	if relationErr != nil {
		return artists, relationErr
	}
	json.Unmarshal(artistRelation, &r)
	artists[index].Concerts = RelationFormat(r.Concerts)
	return artists, nil
}

/*
populateLocationData takes a locationData struct, slice of ArtistsData structs and an index integer
as inputs, and returns the slice of ArtistsData structs with populated location data retrieved from the
Location API, along with an error value, which is non-nil if errors are encountered.
*/
func populateLocationData(l LocationData, artists []ArtistsData, index int) ([]ArtistsData, error) {
	artistLocation, locationErr := (getData(artists[index].LocationAPI))
	if locationErr != nil {
		return artists, locationErr
	}
	json.Unmarshal(artistLocation, &l)
	artists[index].ConcertLocations = LocationCleaner(l.Locations)
	artists[index].CFirstAlbum = DateFormat(artists[index].FirstAlbum)
	return artists, nil
}

/*
GatherDataUp takes as input an API link in the form of a string,
and uses the slice of bytes returned by GetData to populate an
ArtistsData data-structure which is then returned.
*/
func gatherDataUp(link string) ([]ArtistsData, error) {
	Artists := []ArtistsData{}
	var relationErr error
	var locationErr error
	r := RelationData{}
	l := LocationData{}

	// GET and initial unmarshall of API data
	macroData, macroErr := getData(link)
	if macroErr != nil {
		return Artists, macroErr
	}
	err := json.Unmarshal(macroData, &Artists)
	if err != nil {
		return nil, err
	}

	// Print progress to terminal
	unicodeBlock := '\u2588'
	fmt.Print("   ... (╯°□°）╯︵ ┻━┻ Just busy gathering data ...\n")
	fmt.Print(strings.Repeat("-", len(Artists)) + "\n")

	for i := 0; i < len(Artists); i++ {
		// Reset structs
		r = RelationData{}
		l = LocationData{}

		// GET and Unmarshall relation data
		Artists, relationErr = populateRelationData(r, Artists, i)
		if relationErr != nil {
			return Artists, relationErr
		}

		// GET and Unmarshall location data
		Artists, locationErr = populateLocationData(l, Artists, i)
		if locationErr != nil {
			return Artists, locationErr
		}

		fmt.Printf("%c", unicodeBlock)
	}
	fmt.Println()
	fmt.Printf("Current Data:%v Artists\n", len(Artists))
	return Artists, nil
}

/*
SaveData calls GatherDataUp (which in its turn calls GetData) and
saves the retreived and formatted datat from the API to the global
NewData variable (an []ArtistsData struct)
*/
func SaveData(api string) error {
	var gatherErr error
	NewData, gatherErr = gatherDataUp(api)
	if NewData == nil || gatherErr != nil {
		return fmt.Errorf("failed to save data from API:\n%w", gatherErr)
	}
	return nil
}
