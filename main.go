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
	fmt.Println("Bem-vindo(a) ao jogo da adivinhação!")
	fmt.Println("| Irei sortear um número de 0 a 100. Sua missão é adivinhar o número sorteado, acha que consegue?")

	random := rand.Int64N(101) // Gera um número aleatório entre 0 e 100
	chutes := [10]int64{}

	scanner := bufio.NewScanner(os.Stdin)

	for i := range chutes {
		fmt.Println("Qual é a sua tentativa?")

		scanner.Scan() // Blocking operation

		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("Ops! Parece que você digitou um valor não permitido.\n" +
				"Isso é contra as regras!")
			return
		}

		if i == 0 && chuteInt == random { // Caso o jogador adivinhe na primeira tentativa.
			fmt.Printf(
				"\n😱 NFUHF98#$#*0 😱 Afaste-se de mim, bruxo(a)! \n"+
					"O número sorteado era: %d\n", random)
			return
		}

		switch {
		case chuteInt > random:
			fmt.Println("Hum... o número sorteado é menor que: ", chuteInt)
		case chuteInt < random:
			fmt.Println("Hum... o número sorteado é maior que: ", chuteInt)
		case chuteInt == random:
			fmt.Printf("\nParabéns! 😄\n"+
				"Você acertou o número que eu estava adivinhando.\n"+
				"O número era %d.\n"+
				"Você demorou: %d tentativa(s).\n"+
				"O(s) seu(s) palpite(s) foram: %v \n", random, i+1, chutes[:i])
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"\nOpa! Suas chances acabaram e a terra explodiu 💥. Você é o grande culpado(a).\n"+
			"O número que eu estava pensando era %d.\n"+
			"Os números que você tentou foram: %v\n", random, chutes)
}
