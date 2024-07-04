package lastfm

type Options map[string][]string

func (o Options) Merge(other Options) Options {
	if other != nil {
		for k, v := range other {
			o[k] = v
		}
	}
	return o
}
