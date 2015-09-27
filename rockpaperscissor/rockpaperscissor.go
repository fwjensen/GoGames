package main

import(
	"fmt"
	"math/rand"
	"time"
)

func NewGame(){
	fmt.Println("How many sets do you wanna play(1-100)?")
	s := Turns()
	p := NewPoints()
	for i := 0;i < *s;i++{
		NewSet(p)
	}
	if p.cp > p.pp{
		fmt.Println("I win, better luck next time!")
	}else if p.cp < p.pp{
		fmt.Println("You win, GG!")
	}else{
		fmt.Println("It's a Tie.")
	}
}

func Turns() *int{
	var s int
	fmt.Scan(&s)
	if s > 0 && s <= 100{
		return &s
	}else{
		i := 1
		return &i
	}
}

func Hand() [3]string{
	var h [3]string
	h[0] = "rock"
	h[1] = "paper"
	h[2] = "scissor"
	return h
}

func NewCompHand() string{
	h := Hand()
	rand.Seed(time.Now().UTC().UnixNano())
	return h[rand.Intn(3)]
}

func NewPlayerHand() string{
	fmt.Print("Choose(rock, paper or scissor): ")
	h := Hand()
	var ph string
	if fmt.Scan(&ph);ph == h[0]{
		return h[0]
	}else if ph == h[1]{
		return h[1]
	}else if ph == h[2]{
		return h[2]
	}else{
		fmt.Println("Not a valid input.")
		return NewPlayerHand()
	}
}

func NewSet(p *Points){
	ch := NewCompHand()
	ph := NewPlayerHand()
	CompareHands(&ch, &ph, p)

}

type Points struct{
	cp, pp int
}

func NewPoints() *Points{
	return &Points{cp: 0, pp: 0}
}

func (p *Points)AddPoint(cp, pp int){
	p.cp += cp
	p.pp += pp
	fmt.Printf("Me: %d:%d :You\n",p.cp,p.pp)
}

func CompareHands(ch, ph *string, p *Points){
	if *ch == *ph{
		fmt.Printf("%s vs %s - It's a tie, lets go again.\n",*ch,*ph)
		NewSet(p)
	}else if *ch == "rock"{
		if *ph == "paper"{
			fmt.Println("rock vs paper - You won that round")
			p.AddPoint(0,1)
		}else{
			fmt.Println("rock vs scissor - I won that round")
			p.AddPoint(1,0)
		}
	}else if *ch == "scissor"{
		if *ph == "rock"{
			fmt.Println("scissor vs rock - You won that round")
			p.AddPoint(0,1)
		}else{
			fmt.Println("scissor vs paper - I won that round")
			p.AddPoint(1,0)
		}
	}else{
		if *ph == "scissor"{
			fmt.Println("paper vs scissor - You won that round")
			p.AddPoint(0,1)
		}else{
			fmt.Println("paper vs rock - I won that round")
			p.AddPoint(1,0)
		}
	}
}

func main(){
	NewGame()
}
