package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gainaleks189/termwords/internal/dictionary"
)

func Run(words []dictionary.Word, start int, end int) {
	if len(words) == 0 {
		fmt.Println("No words available.")
		return
	}

	if end >= len(words) {
		end = len(words) - 1
	}

	reader := bufio.NewReader(os.Stdin)

	for i := start; i <= end; i++ {
		word := words[i]

		for {
			fmt.Printf("[%d] %s: ", word.ID, word.Prompt)

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == word.Answer {
				break
			}

			fmt.Printf("Wrong. Correct: %s\n", word.Answer)
		}
	}
}
