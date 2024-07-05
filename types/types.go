package types

type Music struct {
	Title    string   `json:"title"`
	Artist   string   `json:"artist"`
	Album    string   `json:"album"`
	Author   string   `json:"author"`
	By       string   `json:"by"`
	Length   string   `json:"length"` // mm:ss.xx
	Language string   `json:"language"`
	Lyrics   []Lyrics `json:"lyrics"`
}

type Lyrics struct {
	Time  string `json:"time"`
	Value string `json:"value"`
}
