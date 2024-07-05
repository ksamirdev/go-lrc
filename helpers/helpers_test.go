package helpers

import (
	"testing"

	"github.com/samocodes/go-lrc/types"
)

func TestGenerateLRC(t *testing.T) {
	music := types.Music{
		Title:    "Please Please Please",
		Artist:   "Sabrina Carpenter",
		Album:    "Please Please Please",
		Author:   "Amy Allen",
		By:       "Samir",
		Language: "eng",
		Lyrics: []types.Lyrics{
			{
				Time:  "00:18.00",
				Value: "I know I have good judgment",
			}, {
				Time:  "00:20.00",
				Value: "I know I have good taste",
			},
		},
	}

	t.Log(GenerateLRC(music))
}
