package audio

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("when has no path", func(t *testing.T) {
		audio := New("")

		if audio.ffmpegPath != "ffmpeg" {
			t.Error("ffmpegPath should be ffmpeg")
		}
	})
	t.Run("when has path", func(t *testing.T) {
		audio := New("/usr/bin")

		if audio.ffmpegPath != "/usr/bin/ffmpeg" {
			t.Error("ffmpegPath should be /usr/bin/ffmpeg")
		}
	})
}

func TestAudio_Metadata(t *testing.T) {
	file, _ := os.Open("/mnt/c/Users/Mitsuha/Desktop/music.mp3")

	metadata, err := New("").LoadFromReader(file).Metadata()

	t.Run("Metadata", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}

		fmt.Println(metadata)
	})
}
