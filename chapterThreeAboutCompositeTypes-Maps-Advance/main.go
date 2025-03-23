package main

import (
	"fmt"
)
const (
	fullName       string  = "Cedric K. Wachira"			// Full name
        referToMeAs    string  = "Mr.Cedric"				// Preferenc on how to be addressed
	favoriteQuote  string  = "Code is life"				// Favorite quote
	favoriteEmoji  rune    = '🚀'					// Favorite emoji
	myFavoriteTeam   string    = "Real Madrid F.C"			// Favorite football team name
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


func main() {
	lifeInSummary := LifeInSummary{
		TotalFriends: 6,
		DailyCoffees: 2,
		TotalBooksRead: 20,
		IsMorningPerson: true,
	}

        friendsFirstNameLastNameInitialsCombined := map[string]string{
		"P": "Mu",
		"L": "Ja",
		"F": "Sq",
		"D": "Du",
		"C": "De",
		"B": "Ol",
	}
	
	friendsIdMappedToTheFirstNameLastNameInitials:= make(map[int]string)

	friendsIdMappedToTheFirstNameLastNameInitials[10] = "P.Mu"
	friendsIdMappedToTheFirstNameLastNameInitials[20] = "C.De"
	friendsIdMappedToTheFirstNameLastNameInitials[30] = "D.Du"
	friendsIdMappedToTheFirstNameLastNameInitials[40] = "L.Ja"
	friendsIdMappedToTheFirstNameLastNameInitials[50] = "F.Sq"
	friendsIdMappedToTheFirstNameLastNameInitials[60] = "B.Ol"
	
	englishTeamNameAndStadiumName:= map[string]string{
		"Arsenal":"Emirates Stadium",
		"Aston Villa":"Villa Park",
		"AFC Bournemouth":"Vitality Stadium",
		"Brentford":"Gtech Community Stadium",
		"Brighton & Hove Albion":"American Express Stadium",
		"Chelsea":"Stamford Bridge",
		"Crystal Palace":"Selhurst Park",
		"Everton":"Goodison Park",
		"Fulham":"Craven Cottage",
		"Ipswich Town":"Portman Road",
		"Leicester City":"King Power Stadium",
		"Liverpool":"Anfield",
		"Manchester City":"Etihad Stadium",
		"Manchester United":"Old Trafford",
		"Newcastle United":"St James Park",
		"Nottingham Forest":"City Ground",
		"Southampton":"St Mary's Stadium",
		"Tottenham Hotspur":"Tottenham Hotspur Stadium",
		"West Ham United":"London Stadium",
		"Wolverhampton Wanderers":"Molineux Stadium",
	}
	englishTeamStadiumNameAndLocation:= map[string]string{
		"Emirates Stadium":"London",
		"Villa Park":"Birmingham",
		"Vitality Stadium":"Bournemouth",
		"Gtech Community Stadium":"London",
		"American Express Stadium":"Falmer",
		"Stamford Bridge":"London",
		"Selhurst Park":"London",
		"Goodison Park":"Liverpool",
		"Craven Cottage":"London",
		"Portman Road":"Ipswich",
		"King Power Stadium":"Leicester",
		"Anfield":"Liverpool",
		"Etihad Stadium":"Manchester",
		"Old Trafford":"Manchester",
		"St James Park":"Newcastle upon Tyne",
		"City Ground":"West Bridgford",
		"St Mary's Stadium":"Southampton",
		"Hotspur Tottenham Hotspur Stadium":"London",
		"United London Stadium":"London",
		"Wanderers Molineux Stadium":"Wolverhampton",
	}

friendNameInitialsWithTheTeamTheySupport:= map[string]string{
		"P.Mu":"Manchester United",
		"D.Du":"Arsenal",
		"F.Sq":"Manchester United",
	}
friendNameInitialsWithoutATeamToSupport:= map[string]string{
		"C.De":"",
		"L.Ja":"",
		"B.Ol":"",
	}
teamIdWithTeamName:= map[int]string{
		1:"Arsenal",
		2:"Aston Villa",
		3:"AFC Bournemouth",
		4:"Brentford",
		5:"Brighton & Hove Albion",
		6:"Chelsea",
		7:"Crystal Palace",
		8:"Everton",
		9:"Fulham",
		10:"Ipswich Town",
		11:"Leicester City",
		12:"Liverpool",
		13:"Manchester City",
		14:"Manchester United",
		15:"Newcastle United",
		16:"Nottingham Forest",
		17:"Southampton",
		18:"Tottenham Hotspur",
		19:"West Ham United",
		20:"Wolverhampton Wanderers",
}

	fmt.Printf("Welcome to my friend's app, my names are: %v \n", fullName)
	fmt.Printf("You can call me: %v\n", referToMeAs)
	fmt.Printf("My Favorite quote is: %v\n", favoriteQuote)
	fmt.Printf("I am a football fan and, my favorite team is: %v\n", myFavoriteTeam )
	fmt.Printf("I am always asking when: %c\n", favoriteEmoji)
        fmt.Printf("Combined First and Last Name Initials of friends %v\n", friendsFirstNameLastNameInitialsCombined)
	fmt.Printf("Friends ID mapped to First.Last Names: %+v\n", friendsIdMappedToTheFirstNameLastNameInitials)
	fmt.Printf("These friends have teams that they support and they are: %+v\n", friendNameInitialsWithTheTeamTheySupport)
	fmt.Printf("These friends don't have teams that they support and they are: %+v\n", friendNameInitialsWithoutATeamToSupport)
        fmt.Printf("EPL Team and Home Stadium: %+v\n", englishTeamNameAndStadiumName)
	fmt.Printf("EPL Team and City: %+v\n", englishTeamStadiumNameAndLocation) 
	fmt.Printf("Team Id and Name: %+v\n", teamIdWithTeamName)
	fmt.Printf("This sums up who I am: %+v\n", lifeInSummary)

}
