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

func (t TeamId) String() string{
	return teamNames[t]
}

func (f FriendId) String() string{
	return friendNames[f]
}

func main(){
	fmt.Println("Testing\n")
	fmt.Println(Arsenal,"= Id for Arsenal.\n",ManchesterUnited,"= Id for Manchester United.\n",WestHamUnited,"= Id for WestHam United.\n",WolverhamptonWanderers,"= Id for Wolverhampton Wanderers.\n",PMu,"= Id for PMu user.\n",CDe,"= Id for CDe user.\n",WaD,"= Id for WaD user.\n and",MbE,"= Id for MbE user.\n")
	fmt.Println("Testing Done, time for some logic...\n")

	for team, stadium := range eplTeamWithStadiumName{
		location := eplStadiumNameAndLocation[stadium]
		fmt.Printf("%s plays at %s, located in %s\n", team, stadium, location)
	}
}
