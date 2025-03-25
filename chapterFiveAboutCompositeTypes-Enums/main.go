package main

import "fmt"

type TeamId int 
const(
	Arsenal TeamId = 1 
	AstonVilla = iota + 1
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

type FriendId int
const(
	PMu FriendId = 1
	CDe = iota + 1
	DDu
	LJa
	FSq
	BOl
	WaD
	MbE
)

type SupportStatus int
const(
	NoTeam = iota
	SupportedTeam 
)

func main(){
	fmt.Println("Testing")
	fmt.Println(Arsenal,"=Id for Arsenal,", ManchesterUnited, "=Id for ManchesterUnited,",WestHamUnited, "=Id for WestHamUnited",WolverhamptonWanderers, "=Id for WolverhamptonWanderers", PMu, "=Id for PMu user,", CDe, "=Id for CDe user,", WaD,"=Id for WaD user, and",MbE,"=Id for MbE user.")

	eplTeamWithStadiumName:= map[string]string{
		"Arsenal":"Emirates Stadium",
		"AstonVilla":"Villa Park",
		"AFCBournemouth":"Vitality Stadium",
		"Brentford":"Gtech Community Stadium",
		"BrightonAndHoveAlbion":"American Express Stadium",
		"Chelsea":"Stamford Bridge",
		"CrystalPalace":"Selhurst Park",
		"Everton":"Goodison Park",
		"Fulham":"Craven Cottage",
		"IpswichTown":"Portman Road",
		"LeicesterCity":"King Power Stadium",
		"Liverpool":"Anfield",
		"ManchesterCity":"Etihad Stadium",
		"ManchesterUnited":"Old Trafford",
		"NewcastleUnited":"St James Park",
		"NottinghamForest":"City Ground",
		"Southampton":"St Mary's Stadium",
		"TottenhamHotspur":"Tottenham Hotspur Stadium",
		"WestHamUnited":"London Stadium",
		"WolverhamptonWanderers":"Molineux Stadium",
	}

	eplTeamStadiumNameAndLocation:= map[string]string{
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

}
