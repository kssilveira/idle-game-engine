package game

type Resource struct {
	Name string
	Quantity int
	Capacity int
}

type Action struct {
	Name string
	Add []Resource
}

type Game struct {
	Resources []*Resource
	ResourceToIndex map[string]int
	Actions []Action
}

func NewGame() *Game {
	return &Game{
		ResourceToIndex: map[string]int{},
	}
}

func (g *Game) AddResource(r *Resource) {
	g.ResourceToIndex[r.Name] = len(g.Resources)
	g.Resources = append(g.Resources, r)
}

