package audio

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"os/exec"
	"strconv"
	"time"
)

type FFProbeFormat struct {
	StartTime  string `json:"start_time,omitempty"`
	Duration   string `json:"duration,omitempty"`
	ProbeScore int    `json:"probe_score,omitempty"`
	Tags       struct {
		CreationTime time.Time `json:"creation_time,omitempty"`
		Artist       string    `json:"artist,omitempty"`
		Album        string    `json:"album,omitempty"`
		Comment      string    `json:"comment,omitempty"`
		Title        string    `json:"title,omitempty"`
		Disc         string    `json:"disc,omitempty"`
		Track        string    `json:"track,omitempty"`
		Genre        string    `json:"GENRE,omitempty"`
	} `json:"tags"`
}

type showFormatResult struct {
	Format FFProbeFormat `json:"format"`
}

func ffprobeReadMetadata(ffprobe string, reader io.Reader) (*FFProbeFormat, error) {
	cmd := exec.Command(ffprobe, "-i", "pipe:", "-show_format", "-v", "quiet", "-print_format", "json")

	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	errBuf := bytes.NewBuffer(make([]byte, 0))

	cmd.Stdin, cmd.Stdout, cmd.Stderr = reader, buffer, errBuf

	if err := cmd.Run(); err != nil {
		return nil, errors.New("ffprobe error: " + errBuf.String())
	}

	var result showFormatResult

	if err := json.Unmarshal(buffer.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	return &result.Format, nil
}

func ffmpegReadCover(ffmpeg string, reader io.Reader) (*Picture, error) {
	cmd := exec.Command(ffmpeg, "-i", "pipe:", "-f", "image2pipe", "-")

	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	errBuf := bytes.NewBuffer(make([]byte, 0))

	cmd.Stdin, cmd.Stdout, cmd.Stderr = reader, buffer, errBuf

	if err := cmd.Run(); err != nil {
		return nil, errors.New("ffmpeg error: " + errBuf.String())
	}

	if buffer.Len() == 0 {
		return nil, errors.New("no cover found")
	}

	return &Picture{
		Data: buffer,
	}, nil
}

type MetadataFromFfmpeg struct {
	Format *FFProbeFormat
	Cover  *Picture
}

func (m *MetadataFromFfmpeg) HasTag() bool {
	return m.Format != nil && m.Format.Tags.Title != "" && m.Format.Tags.Artist != ""
}

func (m *MetadataFromFfmpeg) Title() string {
	return m.Format.Tags.Title
}

func (m *MetadataFromFfmpeg) Album() string {
	return m.Format.Tags.Album
}

func (m *MetadataFromFfmpeg) Artist() string {
	return m.Format.Tags.Artist
}

func (m *MetadataFromFfmpeg) AlbumArtist() string {
	return m.Format.Tags.Artist
}

func (m *MetadataFromFfmpeg) Composer() string {
	return m.Format.Tags.Artist
}

func (m *MetadataFromFfmpeg) Year() int {
	return m.Format.Tags.CreationTime.Year()
}

func (m *MetadataFromFfmpeg) Genre() string {
	return ""
}

func (m *MetadataFromFfmpeg) Track() int {
	i, err := strconv.Atoi(m.Format.Tags.Track)
	if err != nil {
		return 0
	}
	return i
}

func (m *MetadataFromFfmpeg) Disc() int {
	i, err := strconv.Atoi(m.Format.Tags.Disc)
	if err != nil {
		return 0
	}
	return i
}

func (m *MetadataFromFfmpeg) Picture() *Picture {
	return m.Cover
}

func (m *MetadataFromFfmpeg) Duration() float64 {
	duration, err := strconv.ParseFloat(m.Format.Duration, 64)
	if err != nil {
		return 0
	}
	return duration
}

func (m *MetadataFromFfmpeg) Lyrics() string {
	//TODO implement me
	panic("implement me")
}

func (m *MetadataFromFfmpeg) Comment() string {
	return m.Format.Tags.Comment
}

func (m *MetadataFromFfmpeg) Raw() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}
