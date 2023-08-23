package internal
type Structure interface {
	returnPostion() Position
}

type HomeBase struct {
	Position Position
}
func (h HomeBase) returnPosition() Position{
	return h.Position
}

type Obstacle struct {
	Position Position
}

func (o Obstacle) returnPosition() Position{
	return o.Position
}
type EnergySource struct {
	Position Position
}
func (e EnergySource) returnPosition() Position{
	return e.Position
}
