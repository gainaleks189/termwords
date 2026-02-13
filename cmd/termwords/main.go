package main

import (
	"fmt"
	"log"

	"github.com/gainaleks189/termwords/internal/dictionary"
	"github.com/gainaleks189/termwords/internal/engine"
	"github.com/gainaleks189/termwords/internal/progress"
	"github.com/gainaleks189/termwords/internal/session"
)

func main() {
	p, err := progress.Load()
	if err != nil {
		log.Fatal(err)
	}

	words, err := dictionary.Load(p.CurrentLanguage)
	if err != nil {
		log.Fatal(err)
	}

	lang := p.CurrentLanguage
	current := p.Languages[lang].CurrentIndex

	start, end := engine.CalculateWindow(current, p.DailyNewWords)

	fmt.Printf("Active window: %d - %d\n", start, end)

	session.Run(words, start, end)

	// 🔥 ВАЖНО: двигаем индекс
	current += p.DailyNewWords

	if current >= len(words) {
		current = len(words) - 1
	}

	p.Languages[lang] = progress.LanguageProgress{
		CurrentIndex: current,
	}
	fmt.Println("Saving index:", current)
	if err := progress.Save(p); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Session complete.")
}