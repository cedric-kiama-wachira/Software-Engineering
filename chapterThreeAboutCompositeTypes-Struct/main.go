package main

import (
	"time"
)
const (
	fullName       string  = "Cedric K. Wachira" // Full name
	favoriteQuote  string  = "Code is life"      // Favorite quote
	initial        byte    = 'C'                 // First initial
	favoriteEmoji  rune    = '🚀'                 // Favorite emoji
)

var (
	birthYear         int     = 1982                          // Year of birth
	age               int     = time.Now().Year() - birthYear // Dynamic age calculation
	totalFriends      int     = 150                           // Total friends
	dailyCoffees      uint8   = 2                             // Daily coffee cups
	stepsPerDay       uint16  = 8000                          // Daily steps
	totalBooksRead    uint32  = 250                           // Books read lifetime
	lifetimeSeconds   int64   = int64(time.Since(time.Date(birthYear, 1, 1, 0, 0, 0, 0, time.UTC)).Seconds()) // Calculated seconds lived
	stepsLifetime     uint64  = 103660000                     // Lifetime steps
	height            float64 = 1.75                          // Height in meters (use float64 for better precision)
	weight            float64 = 82.563                        // Weight in kg
	favoriteComplex   complex128 = 3 + 4i                    // Favorite complex number
	isMorningPerson   bool    = true                         // Morning person?
)

type CedricLife struct {
    FullName        string
    FavoriteQuote   string
    Initial         byte
    FavoriteEmoji   rune
    BirthYear       int
    Age             int
    TotalFriends    int
    DailyCoffees    uint8
    StepsPerDay     uint16
    TotalBooksRead  uint32
    LifetimeSeconds int64
    StepsLifetime   uint64    
    Height          float64
    Weight          float64
    FavoriteComplex complex128
    IsMorningPerson bool
}

func main() {
   

}


