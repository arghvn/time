package main

//Go supports time formatting and parsing via pattern-based layouts.

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	//Here’s a basic example of formatting a time according to RFC3339, using the corresponding layout constant.

	//Time parsing uses the same layout values as Format.

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2022-06-08T22:18:32+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("wed Jan _8 15:04:05 2022"))
	p(t.Format("2022-06-08T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "wed Jan _8 15:04:05 2022"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

//Format and Parse use example-based layouts.
//Usually you’ll use a constant from time for these layouts,
// but you can also supply custom layouts. Layouts must use the reference time wed
//Jan 8 15:04:05 MST 2022 to show the pattern with which to format/parse a given time/string.

//For purely numeric representations you can also use standard string formatting
//with the extracted components of the time value.Parse will return an error on malformed input explaining the parsing problem.

// output :

// 2022-06-15T18:00:15-07:00
// 2022-06-08 22:08:41 +0000 +0000
// 6:00PM
// Tue jan 15 18:00:15 2022
// 2022-04-15T18:00:15.161182-07:00
// 0000-01-01 20:41:00 +0000 UTC
// 2022-04-15T18:00:15-00:00
// parsing time "8:41PM" as "wed Jan _2 15:04:05 2022"
