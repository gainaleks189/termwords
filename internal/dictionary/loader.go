package dictionary

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed dict/*.json
var embeddedDicts embed.FS

func Load(language string) ([]Word, error) {
	filename := fmt.Sprintf("dict/%s.json", language)

	data, err := embeddedDicts.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var words []Word
	if err := json.Unmarshal(data, &words); err != nil {
		return nil, err
	}

	return words, nil
}
