/*
This file holds all calculation like sums, calculation of "magic number"
*/
package main

import "time"

// calcRate Calculates the so-called "Magic Number"
// The daily amount of money you can spend for a signle fixed expense/income
func calcRate(trans Transaction) float64 {
	magicAdd := 0.0
	switch trans.Recurrence {
	case "monthly":
		magicAdd = (trans.Amount * 12) / daysInYear(time.Now().Year())
	case "yearly":
		magicAdd = trans.Amount / daysInYear(time.Now().Year())
	case "twice a year":
		magicAdd = (trans.Amount * 2) / daysInYear(time.Now().Year())
	case "quarterly":
		magicAdd = (trans.Amount * 4) / daysInYear(time.Now().Year())
	}
	switch trans.Income {
	case true:
		return magicAdd
	case false:
		magicAdd = -magicAdd
	}
	return magicAdd
}

// Helper to calculate the amount of days in a month
func daysInMonth(y int, m time.Month) int {
	return time.Date(y, m, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1).Day()
}

// Gets the amount of days in a single year
func daysInYear(year int) float64 {
	numdays := 0.0
	for month := time.January; month <= time.December; month++ {
		numdays += float64(daysInMonth(year, month))
	}
	return numdays
}

// Calculates the total expenses per period
func expensesPerPeriod(period string, channel chan float64) {
	expenses := totalExpenses(db, period)
	magicNumber := baseMagic(db)
	var total float64
	switch period {
	case "week":
		total = (magicNumber * 7) + expenses
	case "month":
		total = (magicNumber * float64(daysInMonth(time.Now().Year(), time.Now().Month()))) + expenses
	case "year":
		total = (magicNumber * daysInYear(time.Now().Year())) + expenses
	}
	channel <- total
}

func percentages(total, transam float64) float64 {
	if transam < 0 {
		transam *= -1
	}
	return (transam * 100) / total
}
