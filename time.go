package main

//Go offers extensive support for times and durations

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	//We’ll start by getting the current time.

	now := time.Now()

	p(now)

	then := time.Date(

		//we can build a time struct by providing the year, month, day, etc.
		//Times are always associated with a Location, i.e. time zone.

		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)

	//we can extract the various components of the time value as expected.

	p(then.Year())

	p(then.Month())

	p(then.Day())

	p(then.Hour())

	p(then.Minute())

	p(then.Second())

	p(then.Nanosecond())

	p(then.Location())

	//The Monday-Sunday Weekday is also available.

	p(then.Weekday())

	//These methods compare two times, testing if the first occurs before, after,
	// or at the same time as the second, respectively.

	p(then.Before(now))

	p(then.After(now))

	p(then.Equal(now))

	//The Sub methods returns a Duration representing the interval between two times.

	diff := now.Sub(then)

	p(diff)

	//We can compute the length of the duration in various units.

	p(diff.Hours())

	p(diff.Minutes())

	p(diff.Seconds())

	p(diff.Nanoseconds())

	p(then.Add(diff))

	p(then.Add(-diff))

	//You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.

}

// output :
//2012-10-31 15:50:13.793654 +0000 UTC
//2009-11-17 20:34:58.651387237 +0000 UTC
//2009
//November
//17
//20
//34
//58
//651387237
//UTC
//Tuesday
//true
//false
//false
//25891h15m15.142266763s
//25891.25420618521
//1.5534752523711128e+06
//9.320851514226677e+07
//93208515142266763
//2022-05-16 22:50:13.793654 +0000 UTC
//2021-12-05 01:19:43.509120474 +0000 UTC
