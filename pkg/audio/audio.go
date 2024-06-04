package audio

import (
	"io"
	p "path"
)

const OSPath = ""

type Audio struct {
	ffmpegPath  string
	ffProbePath string
	media       io.ReadSeeker
}

func New(path string) *Audio {
	return &Audio{
		ffmpegPath:  p.Join(path, "ffmpeg"),
		ffProbePath: p.Join(path, "ffprobe"),
	}
}

func (a *Audio) LoadFromReader(reader io.ReadSeeker) *Audio {
	a.media = reader
	return a
}

func (a *Audio) Metadata() (Metadata, error) {
	metadata, err := ffprobeReadMetadata(a.ffProbePath, a.media)
	if err != nil {
		return nil, err
	}

	if _, err = a.media.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	cover, _ := ffmpegReadCover(a.ffmpegPath, a.media)

	return &MetadataFromFfmpeg{
		Format: metadata,
		Cover:  cover,
	}, nil
}
