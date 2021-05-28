package models

type Controller struct {
	IsLeftMouseDown bool
	LeftMouseStartX int
	LeftMouseStartY int

	IsRightMouseDown bool
	RightMouseStartX int
	RightMouseStartY int
}
