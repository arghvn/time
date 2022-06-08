package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

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
