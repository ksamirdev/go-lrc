package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/samocodes/go-lrc/types"
)

func SupportsHTML(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "text/html")
}

var (
	title    = "ti"
	artist   = "ar"
	album    = "al"
	author   = "au"
	by       = "by"
	length   = "length"
	language = "la"
)

func GenerateLRC(music types.Music) string {
	var data strings.Builder

	data.WriteString(fmt.Sprintf("[%s:%s]\n", title, music.Title))
	data.WriteString(fmt.Sprintf("[%s:%s]\n", artist, music.Artist))
	data.WriteString(fmt.Sprintf("[%s:%s]\n", album, music.Album))
	data.WriteString(fmt.Sprintf("[%s:%s]\n", author, music.Author))
	data.WriteString(fmt.Sprintf("[%s:%s]\n", by, music.By))

	if len(music.Lyrics) > 0 {
		data.WriteString(fmt.Sprintf("[%s:%s]\n", length, music.Lyrics[len(music.Lyrics)-1].Time))
	} else {
		data.WriteString(fmt.Sprintf("[%s:%s]\n", length, music.Length))
	}

	data.WriteString(fmt.Sprintf("[%s:%s]\n\n", language, music.Language))

	for _, lyric := range music.Lyrics {
		data.WriteString(fmt.Sprintf("[%s]%s\n", lyric.Time, lyric.Value))
	}

	return data.String()
}
