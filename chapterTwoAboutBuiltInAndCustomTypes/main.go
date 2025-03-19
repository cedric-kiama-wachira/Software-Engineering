package main

import "fmt"

const fullnames string = "Cedric K. Wachira"       // Full name
var birthyear int = 1982                           // Year of birth
var age int = 2025 - 1982                          // Age in 2025 (43)
var totalFriends int = 150                         // Total friends
var dailyCoffees int8 = 2                          // Daily coffee cups
var stepsPerDay int16 = 8000                       // Daily steps
var totalBooksRead int32 = 250                     // Books read lifetime
var lifetimeSeconds int64 = 1356048000             // Seconds lived by 2025
var daysSinceLastVacation uint = 180               // Days since last trip
var favoriteHour uint8 = 7                         // Favorite hour (7 AM)
var yearlyExpenses uint16 = 45000                  // Yearly expenses ($)
var stepsLifetime uint32 = 103660000               // Lifetime steps
var starsVisible uint64 = 3000                     // Stars visible at night
var height float32 = 1.75                          // Height in meters
var weight float64 = 82.563                        // Weight in kg
var favoriteComplex complex64 = 3 + 4i             // Favorite complex number
var lifeVector complex128 = 43.5 + 1982.0i         // Life as a complex number
var isMorningPerson bool = true                    // Morning person?
var favoriteQuote string = "Code is life"          // Favorite quote
var initial byte = 'C'                             // First initial
var favoriteEmoji rune = '🚀'                      // Favorite emoji

func main(){

fmt.Println(birthyear)
fmt.Println(favoriteEmoji)

}
