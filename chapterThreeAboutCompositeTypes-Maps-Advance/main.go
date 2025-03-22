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

type LifeInDetail struct {
    FirstFriendsNameInitials  string
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

	friendsIdMappedToTheFirstNameLastNameInitials[1] = "P.Mu"
	friendsIdMappedToTheFirstNameLastNameInitials[2] = "C.De"
	friendsIdMappedToTheFirstNameLastNameInitials[3] = "D.Du"
	friendsIdMappedToTheFirstNameLastNameInitials[4] = "L.Ja"
	friendsIdMappedToTheFirstNameLastNameInitials[5] = "F.Sq"
	friendsIdMappedToTheFirstNameLastNameInitials[6] = "B.Ol"
	
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
		"C.De":"",
		"D.Du":"Arsenal",
		"L.Ja":"",
		"F.Sq":"Manchester United",
		"B.Ol":"",
	}

teamIdWithTeamName:= map[int]string{
		10:"Arsenal",
		20:"Aston Villa",
		30:"AFC Bournemouth",
		40:"Brentford",
		50:"Brighton & Hove Albion",
		60:"Chelsea",
		70:"Crystal Palace",
		80:"Everton",
		90:"Fulham",
		100:"Ipswich Town",
		110:"Leicester City",
		120:"Liverpool",
		130:"Manchester City",
		140:"Manchester United",
		150:"Newcastle United",
		160:"Nottingham Forest",
		170:"Southampton",
		180:"Tottenham Hotspur",
		190:"West Ham United",
		200:"Wolverhampton Wanderers",
}

	fmt.Printf("Welcome to my friend's app, my names are: %v \n", fullName)
	fmt.Printf("You can call me: %v\n", referToMeAs)
	fmt.Printf("My Favorite quote is: %v\n", favoriteQuote)
	fmt.Printf("I am a football fan and, my favorite team is: %v\n", myFavoriteTeam )
	fmt.Printf("I am always asking when: %c\n", favoriteEmoji)
	fmt.Printf("This will be improved: %s\n", lifeInDetail.getLifeInDetail())
        fmt.Printf("Combined First and Last Name Initials of friends %v\n", friendsFirstNameLastNameInitialsCombined)
	fmt.Printf("Friends ID mapped to First.Last Names: %+v\n", friendsIdMappedToTheFirstNameLastNameInitials)
        fmt.Printf("EPL Team and Home Stadium: %+v\n", englishTeamNameAndStadiumName)
	fmt.Printf("EPL Team and City: %+v\n", englishTeamStadiumNameAndLocation) 
	fmt.Printf("My Friend and the team they support: %+v\n", friendNameInitialsWithTheTeamTheySupport) 
	fmt.Printf("Team Id and Name: %+v\n", teamIdWithTeamName)
	fmt.Printf("This sums up who I am: %+v\n", lifeInSummary)

}
