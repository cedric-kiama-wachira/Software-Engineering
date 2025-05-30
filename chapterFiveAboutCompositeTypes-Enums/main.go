package main

import "fmt"

type TeamId int 
const(
	Arsenal = iota + 1 
	AstonVilla
	AFCBournemouth
	Brentford
	BrightonAndHoveAlbion
	Chelsea
	CrystalPalace
	Everton
	Fulham
	IpswichTown
	LeicesterCity
	Liverpool
	ManchesterCity
	ManchesterUnited
	NewcastleUnited
	NottinghamForest
	Southampton
	TottenhamHotspur
	WestHamUnited
	WolverhamptonWanderers
)

var teamNames = map[TeamId]string{
	Arsenal: "Arsenal",
	AstonVilla: "Aston Villa",
	AFCBournemouth: "AFC Bournemouth",
	Brentford: "Brentford",
	BrightonAndHoveAlbion: "Brighton & Hove Albion",
	Chelsea: "Chelsea",
	CrystalPalace: "Crystal Palace",
	Everton: "Everton",
	Fulham: "Fulham",
	IpswichTown: "Ipswich Town",
	LeicesterCity: "Leicester City",
	Liverpool: "Liverpool",
	ManchesterCity: "Manchester City",
	ManchesterUnited: "Manchester United",
	NewcastleUnited: "Newcastle United",
	NottinghamForest: "Nottingham Forest",
	Southampton: "Southampton",
	TottenhamHotspur: "Tottenham Hotspur",
	WestHamUnited: "WestHam United",
	WolverhamptonWanderers: "Wolverhampton Wanderers",
}

var eplTeamWithStadiumName = map[TeamId]string{
	Arsenal:"Emirates Stadium",
	AstonVilla:"Villa Park",
	AFCBournemouth:"Vitality Stadium",
	Brentford:"Gtech Community Stadium",
	BrightonAndHoveAlbion:"American Express Stadium",
	Chelsea:"Stamford Bridge",
	CrystalPalace:"Selhurst Park",
	Everton:"Goodison Park",
	Fulham:"Craven Cottage",
	IpswichTown:"Portman Road",
	LeicesterCity:"King Power Stadium",
	Liverpool:"Anfield",
	ManchesterCity:"Etihad Stadium",
	ManchesterUnited:"Old Trafford",
	NewcastleUnited:"St James Park",
	NottinghamForest:"City Ground",
	Southampton:"St Mary's Stadium",
	TottenhamHotspur:"Tottenham Hotspur Stadium",
	WestHamUnited:"London Stadium",
	WolverhamptonWanderers:"Molineux Stadium",
}

var eplStadiumNameAndLocation = map[string]string{
	"Emirates Stadium":"London.",
	"Villa Park":"Birmingham.",
	"Vitality Stadium":"Bournemouth.",
	"Gtech Community Stadium":"London.",
	"American Express Stadium":"Falmer.",
	"Stamford Bridge":"London.",
	"Selhurst Park":"London.",
	"Goodison Park":"Liverpool.",
	"Craven Cottage":"London.",
	"Portman Road":"Ipswich.",
	"King Power Stadium":"Leicester.",
	"Anfield":"Liverpool.",
	"Etihad Stadium":"Manchester.",
	"Old Trafford":"Manchester.",
	"St James Park":"Newcastle.",
	"City Ground":"West Bridgford.",
	"St Mary's Stadium":"Southampton.",
	"Tottenham Hotspur Stadium":"London.",
	"London Stadium":"London.",
	"Molineux Stadium":"Wolverhampton.",
}

type FriendId int
const(
	PMu = iota + 1
	CDe 
	DDu
	LJa
	FSq
	BOl
	WaD
	MbE
)

var friendNames = map[FriendId]string{
	PMu: "PMu",
	CDe: "CDe",
	DDu: "DDu",
	LJa: "LJa",
	FSq: "FSq",
	BOl: "BOl",
	WaD: "WaD",
	MbE: "MbE",
}

type SupportStatus int
const(
	NoTeam SupportStatus = iota
	SupportedTeam 
)

var friendSupport = map[FriendId]struct {
	Status SupportStatus
	Team   TeamId
	}{
		PMu:{SupportedTeam, ManchesterUnited},
		CDe:{NoTeam,0},
		DDu:{SupportedTeam, Arsenal},
		LJa:{NoTeam,0},
		FSq:{SupportedTeam, ManchesterUnited},
		BOl:{NoTeam,0},
		WaD:{SupportedTeam, Arsenal},
		MbE:{SupportedTeam, ManchesterUnited},
	}

func (t TeamId) String() string{
	return teamNames[t]
}

func (f FriendId) String() string{
	return friendNames[f]
}

func (s SupportStatus) String() string{
	switch s{
	case NoTeam:
	return "No Team"
	case SupportedTeam:
	return "Supports a team"
	default:
	panic("Team does not exist.")
	}
}


func main(){
	fmt.Println("Testing\n")
	
	fmt.Println("Let's get the Team Name mapped to the Team ID\n")
	//fmt.Println(TeamId(Arsenal),"team ID is", Arsenal,"\n",
	fmt.Println(Arsenal,"is the ID of Team", TeamId(Arsenal),"\n",
		    ManchesterUnited,"is the ID of Team", TeamId(ManchesterUnited),"\n",
		    WestHamUnited,"is the ID of team", TeamId(WestHamUnited),"\n",
		    WolverhamptonWanderers,"is the ID of team",TeamId(WolverhamptonWanderers),"\n")

	fmt.Println("Testing done, it's time for some logic, I'll pair a team, it's stadium and location...\n")
	for team, stadium := range eplTeamWithStadiumName{
		location := eplStadiumNameAndLocation[stadium]
		fmt.Printf("%s plays at %s, located in %s\n", team, stadium, location)
	}

	fmt.Println("Let's see if the IDs assigned are properly mapped to my friends correctly\n")
	fmt.Println(PMu,"= Id for my friend PMu.\n",
		    CDe,"= Id for my friend CDe.\n",
		    WaD,"= Id for my friend WaD.\n",
		    "and",MbE,"= Id for my friend MbE.\n")

	fmt.Println("I'll now check which teams my friends support\n")
	for friend, support := range friendSupport {
		if support.Status == SupportedTeam {
			fmt.Printf("%s supports %s\n", friend, support.Team)
		} else {
			fmt.Printf("%s does not support any team.\n", friend)
		}
	}
}
