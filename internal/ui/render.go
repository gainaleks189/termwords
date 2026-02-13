package ui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gainaleks189/termwords/internal/dictionary"
)
const (
	// Отступ от вертикальной черты │ до позиции ввода (в пробелах)
	inputMarginFromSeparator = 2
)
func RenderWindow(
	words []dictionary.Word,
	start, end int,
	language string,
	daily int,
	current int,
) (inputStartRow int, inputCol int) {

	// Очистка экрана
	fmt.Print("\033[H\033[2J")

	total := len(words)

	// Заголовок
	fmt.Printf("termwords · %s · %d/day · %d/%d\n",
		language,
		daily,
		current,
		total,
	)

	fmt.Println(strings.Repeat("─", 60))

	// Таблица начинается с 3 строки
	inputStartRow = 3

	// Вычисляем максимальную длину слова
	// Вычисляем максимальную длину слова (в символах, не в байтах)
maxLen := 0
for i := start; i <= end && i < total; i++ {
	l := utf8.RuneCountInString(words[i].Prompt)
	if l > maxLen {
		maxLen = l
	}
}

	// Печать слов
	// Печать слов
row := inputStartRow
for i := start; i <= end && i < total; i++ {
	rcount := utf8.RuneCountInString(words[i].Prompt)
	padding := maxLen - rcount
	fmt.Printf("%03d  %s%s │\n",
		i+1,
		words[i].Prompt,
		strings.Repeat(" ", padding),
	)
	row++
}

	fmt.Println(strings.Repeat("─", 60))
	fmt.Println("10 minutes · focus")

	// Колонка ввода: после │ с отступом
// 3 (номер) + 2 пробела + maxLen + 1 (│) + отступ
inputCol = 3 + 2 + maxLen + 1 + inputMarginFromSeparator

	return inputStartRow, inputCol
}