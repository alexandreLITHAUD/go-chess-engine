package position

type Move struct {
	Destination Coordinates
	IsCapture   bool
	IsCheck     bool
}
