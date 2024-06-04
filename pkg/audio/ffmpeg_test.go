package audio

import (
	"io"
	"os"
	"testing"
)

var file *os.File

func TestMain(m *testing.M) {
	file, _ = os.Open("/mnt/c/Users/Mitsuha/Desktop/music.mp3")

	m.Run()

	_ = file.Close()
}

func TestFfprobeReadMetadata(t *testing.T) {
	_, _ = file.Seek(0, io.SeekStart)

	metadata, err := ffprobeReadMetadata("ffprobe", file)

	t.Run("FfprobeReadMetadata", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}

		t.Log(metadata)
	})
}

func TestFfmpegReadCover(t *testing.T) {
	_, _ = file.Seek(0, io.SeekStart)

	_, err := ffmpegReadCover("ffmpeg", file)

	t.Run("FfmpegReadCover", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})
}
