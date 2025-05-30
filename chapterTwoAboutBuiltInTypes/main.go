package main

import (
	"fmt"
	"time"
	"unsafe"
)

const (
	fullName       string  = "Cedric K. Wachira" // Full name
	favoriteQuote  string  = "Code is life"      // Favorite quote
	favoriteHour   uint8   = 7                   // Favorite hour (7 AM)
	initial        byte    = 'C'                 // First initial
	favoriteEmoji  rune    = '🚀'                 // Favorite emoji
)

var (
	birthYear         int     = 1982                          // Year of birth
	age              int     = time.Now().Year() - birthYear // Dynamic age calculation
	totalFriends      int     = 150                           // Total friends
	dailyCoffees      uint8   = 2                             // Daily coffee cups
	stepsPerDay       uint16  = 8000                          // Daily steps
	totalBooksRead    uint32  = 250                           // Books read lifetime
	lifetimeSeconds   int64   = int64(time.Since(time.Date(birthYear, 1, 1, 0, 0, 0, 0, time.UTC)).Seconds()) // Calculated seconds lived
	daysSinceLastTrip uint16  = 180                           // Days since last trip
	yearlyExpenses    uint32  = 45000                         // Yearly expenses ($)
	stepsLifetime     uint64  = 103660000                     // Lifetime steps
	starsVisible      uint64  = 3000                          // Stars visible at night
	height            float64 = 1.75                          // Height in meters (use float64 for better precision)
	weight            float64 = 82.563                        // Weight in kg
	favoriteComplex   complex128 = 3 + 4i                    // Favorite complex number
	lifeVector        complex128 = complex(float64(age), float64(birthYear)) // Life as a complex number
	isMorningPerson   bool    = true                         // Morning person?

	memoryAddress     uintptr = uintptr(unsafe.Pointer(&age)) // Memory address of age
	lastError         error   = nil                           // Last encountered error
)

func main() {
	fmt.Println("Full Name:", fullName)
	fmt.Println("Age:", age)
	fmt.Println("Birth Year:", birthYear)
	fmt.Println("Total Friends:", totalFriends)
	fmt.Println("Daily Coffees:", dailyCoffees)
	fmt.Println("Steps Per Day:", stepsPerDay)
	fmt.Println("Total Books Read:", totalBooksRead)
	fmt.Println("Lifetime Seconds:", lifetimeSeconds)
	fmt.Println("Days Since Last Trip:", daysSinceLastTrip)
	fmt.Println("Yearly Expenses:", yearlyExpenses)
	fmt.Println("Steps Lifetime:", stepsLifetime)
	fmt.Println("Stars Visible at Night:", starsVisible)
	fmt.Println("Height (m):", height)
	fmt.Println("Weight (kg):", weight)
	fmt.Println("Favorite Complex Number:", favoriteComplex)
	fmt.Println("Life Vector:", lifeVector)
	fmt.Println("Is Morning Person?:", isMorningPerson)
	fmt.Println("Favorite Quote:", favoriteQuote)
	fmt.Println("Initial:", string(initial))
	fmt.Println("Favorite Emoji:", string(favoriteEmoji))
	fmt.Println("Memory Address of Age:", memoryAddress)
	fmt.Println("Last Error:", lastError)
}
