# termwords

Terminal vocabulary trainer written in Go with sliding-window repetition and interactive TUI.

## Overview

termwords is a CLI application that helps users learn and practice vocabulary directly in the terminal. The program presents a word or phrase and prompts the user to enter the translation, immediately validating the answer.

Progress is stored locally, and each session combines new and review words. The review system uses a sliding window model, where words remain active for 10 days to ensure consistent repetition without overwhelming the user. The number of new words per session is configurable.

The application supports multiple languages via embedded JSON dictionaries. Dictionaries are compiled into the binary using Go embed, and user progress is stored in the home directory.

This project demonstrates clean CLI architecture, persistent storage, and terminal UI design using modern Go libraries.

## Features

- interactive terminal UI (TUI)
- sliding-window repetition model
- configurable number of new words per session
- persistent progress storage
- embedded dictionaries using Go embed
- multi-language support
- fast startup and minimal dependencies

## Tech Stack

- Go
- Bubble Tea (terminal UI framework)
- Lipgloss (terminal styling)
- JSON (dictionary and progress storage)
- Go embed (static assets)

## Installation

Clone repository:

```bash
git clone https://github.com/gainaleks189/termwords
cd termwords
```

Run:

```bash
go run ./cmd/termwords
```

Or build binary:

```bash
go build -o termwords ./cmd/termwords
./termwords
```

## Architecture

Project follows standard Go layout:

```
cmd/termwords    # application entry point
internal/        # core application logic
```

This structure separates CLI entry point from business logic and improves maintainability.

## Author

Aleksandr Gainullin

Designed and implemented the complete application, including architecture, repetition logic, terminal UI, dictionary handling, and persistent storage.

## Planned Improvements

- unit tests for repetition logic
- spaced repetition algorithm
- support for additional input methods
- extended dictionary format documentation
