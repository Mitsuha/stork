package audio

import "io"

// Metadata is an interface which is used to describe metadata retrieved by this package.
type Metadata interface {
	HasTag() bool

	// Duration returns the duration of the track.
	Duration() float64

	// Title returns the title of the track.
	Title() string

	// Album returns the album name of the track.
	Album() string

	// Artist returns the artist name of the track.
	Artist() string

	// AlbumArtist returns the album artist name of the track.
	AlbumArtist() string

	// Composer returns the composer of the track.
	Composer() string

	// Year returns the year of the track.
	Year() int

	// Genre returns the genre of the track.
	Genre() string

	// Track returns the track number and total tracks, or zero values if unavailable.
	Track() int

	// Disc returns the disc number and total discs, or zero values if unavailable.
	Disc() int

	//Picture returns a picture, or nil if not available.
	Picture() *Picture

	// Lyrics returns the lyrics, or an empty string if unavailable.
	Lyrics() string

	// Comment returns the comment, or an empty string if unavailable.
	Comment() string

	// Raw returns the raw mapping of retrieved tag names and associated values.
	// NB: tag/atom names are not standardised between formats.
	Raw() map[string]interface{}
}

type Picture struct {
	Data io.Reader
}
