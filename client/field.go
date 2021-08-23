package client

import (
	"strings"

	"github.com/pascallohrer/fillergame/api"
)

type field struct {
	tiles [100]api.Color
}

/*/ <-- toggle between symbol mode and color mode by adding/removing a / at the start of this line

var symbols = map[api.Color]string{
	api.Color_RED:     "°",
	api.Color_GREEN:   "-",
	api.Color_BLUE:    "|",
	api.Color_YELLOW:  "%",
	api.Color_CYAN:    "#",
	api.Color_MAGENTA: "+",
	api.Color_ORANGE:  "X",
	api.Color_PURPLE:  "*",
}

var colors = map[string]api.Color{
	"°": api.Color_RED,
	"-": api.Color_GREEN,
	"|": api.Color_BLUE,
	"%": api.Color_YELLOW,
	"#": api.Color_CYAN,
	"+": api.Color_MAGENTA,
	"X": api.Color_ORANGE,
	"*": api.Color_PURPLE,
}

/*/

var symbols = map[api.Color]string{
	api.Color_RED:     "\x1b[101m  \x1b[0m",
	api.Color_GREEN:   "\x1b[102m  \x1b[0m",
	api.Color_BLUE:    "\x1b[104m  \x1b[0m",
	api.Color_YELLOW:  "\x1b[103m  \x1b[0m",
	api.Color_CYAN:    "\x1b[106m  \x1b[0m",
	api.Color_MAGENTA: "\x1b[105m  \x1b[0m",
	api.Color_ORANGE:  "\x1b[43m  \x1b[0m",
	api.Color_PURPLE:  "\x1b[100m  \x1b[0m",
}

var colors = map[string]api.Color{
	"1": api.Color_RED,
	"2": api.Color_GREEN,
	"3": api.Color_BLUE,
	"4": api.Color_YELLOW,
	"5": api.Color_CYAN,
	"6": api.Color_MAGENTA,
	"7": api.Color_ORANGE,
	"8": api.Color_PURPLE,
}

/**/

// String creates a representation of the field's status fit for printing in a command line environment
func (f *field) String() string {
	lines := make([]string, 10, 10)
	for t := 9; t >= 0; t-- {
		line := make([]string, 10, 10)
		for i := 0; i <= 9; i++ {
			line[i] = symbols[f.tiles[10*t+i]]
		}
		lines[9-t] = strings.Join(line, "")
	}
	return " " + strings.Join(lines, "\n ")
}

func (f *field) update(tileID int, newColor api.Color) {
	f.tiles[tileID] = newColor
}

func (f *field) menu() string {
	menuItems := make([]string, 0, len(colors))
	for input, color := range colors {
		if color == f.tiles[0] || color == f.tiles[99] {
			continue // the player's and opponent's current color are invalid
		}
		menuItems = append(menuItems, input+symbols[color])
	}
	return " " + strings.Join(menuItems, " ")
}
