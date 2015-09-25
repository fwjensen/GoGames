package main

import(
	"fmt"
	"math/rand"
	"time"
)

var(
	x, i = 0, 0
)

func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	randNum := rand.Intn(10)+1
	fmt.Println("Guess a number between 1 and 10?")
	for true{
		fmt.Scan(&x);
		i++
		if(x == randNum){
			fmt.Printf("Wow, you guessed it! Only took you %d trie(s)",i)
			break
		}else{
			fmt.Println("Nope, try again.")
		}
	}
}
