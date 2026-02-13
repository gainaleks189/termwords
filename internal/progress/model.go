package progress

type LanguageProgress struct {
	CurrentIndex int `json:"current_index"`
}

type Progress struct {
	CurrentLanguage string                       `json:"current_language"`
	DailyNewWords   int                          `json:"daily_new_words"`
	Languages       map[string]LanguageProgress  `json:"languages"`
}
