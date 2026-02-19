

Project: termwords

termwords is a terminal-based vocabulary trainer written in Go. It presents a word or phrase and prompts the user to enter the translation, immediately showing whether the answer is correct. Progress is saved locally, and each session combines review and new words. The review uses a sliding window model where words remain active for 10 “days,” ensuring repetition without overwhelming the user. The number of new words per day is configurable (e.g., 5, 10, or 20).

The tool supports multiple languages using embedded JSON dictionaries, including practical languages (English, Finnish, Spanish, French) as well as experimental ones such as Klingon, demonstrating flexible data handling and extensibility. Dictionaries are embedded into the binary, and user progress is stored in the home directory.

My role

I designed and implemented the entire application independently. This includes the sliding window repetition logic, terminal UI using Bubble Tea, visual styling with Lipgloss, dictionary loading via Go embed, and persistent progress storage.

Challenges

The main challenge was designing a clear and efficient terminal UI. I implemented contextual hints only for new words, separated review and new items in the model, and ensured a readable interface without visual overload.

Improvements

Future improvements include adding unit tests, supporting additional input methods for complex alphabets, implementing true spaced repetition, and documenting the dictionary format.

Tech: Go, Bubble Tea, Lipgloss, JSON, embed

