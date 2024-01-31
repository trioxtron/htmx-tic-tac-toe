package gameengine

import "fmt"

type square string

var board [3][3]square

type Game struct {
    board [3][3]square
    turn int
}

func (g *Game) GetBoard() [3][3]square {
    return g.board
}

func (g *Game) SetField(y, x int) string {
    if g.board[y][x] != "" {
        return "This field is already taken"
    } else if g.turn % 2 != 0 {
        g.board[y][x] = "x"
        g.turn++
        fmt.Println(g.board)
        return "x"
    } else {
        g.board[y][x] = "o"
        g.turn++
        fmt.Println(g.board)
        return "o"
    }
}

func (g *Game) ClearField() {
    for y, r := range g.board {
        for x, _ := range r {
            fmt.Println(g.board[y][x])
            g.board[y][x] = ""
        }
    }
}
