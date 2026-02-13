package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gainaleks189/termwords/internal/dictionary"
)

func Run(words []dictionary.Word, start, end int, inputStartRow int, inputCol int) {
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
	row := inputStartRow + (i - start)  // строка таблицы для этого слова

	// Курсор в правую колонку таблицы, напротив слова
	fmt.Printf("\033[%d;%dH", row, inputCol)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input != word.Answer {
		// сообщение об ошибке (например, внизу экрана или рядом)
		fmt.Printf("\033[%d;%dHWrong. Correct: %s", row, inputCol, word.Answer)
		// и повторить ввод для этого слова
	}
}
}
