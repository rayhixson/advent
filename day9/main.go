package main

import "fmt"

type Marble struct {
	Value   int
	Current bool
	Next    *Marble
	Prev    *Marble
}

func ZeroMarble() *Marble {
	return &Marble{
		Value:   0,
		Current: true,
		Next:    nil,
		Prev:    nil,
	}
}

func (m *Marble) Add(v int) *Marble {
	//fmt.Println("Adding:", v)
	new := &Marble{
		Value:   v,
		Current: true,
		Prev:    m,
	}

	if m.Next != nil {
		new.Next = m.Next
		m.Next.Prev = new
	} else {
		new.Next = m
	}
	m.Next = new

	if m.Prev == nil {
		m.Prev = new
	}

	//m.Debug()
	//new.Debug()
	return new
}

func (m *Marble) Remove() (next *Marble) {
	if m.Next == nil {
		panic(fmt.Sprintf("Can't remove: %d", m.Value))
	}

	m.Prev.Next = m.Next
	m.Next.Prev = m.Prev
	return m.Next
}

func (m *Marble) FindCurrent() *Marble {
	if m.Current {
		return m
	}
	return m.Next.FindCurrent()
}

func (m Marble) String() string {
	if m.Current {
		return fmt.Sprintf("(%d) ", m.Value)
	}
	return fmt.Sprintf("%d ", m.Value)
}

func (m Marble) Debug() {
	fmt.Printf("%d <- %d -> %d == Current: %v\n", m.Prev.Value, m.Value, m.Next.Value, m.Current)
}

type Player struct {
	ID    int
	Score int
}

type Players []*Player

type GameBoard struct {
	CurrentTurn int
	LastMarble  int
	MarbleValue int
	ZeroMarble  *Marble
	People      Players
}

func NewGameBoard(players, lastMarble int) GameBoard {
	g := GameBoard{LastMarble: lastMarble}
	g.ZeroMarble = ZeroMarble()

	for i := 0; i < players; i++ {
		g.People = append(g.People, &Player{ID: i + 1})
	}
	return g
}

func (g GameBoard) Dump(playerID int) {
	s := fmt.Sprintf("[%d] ", playerID)

	for c := g.ZeroMarble; c.Next != nil; c = c.Next {
		s += c.String()

		if c.Next == g.ZeroMarble {
			break
		}
	}
	fmt.Println(s)
}

func (g GameBoard) HighScore() int {
	h := 0
	for _, p := range g.People {
		if p.Score > h {
			h = p.Score
		}
	}
	return h
}

func (g *GameBoard) PlaceMarble(value int, player *Player) {
	c := g.ZeroMarble.FindCurrent()
	c.Current = false

	if value%23 == 0 {
		//fmt.Printf("mod 23: %d\n", value)
		player.Score += value

		// remove the marble 7 counter clockwise

		for i := 0; i < 7; i++ {
			c = c.Prev
		}

		c.Remove().Current = true
		player.Score += c.Value
	} else {
		if c.Next != nil {
			c.Next.Add(value)
		} else {
			c.Add(value)
		}
	}
}

func (g GameBoard) CurrentPlayer() *Player {
	if g.CurrentTurn == 0 {
		return g.People[0]
	}
	player := (g.CurrentTurn - 1) % len(g.People)
	// but zero indexed
	return g.People[player]
}

func (g *GameBoard) Play() {
	//g.Dump(0)
	for i := 0; i < g.LastMarble; i++ {
		g.CurrentTurn++
		p := g.CurrentPlayer()
		g.PlaceMarble(g.CurrentTurn, p)
		//g.Dump(p.ID)
	}
}

func main() {
	g := NewGameBoard(439, 71307)
	g.Play()
	fmt.Println("High Score:", g.HighScore())
}

const data = `439 players; last marble is worth 71307 points`
