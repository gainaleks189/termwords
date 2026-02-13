package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gainaleks189/termwords/internal/dictionary"
	"github.com/gainaleks189/termwords/internal/engine"
	"github.com/gainaleks189/termwords/internal/progress"
	"github.com/gainaleks189/termwords/internal/session"
)

func main() {
	// 1. Load progress
	p, err := progress.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Handle CLI arguments
	args := os.Args
	if len(args) >= 2 && args[1] == "reset" {
		p.Languages[p.CurrentLanguage] = progress.LanguageProgress{
			CurrentIndex: 0,
		}
	
		if err := progress.Save(p); err != nil {
			log.Fatal(err)
		}
	
		fmt.Println("Progress reset.")
		return
	}
	// STATUS
if len(args) >= 2 && args[1] == "status" {

	words, err := dictionary.Load(p.CurrentLanguage)
	if err != nil {
		log.Fatal(err)
	}

	current := p.Languages[p.CurrentLanguage].CurrentIndex
	start, end := engine.CalculateWindow(current, p.DailyNewWords, len(words))

	fmt.Println("Status:")
	fmt.Println("Language:", p.CurrentLanguage)
	fmt.Println("Daily new words:", p.DailyNewWords)
	fmt.Println("Current index:", current)
	fmt.Printf("Active window: %d - %d\n", start, end)

	return
}
	if len(args) >= 3 && args[1] == "set" {
		value, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal("Invalid number")
		}

		if value <= 0 {
			log.Fatal("Value must be positive")
		}

		p.DailyNewWords = value

		if err := progress.Save(p); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Daily new words set to:", value)
		return
	}

	// 3. Load dictionary
	words, err := dictionary.Load(p.CurrentLanguage)
	if err != nil {
		log.Fatal(err)
	}

	lang := p.CurrentLanguage
	current := p.Languages[lang].CurrentIndex

	// 4. Calculate window
	start, end := engine.CalculateWindow(current, p.DailyNewWords, len(words))

	fmt.Printf("Active window: %d - %d\n", start, end)

	// 5. Run session
	session.Run(words, start, end)

	// 6. Move index forward
	current += p.DailyNewWords

	if current >= len(words) {
		current = len(words) - 1
	}

	p.Languages[lang] = progress.LanguageProgress{
		CurrentIndex: current,
	}

	if err := progress.Save(p); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Session complete.")
}