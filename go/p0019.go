/**
Problem 19: Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

// While it would be tempting to implement Zeller's Congruence
// or even something exotic like Tomohiko Sakamoto's Algorithm,
// it is probably more fruitful to get some practice with golang's time library.

import "time"

func SundaysOnFirstOfMonth(begin_year, end_year int) int {
	count := 0

	for year := begin_year; year <= end_year; year++ {
		for month := 1; time.Month(month) <= time.December; month++ {
			date := time.Date(year, time.Month(month), 1,
				// wake up wake up wake up, it's the first of the month...
				4, 20, 69, 808, time.UTC)
			if date.Weekday() == time.Sunday {
				// ...oh wait it's Sunday, we can sleep in some
				count += 1
			}
		}
	}

	return count
}
