package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// define a structure.

type Entry struct {
	Name string 
	Surname string
	Tel string
}

// define global variable

var data = []Entry{}
var MIN = 0
var MAX = 26

// serach functionality.

func search(key string) *Entry{
	for i,v := range data{
		if v.Tel == key{
			return &data[i]
		}
	}

	return nil
}

// list functionality.

func list(){
	for _,v := range data{
		fmt.Println(v)
	}
}

func randomPhoneNumberGen(min, max int) int {
    return rand.Intn(max-min) + min
}
 


func getString(len int64) string {
	startChar := "A"
    temp := ""
    var i int64 = 1
    for {
        myRand := randomPhoneNumberGen(MIN,MAX)
        newChar := string(startChar[0] + byte(myRand))
        temp = temp + newChar
        if i == len {
            break
        }
        i++
    }
    return temp
}


func populate(n int, s []Entry) {
    for i := 0; i < n; i++ {
        name := getString(4)
        surname := getString(5)
        n := strconv.Itoa(randomPhoneNumberGen(100, 199))
        data = append(data, Entry{name, surname, n})
    }
}

func generateBytes(n int64) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}



func main(){

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Usage: search|list <arguments>")
		return
	}
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	// How many records to insert
	n := 100
	populate(n, data)
	fmt.Printf("Data has %d entries.\n", len(data))

	data = append(data, Entry{"Messi", "the-Goat", "2204001010"})
	data = append(data, Entry{"Kimbo", "slice", "1205005050"})
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
    data = append(data, Entry{"Mary", "Doe", "2109416871"})
    data = append(data, Entry{"John", "Black", "2109416123"})

	// differentiate between commands.
	switch arguments[1]{
	case "search":
		if len(arguments) != 3{
			fmt.Println("Usage: Search Surname")
			return
		}
		result := search(arguments[2])

		if result == nil {
			fmt.Println("Entry can't be found:", arguments[2])
			return
		}
		fmt.Println(*result)

		// list commands.
	case "list":
		list()
		// response to anything that is not a match.
	default:
		fmt.Println("Not a valid option, sorry.")
	}


}