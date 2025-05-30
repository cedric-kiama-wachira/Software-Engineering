package main

import (
	"fmt"
)
const (
	fullName       string  = "Cedric K. Wachira"			// Full name
        referToMeAs    string  = "Mr.Cedric"				// Preferenc on how to be addressed
	favoriteQuote  string  = "Code is life"				// Favorite quote
	favoriteEmoji  rune    = '🚀'					// Favorite emoji
	favoriteTeam   string    = "Real Madrid F.C"			// Favorite Football Crest

)

var (
	totalFriends      int          // Total friends
	dailyCoffees      uint8        // Daily coffee cups
	totalBooksRead    uint32       // Books read lifetime
	isMorningPerson   bool         // Morning person?
)

type LifeInSummary struct {
    TotalFriends    int
    DailyCoffees    uint8
    TotalBooksRead  uint32
    IsMorningPerson bool
}

type LifeInDetail struct {
    FirstFriendsNameInitials  string
    //SecondFriendsNameInitials string
    //ThirdFriendsNameInitials  string
    //FourthFriendsNameInitials string
    //FifthFriendsNameInitials  string
    //SixthFriendsNameInitials  string
    //DailyCoffeeName   	      string
    //FavoriteBookName          string
    //MyMorningStartsAt         string
}

func (lifeInDetail LifeInDetail) getLifeInDetail() string{
	return  lifeInDetail.FirstFriendsNameInitials
}
func main() {
	lifeInSummary := LifeInSummary{
		TotalFriends: 6,
		DailyCoffees: 2,
		TotalBooksRead: 20,
		IsMorningPerson: true,
	}

	lifeInDetail := LifeInDetail{
		FirstFriendsNameInitials:  "P.Mu",
		//SecondFriendsNameInitials: "D.Gi",
		//ThirdFriendsNameInitials:  "D.Du",
		//FourthFriendsNameInitials: "B.Ol",
		//FifthFriendsNameInitials:  "F.Sq",
		//SixthFriendsNameInitials:  "L.Ja",
		//DailyCoffeeName:   	  "Americano no sugar",    
		//FavoriteBookName:          "African Bible",
		//MyMorningStartsAt:         "4 AM that's everyday",
	}

        friendsFirstNameLastNameInitialsCombined := map[string]string{
		"P": "Mu",
		"L": "Ja",
		"F": "Sq",
		"D": "Du",
		"C": "De",
		"B": "Ol",
	}

	friendsFirstNameLastNameInitialsWithId:= make(map[string]int)

	friendsFirstNameLastNameInitialsWithId["P.Mu"] = 1
	friendsFirstNameLastNameInitialsWithId["D.Gi"] = 2
	friendsFirstNameLastNameInitialsWithId["D.Du"] = 3
	friendsFirstNameLastNameInitialsWithId["L.Ja"] = 4
	friendsFirstNameLastNameInitialsWithId["F.Sq"] = 5
	friendsFirstNameLastNameInitialsWithId["B.Ol"] = 6

	fmt.Printf("Welcome to my friend's app, my names are: %v \n", fullName)
	fmt.Printf("You can call me: %v\n", referToMeAs)
	fmt.Printf("My Favorite quote is: %v\n", favoriteQuote)
	fmt.Printf("I am a football fan and, my favorite team is: %v\n", favoriteTeam )
	fmt.Printf("I am always asking when: %c\n", favoriteEmoji)
	fmt.Printf("This will be improved: %s\n", lifeInDetail.getLifeInDetail())
        fmt.Printf("Combined First and Last Name Initials of friends %v\n", friendsFirstNameLastNameInitialsCombined)
	fmt.Printf("Friends First.Last Names Initials with an ID: %+v\n", friendsFirstNameLastNameInitialsWithId)
	fmt.Printf("This sums up who I am: %+v\n", lifeInSummary)

}
