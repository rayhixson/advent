package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Turn int

const (
	LeftTurn  Turn = 1
	RightTurn Turn = 2
	Straight  Turn = 3
)

func parse(data string) *Grid {
	g := Grid{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	row := 0
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			continue
		}

		for col, c := range val {
			switch c {
			case '^', 'v', '<', '>':
				// this is a cart
				g.AllCarts = append(g.AllCarts, NewCart(col, row, c))
			}

		}
		g.Tracks = append(g.Tracks, val)
		row++
	}
	return &g
}

func main() {
	/*
		g := parse(data)

		ticks := 1000
		click, a, _ := g.Run(ticks)
		if a == nil {
			fmt.Println("DIdn't find collider in n ticks:", ticks)
			return
		}

		fmt.Println("Collision at:", a.X, a.Y, click)
	*/
	g := parse(data)

	ticks := 20000
	click, a := g.Run2(ticks)

	if a != nil {
		fmt.Println("Last cart:", a)
	} else {
		fmt.Println("Nope last after ticks:", click)
		fmt.Println("Carts left:", g.AllCarts)
	}

}

const data = `
                /-------------------------------------------------\                                                    /---------------------------\  
 /--------------+---------------------------------------\         |                                                    |                           |  
 |          /---+-------->------------------------------+-->-----\|                                           /--------+---\                       |  
 |          |   |            /---------------------\    |        ||                                           |        |   |                       |  
 |          |   |            |                     |    |/-------++-------------------------------------------+--------+---+-----------------------+-\
 |          |   |            |                     |/---++-------++-------------------------------------------+--------+---+-----------------------+\|
 |          |   |       /----+---------------------++---++-------++----------------\                         /+--------+---+-------\               |||
 |          |   |  /----+----+-------------\       ||   ||       ||                |                         ||        |   |       |               |||
 |          |   |  |    |   /+--<----------+-\/----++---++-------++----------------+-------------------------++--------+---+-----\ |               |||
 |          |   |  |    |   ||             | ||    ||   ||       ||                |               /---------++-------\|   |   /-+-+------------\  |||
 |          |   |  |    |   ||        /----+-++----++---++-------++----\           |               |         ||       ||   |   | | |            |  |||
 |          |   |  |    |   ||        |    | ||    ||   ||       ||    |           |               |         ||       ||   |   | | |            |  |||
 |          |   |  |    |   ||        |    | ||    ||  /++-------++----+-----------+---------------+---------++-------++---+---+-+-+---\        |  |||
 |          |   |  |    |   ||        |    | ||    ||  |||     /-++----+-------\   |               |         ||       ||   |   | | |   |        |  |||
 |   /------+---+--+----+---++--------+-\  | ||    ||  |||     | ||    |       |   |               |         ||       ||   |   | | |   |        |  |||
 |   |      |   |  |    |   ||        | |  | ||    |\--+++-----+-++----+-------+---+---------------+---------++-------++---+---+-+-+---+--------+--+/|
 |   |      |   |  |    |   ||        | |  |/++----+---+++-----+-++--\ |       |   |               |         ||       ||   |   | | |   |        |  | |
 |   |      |   |  |    |   \+--------+-+--++/|    |   |||     | ||  | |       |   |       /-------+\        ||       ||   |   | | |   v        |  | |
 |   |      |   |  |    |    |        |/+--++-+----+---+++-----+-++--+-+-------+---+-------+-------++--------++-------++---+---+-+-+\  |        |  | |
 |/--+------+---+--+----+--\ |    /---+++--++-+----+---+++-----+-++\ | |       |   |       |       ||        ||       ||   |   | | ||  |        |  | |
 ||  |      |   |  |    |  | |    |   |||  || |    | /-+++-----+-+++-+-+-------+---+-------+\      ||        ||       ||   |   | | ||  |        |  | |
 ||  |      |   |  |    |  | |    |   |||  || |    | | |||     | ||| | |       |   |  /----++------++--------++-------++---+---+-+-++--+---\    |  | |
 ||  |      |   |  |    |  | |    |   |||  || |    | | |||     | ||| | |       |   |  |    ||      ||        |\-------++---/   | | ||  |   |    |  | |
 ||  |      |   |  |    |  | |/---+---+++--++-+----+-+-+++-----+-+++-+-+-------+---+--+----++------++--------+--------++\      | | ||  |   |    |  | |
 ||  |/-----+---+-<+----+--+-++---+---+++--++-+----+-+-+++----\| ||| | |  /----+---+--+----++------++--------+--------+++------+-+-++--+---+----+--+\|
 ||  ||     |   |  |    |  | ||   |   |||  || |    | | |||    || ||| | |  |    |   |  |    ||      ||        |        |||      | | ||  |   |    |  |||
 ||  \+-----+---+--+----+--+-++---+---++/  || |    | | |||    || ||| | |  |   /+---+--+----++------++--------+-\      |||      | | ||  |   |    |  |||
 ||   |     |   |  |    |/-+-++---+---++---++-+-\  | | |||    || ||| | |  |   ||   |  |    ||      || /------+-+------+++----\ | | ||  |   |    |  |||
 ||   |     |   |  |    || | ||   |   || /-++-+-+--+-+-+++----++-+++-+-+--+---++---+--+----++------++-+------+-+------+++----+-+-+-++--+---+\   |  |||
 \+---+-----+---+--+----++-+-++---+---++-+-++-+-+--+-+-+/|    || ||| | |  |   ||   |  |    ||      || |      | |   /--+++----+-+-+-++--+---++---+\ |||
  \---+-----+--<+--+----++-/ ||   |   || | |\-+-+--+-+-+-+----++-+++-/ |  |   ||   |  |    ||      || |      | |   |  |||    | | | ||  |   ||   || |||
      |     |   |  |    ||   ||   |/--++-+-+--+-+--+-+-+-+----++-+++---+--+---++---+--+----++------++-+------+-+---+--+++---\| | | ||  |   ||   || |||
      |     |   |  |    ||   ||   ||  \+-+-+--+-+--+-+<+-+----++-+++---/  |   ||   |  |    ||/-----++-+------+-+---+-\|||   || | | ||/-+---++--\|| |||
      |     |   |  |    ||   ||  /++---+-+-+--+-+--+-+-+-+----++-+++------+---++--\|  |    \++-----+/ |      | |   | ||||   || | | ||| |   ||  ||| |||
      |    /+---+--+----++---++--+++---+-+-+--+-+--+-+-+-+----++-+++------+---++--++--+-----++--\  |  |      | |   | ||||   || | | ||| |   ||  ||| |||
      | /--++---+--+----++---++--+++---+-+-+--+\|  | | | |/---++-+++------+---++--++--+-----++-\|  |  |      | |   | ||||   || | | ||| |   ||  ||| |||
    /-+-+--++---+--+----++---++--+++---+-+-+--+++--+-+-+-++---++-+++------+---++--++-\|    /++-++--+--+------+-+\  | ||||   || | | ||| |   ||  ||| |||
    | | |  ||   |  |/---++---++--+++---+-+-+--+++--+-+-+-++---++-+++------+---++--++-++----+++-++--+--+------+-++--+-++++--\|| | | ||| |   ||  ||| |||
    | | |  ||/--+--++---++---++--+++---+-+-+--+++--+-+-+-++---++-+++------+---++--++-++----+++-++--+--+------+-++--+-++++\ ||| | | ||| |   ||  ||| |||
/---+-+-+--+++--+--++\  ||   ||  |||   | | |  ||| /+-+-+-++---++-+++------+---++--++-++----+++-++--+\ |/-----+-++--+\||||| ||| | | ||| |   ||  ||| |||
|   | | |  |||  |  |||  ||   ||  |||   |/+-+--+++-++-+-+-++---++-+++------+---++--++-++----+++-++--++-++-----+\||  ||||||| ||| | | ||| |   ||  ||| |||
|   | | |  |||  |  |||  ||   ||  |||   ||| |  ||| || | | ||   || |||      |   ||  || ||    |||/++--++-++-----++++--+++++++-+++-+-+-+++-+\  ||  ||| |||
|   | | |  |||  |  |||  ||  /++--+++--\||| |  ||| || | | ||   || |||      |   ||  || ||    ||||||  \+-++-----++++--+++/\++-+++-+-+-+++-++--++--+++-/||
|   | | |  |||  |  |||  ||  |||  |||  |||| |  ||| || | | ||   || |||      |   ||  || ||    ||||||   | ||     ||||  |||  || ||| | | ||| ||  ||  |||  ||
|   | | |  |||  |  |||  |\--+++--+++--++++-+--++/ || | | ||   || |||      |   ||  || ||    ||||||   | ||     ||||  |||  || ||| | | ||| ||  ||  |||  ||
|   | |/+--+++--+--+++--+---+++--+++--++++-+--++--++-+\| ||   || |||      |   ||  || |\----++++++---+-++-----++++--+++--++-+++-+-+-+++-++--/|  |||  ||
|   | |||  |||  |  |\+--+---+++--+++--++++-+--++--++-+++-++---++-+++------+---++--++-+-----++++++---+-++-----++++--+++--++-/||/+-+-+++-++---+--+++-\||
|   | |||  |||  |  | |/-+---+++--+++--++++-+--++--++-+++-++---++-+++------+---++--++-+\    ||||||   | ||     ||||  |||  ||  |||| | ||| ||   |  ||| |||
|   | ||| /+++--+--+-++-+---+++--+++--++++-+--++--++-+++-++---++-+++--\   |   ||  || ||    ||||||   | ||     ||||  \++--++--++++-+-+++-++---+--++/ |||
|/--+-+++-++++--+--+-++-+---+++--+++--++++-+--++--++-+++-++\  || |||  |   \---++--++-++----++++++---+-++-----++++---++--++--++++-+-+++-++---+--++--+/|
|| /+-+++-++++--+--+-++-+---+++--+++--++++-+--++--++-+++-+++--++-+++--+-------++-\|| ||    |||||| /-+-++-----++++---++-\||  |||| | ||| ||   |  ||  | |
|| || ||| ||||  |  | || |   \++--+++--/||| |  ||  || ||| |||  || |||  |       || ||| ||    |||||| |/+-++-----++++---++-+++--++++-+-+++\||   |  ||  | |
|| || ||| ||||  |  | || |    ||  \++---+++-+--++--++-+++-+++--++-+++--+-------++-+/| ||    |||||| ||| ||     ||||   || |||  |||| | ||||||   |  ||  | |
|| || ||| |\++--+--+-++-+----++---++---+++-+--++--++-+++-+++--++-+++--+-------++-+-+-++----+++++/ ||| ||     ||||   || |||  |||| | ||||||   |  ||  | |
|| || ||| | ||  |  | || |    ||   ||   ||| |  ||  ||/+++-+++--++-+++--+-------++-+-+-++----+++++--+++-++-\   ||||   || |||  |||| | ||||||   |  ||  | |
|| || ||| | ||  |  | || |    ||   ||   ||| |  \+--++++++-+++--++-+++--+-------++-+-+-++----+++++--+++-++-+---++++---++-+++--++++-/ ||||||   |  ||  | |
|| || ||| | ||  |  | ||/+----++---++---+++-+---+--++++++-+++--++-+++--+------\|| | | ||    ||||| /+++-++-+---++++-\ || |||  ||||   ||||||   |  ||  | |
|| || ||| | || /+--+-++++\   ||   ||  /+++-+---+--++++++-+++--++-+++--+------+++-+-+-++----+++++-++++-++-+---++++\| || |||  ||||   ||||||   |  ||  | |
|| || ||| | || || /+-+++++---++---++--++++-+---+--++++++-+++--++-+++--+------+++-+-+-++----+++++-++++-++-+-\ |||||| || |||  ||||   ||||||   |  ||  | |
|| || \++-+-++-++-++-+++++---++---++--++++-+---+--++++++-+++--/| |||  |      ||| | | ||    ||||| ||||/++-+-+-++++++-++\|||  ||||   ||||||   |  ||  | |
|| ||  || | || || || v||||   || /-++--++++-+---+--++++++-+++\  | |||  |      ||| | | ||    ||||| ||||||\-+-+-++++++-/|||||  ||||   ||||||   |  ||  | |
|| ||  || | || || || |||||   || | ||  |||| |   |  |||||| ||||  | |||  |      ||| | | ||    ||||| ||||||  | | ||||||  |||||  ||||   ||||||   |  ||  | |
|| ||  || | || || || |||||   || | ||  |||| |   |  |||||| ||||  | |||  |      |||/+-+-++----+++++-++++++--+-+-++++++--+++++--++++---++++++---+-\||  | |
|| ||  || | || || || |||||   || | \+--++++-+---+--++++++-++++--+-++/  |      ||||| | ||    ||||| ||||||/-+-+-++++++-\|||||  ||||   ||||||   | |||  | |
|| ||  || | || || || |||||   || |  |  |||| | /-+--++++++-++++--+-++---+------+++++-+-++----+++++-+++++++-+-+-++++++-++++++\ ||||   ||||||   | |||  | |
|| ||  || | || || || |||||   || |  |  |||| | | |  |||||| ||||  | ||   |      ||||| | ||    ||||| ||||||| | | |||||| ||||||| ||||   ||||||   | |||  | |
|| ||  || | ||/++-++-+++++---++-+--+--++++-+-+-+--++++++-++++--+-++\  |      ||||| | || /--+++++-+++++++-+\| |||||| ||||||| ||||   ||||||   | |||  | |
|| ||  || | ||||| ||/+++++---++-+--+--++++-+-+-+--++++++-++++\ | |||/-+------+++++-+-++-+--+++++-+++++++-+++-++++++-+++++++-++++---++++++---+\|||  | |
|| ||  || | ||||| ||||||||   || |  |  |||| | | |  |||||| ||||| | |||| |      ||||| | || |  ||||| |||^||| ||| \+++++-+++++++-++++---/|||||   |||||  | |
|| ||  || | ||||| ||||||||   |\-+--+--++++-+-+-+--++++++-+++++-+-++++-+------+++++-+-++-+--+++++-+++++++-+++--+++++-++++/|| ||\+----+++++---+++++--/ |
|| ||  || | ||||| ||||||||   |/-+--+--++++-+-+-+--++++++-+++++-+-++++-+------+++++-+-++-+--+++++-+++++++-+++--+++++-++++-++-++-+----+++++-\ |||||    |
|| ||  || | |||||/++++++++---++-+--+\ |||| | | |  |||||| ||||| | |||| |    /-+++++-+\|| |  ||||| ||||||| |||  ||||| |||| || || |    ||||| | |||||    |
|| ||  || | ||||||||||||||   || |  || |||| | | |  |||||| ||||| | |||| |    | ||||| |||| |  ||||| ||||||| |||  ||||| |||| || || |    ||||| | |||||    |
|| ||  || | ||||||||||||||   || |  || |||| | | |  |||||| ||||| | |||| |    | ||||| |||| |  ||||| |||||\+-+++--+++++-++++-++-+/ |    ||||| | |||||    |
|| ||  || | ||||||||||||||   || |  \+-++++-+-+-+--++++++-+++++-+-++++-+----+-+++++-++++-+--+++++-+++++-+-+++--+++++-++++-++-/  |    ||||| | |||||    |
|| ||  || | ||||\+++++++++---++-+---+-++++-+-+-+--++++++-+++++-+-+/|| |    | ||||| |||| |  ||||| ||||| | |||  ||||| |||| ||    |    ||||| | |||||    |
|| ||  || |/++++-+++++++++-\ || |   | |||| | | |  |||\++-+++++-+-+-++-+----+-+++++-++++-+--+/||| ||||| | |||  ||||| |||| ||    |    ||||| | |||||    |
|| ||  || |||||| ||||||||| | || |   | |||| | | |  ||| || ||||| | | || |    | ||||| |||| |  | ||| ||||| | |||  ||||| |||| ||    |    ||||| | |||||    |
|| ||  || |||||| ||||||||| | || |/--+-++++-+-+-+--+++-++-+++++-+-+-++-+----+-+++++-++++-+\ | ||| ||||| | |||  ||||| |||| ||    |    ||||| | |||||    |
|| ||  || |||||| ||||||\++-+-++-++--+-++++-+-+-+--+++-++-+++++-+-+-++-+----+-/|||| |||| || | ||| ||||| | |||  ||||| |||| ||    |    ||||| | |||||    |
||/++--++-++++++-++++++-++\| || ||  | |||| | | |  ||| v| ||||| | | || |    |  |||| |||| || | ||| ||||| | ||| /+++++-++++-++----+-\  ||||| | |||||    |
|||||  || ||\+++-++++++-++++-++-++--+-++++-+-+-+--+++-++-+++++-+-/ || |    |  |||| |||| || | ||| ||||| | ||| |||||| |||| ||    | |  ||||| | |||||    |
|||||  || || ||\-++++++-+/|| || ||  | |||| | | |  ||| || ||||| |   || |    |  |||| |||| || | ||| ||||\-+-+++-++++++-++/| ||    | |  ||||| | |||||    |
|||||  || || ||  ||\+++-+-++-++-++--+-++++-//+-+--+++-++-+++++-+---++-+----+--++++-++++-++-+-+++-++++--+-+++-++++++-++-+-++----+-+\ ||||| | |||||    |
|||||  || || ||  || ||| | || || ||  | ||||  || |  ||| || ||||| |   || |    |  |||| |||| || | ||| ||||  | ||| |||||| || | ||    \-++-+++++-+-++++/    |
|||||  || || ||  |\-+++-+-++-++-++--+-++++--++-+--+++-++-+++++-+---++-+----+--++++-++++-++-+-+++-++++--+-++/ |||||| || | ||      || ||||| | ||||     |
|||||  || || ||  |  ||| | || || ||  | |||\--++-+--+++-++-+++++-+---++-+----+--++++-++++-++-+-+++-++++--+-++--++++++-++-+-++------++-+++++-+-/|||     |
|||||  || || ||  |  ||| | || || ||  | |||   || |  ||| || ||||| |   || |    |  |||| |||| || | ||| ||||  | ||  |||||| || | ||      || ||||| |  |||     |
|||||  || |\-++--+--+++-+-+/ || ||  | |||   || |  ||| || ||||| |   || |    |  |||| |||| || | ||| ||||  | ||  |||||| || | ||      || ||||| |  |||     |
|||||  || |  ||  | /+++-+-+--++-++--+-+++---++-+--+++-++-+++++-+---++-+----+--++++\|||| || \-+++-++++--+-++--+++/|| || | ||      || ||||| |  |||     |
|||||  || |  ||  | |||| | |  || ||  | |||   || |  ||| || ||||| |   || |    |  ||||||||| ||   ||| ||||  | ||  ||| || || | ||      || ||||| |  |||     |
|||||  || | /++--+-++++-+-+--++-++--+-+++---++-+--+++-++-+++++-+---++-+----+--+++++++++-++---+++-++++--+-++--+++\|| || | ||      || ||||| |  |||     |
|||||  || | |||  | |||| | |  || ||  | |||   || |  ||| || ||||| |   || |    |  ||||||||| ||   ||| ||||/-+-++--++++++-++-+-++-----\|| ||||| |  |||     |
|||||  \+-+-+++--+-++++-+-+--++-++--+-+++---++-+--+++-/| \++++-+---++-+----+--+++++++++-++---+++-+++++-+-++--++++++-++-+-++-----+++-+++++-+--+++-----/
|||||   | | |||  | |||| | |  || ||  | |||   || |  \++--+--++++-+---++-+----+--+++++++++-++---+++-+++/| \-++--++++++-/| | ||     ||| ||||| |  |||      
|||||   | | |||  | |||| | |  |\-++--+-+++---++-+---++--+--++++-+---++-+----+--+++++++++-++---+++-+++-+---++--++++++--+-+-++-----+++-+++++-/  |||      
|||||   | | |||/-+-++++-+-+--+--++--+-+++---++-+---++--+--++++\|   || |    |  ||||||||| ||   ||| ||| |   ||  ||||||  | | ||     ||| |||||    |||      
|||||   | | |||| | |||| | |  |  ||  | |||   || |   ||  |  ||||||   || |    |  ||||||||| ||   ||| ||| |   ||  ||||||  | | ||     ||| |||||    |||      
|||||   | | |||| | |||| | |  |  ||  | |||   || |   ||  |  ||||||   || |    |  ||||||||| ||   ||| |\+-+---++--++++++--+-/ ||     ||| |||||    |||      
|||||   | | |||| | |||| | |  \--++--+-+++---++-+---/|  |  ||||||   || |    |  ||||||||| ||   \++-+-+-+---++--++++++--/   ||     ||| |||||    |||      
|||||   | | |||| | |||| | |     ||  | |||   || |   /+--+--++++++---++-+----+--+++++++++-++----++-+-+-+---++--++++++------++-----+++-+++++----+++----\ 
|||||   | | |||| | |||| | |     ||  | |||   || |   ||  |  \+++++---++-+----+--+++++++++-++----+/ \-+-+---++--+++++/      ||     ||| |||||    |||    | 
||||| /-+-+-++++-+-++++-+-+-----++--+-+++---++-+---++--+-\ |||||   || |    |  ||||||||| ||    |    | |   ||/-+++++-------++-----+++-+++++\   |||    | 
||||| | | | |||| | |||| | |     ||  | |||   || |   ||  | | ||||\---++-+----+--+/||||||| ||    |    | |   ||| |||||       ||     ||| ||||||   |||    | 
||||| | | | |||| | |\++-+-+-----++--+-+++---++-+---++--+-+-++/|    || |    |  | ||||||| ||    |    | \---+++-+++++-------++-----/|| ||||||   |||    | 
||||| | | | |||| | | || | |     ||  | |||   || |   |\--+-+-++-+----++-+----+--+-+++++++-++----+----+-----/|| |||||       ||      || ||||||   |||    | 
||||| | | | |||| | | || | |     ||  | |||   || |   |   |/+-++-+----++-+-\  |  | |^||||| \+----+----+------/| |^|||       ||      || ||||||   |||    | 
||||| | | | |||| | | || | |     ||  | |||   || |   |   ||| || |    || | |  |  | |||||||  |   /+----+-----\ | |||||       ||      || ||||||   |||    | 
||||| | | | |||| | | || | |     ||  | |||   || |   |   ||| || |    || | ^  \--+-++++/||  |   ||    |  /--+-+-+++++--\    ||      || ||||||   |||    | 
||||| | | | |||| | | || \-+-----++--+-+++---++-+---+---+++-++-+----++-+-+-----+-+++/ ||  |   ||  /-+--+--+-+-+++++--+----++\     || ||||||   |||    | 
||||| | | | |||| | | ||   |     ||  | ||\---++-+---+---+++-++-+----++-+-+-----+-+++--++--+---++--+-+--+--+-+-+/|||  |    |||     || ||||||   |||    | 
||||| | | \-++++-+-+-++---+-----++--+-++----++-+---+---+++-++-+----++-/ |     | |||  ||  |   ||  | |  |  | | | |||  |    |||     || |\++++---++/    | 
||||| | |   |||| | | ||   |     ||  | |\----++-+---+---+++-++-+----++---+-----+-+++--++--+---++--+-+--+--+-+-+-+++--+----+++-----++-/ ||||   ||     | 
||||| \-+---++++-+-+-++---+-----++--+-+-----++-+---+---++/ || |    ||   |     | |||  ||  |   ||  | |  |  | \-+-+++--+----+++-----++---+++/   ||     | 
|||||   |   |||| | | ||   |     \+--+-+-----++-+---+---++--+/ |    ||   |     | |||  ||  |   ||  | |  |  |   | |||  |    |||     ||   |||    ||     | 
|||||   |   |||| | | ||   |      |  | |     || |   |   ||  |  |    ||   |/----+-+++--++--+---++--+-+--+--+---+-+++--+----+++-----++---+++--\ ||     | 
|||||   |   |||| | | ||   |      |  | |     || |   |   ||  |  |    |\---++----+-+++--++--+---++--+-+--+--+---+-+++--+----+++-----++---+++--+-/|     | 
||||\---+---++++-+-+-++---+------+--+-+-----++-+---+---++--+--+----+----++----+-+++--/|  |   ||  | \--+--+---+-+++--+----+++-----++---/||  |  |     | 
\+++----+---++++-+-+-/|   |      |  | |     || |   |   ||  |  |    |    ||    | |||   |  |   |\--+----+--+---+-+++--+----+++-----++----+/  |  |     | 
 |||    |   |||| | |  |   |      |  | |     || |   |   \+--+--+----+----++----+-+++---+--+---+---+----+--+---+-+++--+----+++-----++----/   |  |     | 
 |||    \---++++-+-+--+---+------+--+-+-----++-/   |    |  |  |    |    ||    | |||   |  |   |   |    \--+---+-+++--/  /-+++-----++----\   |  |     | 
 |||        |||| | |  |   |   /--+--+-+-----++---\ |    |  |  |    |    ||    | |||   |  |   |   \-------+---+-+++-----+-++/     ||    |   |  |     | 
 |||        |||| | |  |   |   |  |  | |     ||   | |    |  |  |    |    ||    \-+++---+--+---+-------<---+---+-/||     | ||      ||    |   |  |     | 
 ||\--------++++-+-+--+---+---+--+--+-+-----++---+-+----+--+--+----+----++------+/|   |  |   \-----------/   |  ||     | ||      ||    |   |  |     | 
 ||         |||| | |  |   |   |  |  | |     ||   | |    \--+--+----+----/|      | |   |  |                   |  ||   /-+-++------++----+\  |  |     | 
 ||         |||| | |  |   |   |  |  | |     ||   | |       |  |    |     |      | |   |  |                   |  ||   | | ||      ||    ||  |  |     | 
 ||         \+++-+-+--+---+---+--+--+-+-----++---+-+-------+--+----+-----+------+-+---+--+-------------------+--/|   | | ||      ||    ||  |  |     | 
 ||          ||| | |  \---+---+--+--+-+-----++---+-+-------+--+----+-----+------+-+---/  |                   |   |   | | ||      ||    ||  |  |     | 
 ||          ||\-+-+------+---+--+--+-+-----++---+-+-------+--/    |     |      | |      |                   |   |   | | ||      ||    ||  |  |     | 
 ||          ||  | |      |   |  |  | |     ||   | \-------+-------+-----+------+-+------+-------------------+---+---+-+-++------++----++--+--+-----/ 
 ||          ||  | |     /+---+--+--+-+-----++---+--------\|       |     |   /--+-+------+-------------------+-\ |   | | ||      ||    ||  |  |       
 ||          ||  | \-----++---+--+--+-+-----++---+--------++-------+-----+---+--+-/      |                   \-+-+---+-+-++------/|    ||  |  |       
 \+----------++--+-------++---+--+--+-+-----++---+--------+/       |     |   |  |        |                     | |   | \-++-------+----/|  |  |       
  |          \+--+-------++---+--+--+-+-----++---+--------+---->---+-----+---+--+--------+---------------------+-+---+---/|       |     |  |  |       
  |           |  |       ||   |  |  |/+-----++---+------\ |        |     \---+--+--------+---------------------+-+---+----+-------+-----+--/  |       
  |           \--+-------++---+--+--+++-----++---+------+-+--------/         |  |        |                     | |   |    |       |     |     |       
  \--------------+-------+/   |  |  |||     ||   |      | |                  |  \--------+---------------------+-+---+----+-------+-----+-----/       
            /----+-------+----+--+--+++-----++---+------+-+------------------+-----------+-------------------\ | |   |    |       |     |             
            |    |       |    |  |  |||     ||   |      | |                  |           |                   | | |   \----+-------+-----/             
            |    |       |    |  |  |||     ||   |      | |                  |           |                   | | |        |       |                   
            |    |       |    |  |  |||     |\---+------+-+------------------+-----------+-------------------+-+-+--------/       |                   
            |    |       \----+--+--+++-----+----+------+-/                  |           |                   | | |                |                   
            |    |            |  |  |||     \----+------+--------------------+-----------+-------------------+-+-+>---------------/                   
            |    |            |  |  |\+----------+------/                    |           |                   | | |                                    
            |    |            |  \-<+-+----------+---------------------------+-----------/                   | | |                                    
            |    |            |     | \----------+---------------------------+-------------------------------+-+-/                                    
            |    \------------+-----/            |                           |                               | |                                      
            |                 |                  |                           \-------------------------------+-/                                      
            \-----------------+------------------+-----------------------------------------------------------/                                        
                              \------------------/                                                                                                    `