package day9

import "testing"

type Game struct {
	Players int
	currentMarble int
	currentPlayer int
	board []int
	Scores map[int]int
	nextMarble int
}

func (g *Game) Init() {
	g.Scores = make(map[int]int)
	for i := 1;i <= g.Players; i++ {
		g.Scores[i] = 0
	}

	g.currentMarble = 0
	g.board = make([]int,0)
	g.board = append(g.board, 0)
	g.currentPlayer = 1
	g.nextMarble = 1
}

func (g *Game) next() {
	if g.currentPlayer == g.Players {
		g.currentPlayer = 1
	} else {
		g.currentPlayer += 1
	}

	g.nextMarble += 1

}

func (g *Game) Tick() {
	if g.nextMarble == 1{
		g.board = append(g.board, g.nextMarble)
		g.currentMarble = g.nextMarble
	} else if g.nextMarble % 23 == 0 {
		//First, the current player keeps the marble they would have placed
		g.Scores[g.currentPlayer] += g.nextMarble
		//In addition, the marble 7 marbles counter-clockwise from the current marble is removed from the circle
		// and also added to the current player's score.
		removedMarblePosition := (g.currentMarble - 7 )
		if removedMarblePosition < 0 {
			removedMarblePosition = len(g.board) + removedMarblePosition
		}
		g.Scores[g.currentPlayer] += g.board[removedMarblePosition]
		g.board = append(g.board[:removedMarblePosition], g.board[removedMarblePosition+1:]...)
		g.currentMarble = removedMarblePosition
	} else {
		nextPos := (g.currentMarble + 2 ) % len(g.board)
		first := g.board[:nextPos]
		last := make([]int,len(g.board[nextPos:]))
		copy(last, g.board[nextPos:])
		first = append(first, g.nextMarble)
		g.board = append(first, last...)
		g.currentMarble = nextPos
	}

	g.next()

}

func (g *Game) GetHighestScore() (player int, score int) {
	player = 0
	score = 0
	for i:=1; i <= g.Players; i++ {
		if score < g.Scores[i] {
			score = g.Scores[i]
			player = i
		}
	}
	return player,score
}

func NewGame(numberOfPlayers int) *Game {
	g := Game{Players: numberOfPlayers}
	g.Init()
	return &g
}

func Test_PartOneSampleSimulation(t *testing.T){
	testCases := []struct{
		NumberOfPlayers int
		LastMarble int
		HighScore int
		Winner int
	}{
		{
			9,
			25,
			32,
			5,
		}, {
			429,
			70901,
			399645,
			392,
		},
		{
			429,
			7090100,
			3352507536,
			12,
		},
	}

	for _, tc := range testCases {
		g := NewGame(tc.NumberOfPlayers)
		for i := 0; i < tc.LastMarble; i++ {
			g.Tick()
		}
		player, score := g.GetHighestScore()
		if player != tc.Winner || score != tc.HighScore {
			t.Errorf("Expected Player: %d, Score %d, Got: %d and %d", tc.Winner, tc.HighScore, player, score)
		}
	}
}