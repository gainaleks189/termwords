package dictionary

type Word struct {
	ID     int    `json:"id"`
	Prompt string `json:"prompt"`
	Answer string `json:"answer"`
}
