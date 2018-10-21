package main

// Best import
import (
	"example/stringutil"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// PageContext super stuff
type PageContext struct {
	Date string
	Time string
}

// MatchDay Composition
type MatchDay struct {
	Name string
}

// Main
func main() {
	http.HandleFunc("/", HomePage)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// HomePage Handler to MVP Page
func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()               // find the time right now
	homePageContext := PageContext{ //store the date and time in a struct
		Date: now.Format("MM DDDD"),
		Time: now.Format("hh:mm"),
	}

	fmt.Println(stringutil.Reverse("Hello"))

	t, err := template.ParseFiles("1n2.gohtml") //parse the html file 1n2.html
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, homePageContext) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
