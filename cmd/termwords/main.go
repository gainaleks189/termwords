package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gainaleks189/termwords/internal/dictionary"
	"github.com/gainaleks189/termwords/internal/engine"
	"github.com/gainaleks189/termwords/internal/progress"
	"github.com/gainaleks189/termwords/internal/tui"
)

func main() {
	p, err := progress.Load()
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args

	if len(args) >= 2 && args[1] == "help" {
		fmt.Println("Termwords CLI")
		fmt.Println("")
		fmt.Println("Available commands:")
		fmt.Println("  status        Show current status")
		fmt.Println("  set <number>  Set daily new words")
		fmt.Println("  reset         Reset current progress")
		fmt.Println("  use <lang>    Switch language")
		fmt.Println("")
		fmt.Println("Run without arguments to start session.")
		return
	}

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

	if len(args) >= 3 && args[1] == "use" {
		lang := args[2]
		_, err := dictionary.Load(lang)
		if err != nil {
			log.Fatalf("Language '%s' not found.", lang)
		}
		p.CurrentLanguage = lang
		if _, exists := p.Languages[lang]; !exists {
			p.Languages[lang] = progress.LanguageProgress{
				CurrentIndex: 0,
			}
		}
		if err := progress.Save(p); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Language switched to:", lang)
		return
	}

	if len(args) >= 2 && args[1] == "status" {
		words, err := dictionary.Load(p.CurrentLanguage)
		if err != nil {
			log.Fatal(err)
		}
		current := p.Languages[p.CurrentLanguage].CurrentIndex
		start, end := engine.CalculateWindow(current, p.DailyNewWords, len(words))
		fmt.Printf("termwords · %s · %d/day · %d/%d\n", p.CurrentLanguage, p.DailyNewWords, current, len(words))
		fmt.Printf("Window: %d–%d\n", start+1, end+1)
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

	words, err := dictionary.Load(p.CurrentLanguage)
	if err != nil {
		log.Fatal(err)
	}

	lang := p.CurrentLanguage
	current := p.Languages[lang].CurrentIndex
	start, end := engine.CalculateWindow(current, p.DailyNewWords, len(words))

	if len(words) == 0 {
		fmt.Println("No words available.")
		return
	}
	if end >= len(words) {
		end = len(words) - 1
	}

	model := tui.New(words, start, end, lang, p.DailyNewWords)
	program := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	final, err := program.Run()
	fmt.Print("\033[6 q")
	if err != nil {
		log.Fatal(err)
	}

	if fm, ok := final.(tui.Model); ok && fm.Completed {
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
}
