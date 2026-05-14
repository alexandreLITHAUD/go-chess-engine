package position

type Coordinates struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}

type Offset struct {
	Dx int8 `json:"dx"`
	Dy int8 `json:"dy"`
}
