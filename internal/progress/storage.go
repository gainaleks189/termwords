package progress

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	dirName  = ".termwords"
	fileName = "progress.json"
)

func getProgressPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dirPath := filepath.Join(home, dirName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dirPath, fileName), nil
}

func Load() (*Progress, error) {
	path, err := getProgressPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// first run
		return &Progress{
			CurrentLanguage: "en",
			DailyNewWords:   5,
			Languages: map[string]LanguageProgress{
				"en": {CurrentIndex: 0},
			},
		}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p Progress
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func Save(p *Progress) error {
	path, err := getProgressPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
