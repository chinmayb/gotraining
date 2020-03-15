package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

const (
	Singles         = "singles"
	FINISHING_SCORE = 21
)

func NewPlayer(name string) *Player {
	return &Player{Name: name}
}

func NewGame(gt string) Game {
	return Game{
		gameType: gt,
	}
}

type Game struct {
	scoreSheet string
	gameType   string
	players    []*Player
}

type Player struct {
	Score   int
	Name    string
	court   chan ball
	quality int
}

type ball struct {
	quality int32 // 0,1,2..10
}

func generateRandom() int {
	rand2.Seed(time.Now().Unix())
	return rand2.Intn(20)
}

func (p *Player) hit(c chan ball) {
	c <- ball{quality: int32(generateRandom())}
}

func (p *Player) process(b ball) bool {
	// judge it based on the value
	if b.quality == rand2.Int31n(10) {
		return false
	}
	return true
}

func (g *Game) Play(p1, p2 *Player) {
	go p1.hit(p2.court)
	for {
		select {
		// received
		case ball, ok := <-p2.court:
			fmt.Printf("\n%s received %v", p2.Name, ball)
			if !ok {
				fmt.Print("game over")
				return
			}
			isSuccess := p2.process(ball)
			if isSuccess {
				go p2.hit(p1.court)
			} else {
				p1.Score = p1.Score + 1
				return
			}

		case ball, ok := <-p1.court:
			fmt.Printf("\n%s received %v", p1.Name, ball)
			if !ok {
				fmt.Print("game over")
				return
			}
			isSuccess := p1.process(ball)
			if isSuccess {
				go p1.hit(p2.court)
			} else {
				p2.Score = p2.Score + 1
				return
			}
		}
	}
}
func (g *Game) GetResult() string {
	for _, p := range g.players {
		if p.Score == FINISHING_SCORE {
			return fmt.Sprintf("\n%s is the winner!", p.Name)
		}
	}
	return ""
}

func (g *Game) GetScore() string {
	return fmt.Sprintf("%s [%d] - %s [%d]\n", g.players[0].Name, g.players[0].Score,
		g.players[1].Name, g.players[1].Score)
}

func (g *Game) Register(p *Player) {
	g.players = append(g.players, p)
}
func (g *Game) isFinished() bool {
	if g.players[1].Score == FINISHING_SCORE || g.players[0].Score == FINISHING_SCORE {
		return true
	}
	return false
}

func switchSides(p1, p2 *Player) (*Player, *Player) {
	return p2, p1
}

func main() {
	p1 := NewPlayer("a")
	p2 := NewPlayer("b")

	p1.court = make(chan ball)
	p2.court = make(chan ball)
	defer close(p1.court)
	defer close(p2.court)

	g := NewGame(Singles)
	g.Register(p1)
	g.Register(p2)
	for !g.isFinished() {
		if (p1.Score+p2.Score)%5 == 0 {
			p1, p2 = switchSides(p1, p2)
		}
		g.Play(p1, p2)
		fmt.Printf("\nScores: %s", g.GetScore())
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("***********************")
	fmt.Println(g.GetResult())
	fmt.Println(g.GetScore())
}
