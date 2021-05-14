package towers

const (
	TowerHeight int = 6
	TowerWidth  int = 30
)

// Tower is a struct to describe a tower in Hanoi Tower game
type Tower struct {
	Name  string
	Pos   int
	Bases []int
}
