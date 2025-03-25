package main

import (
	"fmt"
	"sort"
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
    InitialFriendsList    int
    DailyCoffees    uint8
    TotalBooksRead  uint32
    IsMorningPerson bool
    UpdatedFriendsList int
}


func main() {

	fmt.Printf("Welcome to my friend's app, my names are: %v \n", fullName)
	fmt.Printf("You can call me: %v\n", referToMeAs)
	fmt.Printf("My Favorite quote is: %v\n", favoriteQuote)
	fmt.Printf("I am a football fan and, my favorite team is: %v\n", myFavoriteTeam )
	fmt.Printf("I am always asking when: %c\n", favoriteEmoji)

	fmt.Printf("Adding more data points, I want to show my list of friends with teams that they support \n")

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

	fmt.Printf("EPL Team and Home Stadium: %+v\n", englishTeamNameAndStadiumName)

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

	fmt.Printf("EPL Team and City: %+v\n", englishTeamStadiumNameAndLocation)

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

        fmt.Printf("Team Id and Name: %+v\n", teamIdWithTeamName)

        friendsFirstNameLastNameInitialsCombined := map[string]string{
		"P": "Mu",
		"L": "Ja",
		"F": "Sq",
		"D": "Du",
		"C": "De",
		"B": "Ol",
	}
	
	fmt.Printf("Combined First and Last Name Initials of friends %v\n", friendsFirstNameLastNameInitialsCombined)

	friendsIdMappedToTheFirstNameLastNameInitials:= make(map[int]string)
	friendsIdMappedToTheFirstNameLastNameInitials[1] = "P.Mu"
	friendsIdMappedToTheFirstNameLastNameInitials[2] = "C.De"
	friendsIdMappedToTheFirstNameLastNameInitials[3] = "D.Du"
	friendsIdMappedToTheFirstNameLastNameInitials[4] = "L.Ja"
	friendsIdMappedToTheFirstNameLastNameInitials[5] = "F.Sq"
	friendsIdMappedToTheFirstNameLastNameInitials[6] = "B.Ol"
	
	fmt.Printf("Initial MAP Range for Friends ID mapped to First.Last Names: \n")

	for firstKeyK, firstValueV := range friendsIdMappedToTheFirstNameLastNameInitials{
		fmt.Printf("The ID is : %d and the Name Initials are : %s\n", firstKeyK, firstValueV)
	} 
	
	lifeInSummary := LifeInSummary{
		InitialFriendsList: 6,
		DailyCoffees: 2,
		TotalBooksRead: 20,
		IsMorningPerson: true,
		UpdatedFriendsList: 8,
	}
	
	fmt.Printf("This sums up who I am: %+v\n", lifeInSummary)
        
	fmt.Printf("Lets Polish this further, I figured some friends have teams that they support while others don't. \n")

	friendsIdMappedToTheFirstNameLastNameInitials[7] = "Wa.D"
	friendsIdMappedToTheFirstNameLastNameInitials[8] = "Mb.E"

	fmt.Printf("Added two new friends to the list, who I know have a team that they support, now we have eight friends with and ID: %+v\n", friendsIdMappedToTheFirstNameLastNameInitials)

        fmt.Printf("Mapping all that have a team they support, leaving out those that don't. \n")

	friendNameInitialsWithTheTeamTheySupport:= make(map[string]string)
		friendNameInitialsWithTheTeamTheySupport["P.Mu"]="Manchester United"
		friendNameInitialsWithTheTeamTheySupport["D.Du"]="Arsenal"
		friendNameInitialsWithTheTeamTheySupport["F.Sq"]="Manchester United"
		friendNameInitialsWithTheTeamTheySupport["Wa.D"]="Arsenal"
		friendNameInitialsWithTheTeamTheySupport["Mb.E"]="Manchester United"

	for secondKeyK, secondValueV := range friendNameInitialsWithTheTeamTheySupport{
		fmt.Printf("The name Initials are : %s and the team they support are : %s\n", secondKeyK, secondValueV)
	}	

	friendNameInitialsWithoutATeamTheySupport:= map[string]string{
		"C.De":"",
		"L.Ja":"",
		"B.Ol":"",
	}

	fmt.Printf("I want to show friends without a team that they support before deleting their respective variables, here is the list %+v\n", friendNameInitialsWithoutATeamTheySupport)
	delete(friendNameInitialsWithoutATeamTheySupport, "C.De")
	delete(friendNameInitialsWithoutATeamTheySupport, "L.Ja")
	delete(friendNameInitialsWithoutATeamTheySupport, "B.Ol")
	
	fmt.Printf("Deleted the 'containers/variables' for the three friends without a team that they support, they are empty %+v\n", friendNameInitialsWithoutATeamTheySupport)

	fmt.Printf("Adding slices to my friendsIdMappedToTheFirstNameLastNameInitials map \n")
	fmt.Printf("Before slicing the list looked like this: %+v\n", friendsIdMappedToTheFirstNameLastNameInitials)
        
        //friendsIdSliced := []int{1,2,3,4,5,6,7,8}

	//fmt.Println("Intro to slice types, the hard coded way: \n",friendsIdSliced)

	//fmt.Println("Fixing the data disconnect as nothing has been achieved by our Initial slice attempt\n")

	var friendInitials []string
	for _, initials := range friendsIdMappedToTheFirstNameLastNameInitials {
		friendInitials = append(friendInitials, initials)
	}

	var supporters []string
	for initials := range friendNameInitialsWithTheTeamTheySupport {
		supporters = append(supporters, initials)
	}

	var teams []string
	seenTeams := make(map[string]bool)
	for _, team := range friendNameInitialsWithTheTeamTheySupport {
		if !seenTeams[team] {
			teams = append(teams, team)
			seenTeams[team] = true
		}
	}
        
	sort.Strings(friendInitials)
	fmt.Println("All Friends Initials:", friendInitials)
	sort.Strings(supporters)
	fmt.Println("Friends Who Support a Team:", supporters)
	fmt.Println("Distinct Teams Supported:", teams)
}
