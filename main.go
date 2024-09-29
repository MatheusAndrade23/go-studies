package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número será sorteado e você terá que adivinhar qual é. O número está entre 0 e 100.")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O seu chute precisa ser um inteiro.")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Seu chute foi menor que o número sorteado.")
		case chuteInt > x:
			fmt.Println("Seu chute foi maior que o número sorteado.")
		case chuteInt == x:
			fmt.Printf(
				"Parabéns! O número era %d\n"+ 
				"Você acertou em %d tentativas.\n"+ 
				"Esses foram suas chutes: %v\n", 
				x, i + 1, chutes[:i])
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Você perdeu! O número sorteado foi %d."+ 
		"Você teve 10 tentativas!\n"+ 
		"Essas foram suas tentativas: %v\n", 
		x, chutes)
}