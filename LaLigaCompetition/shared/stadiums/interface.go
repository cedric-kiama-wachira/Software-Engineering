type StadiumOwner interface {
	GetStadium() Stadium
	ValidateCapacity() error
}
