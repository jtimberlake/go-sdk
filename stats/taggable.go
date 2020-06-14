package stats

// Taggable is an interface for specifying and retrieving default stats tags
type Taggable interface {
	AddDefaultTags(...string)
	DefaultTags() []string
}
