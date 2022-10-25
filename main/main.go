package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))

	data, err = json.Marshal(mo)
	fmt.Println(string(data))

	j := JTest{Aa: 1, Bb: "hhh", Cc: "ppp"}
	data, err = json.Marshal(j)
	fmt.Println(string(data))
}

type JTest struct {
	Aa int    `json:"aa"`
	Bb string `json:"bb"`
	Cc string `json:"cc"`
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var mo = Movie{Title: "Casablanca", Year: 1942, Color: false,
	Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}
