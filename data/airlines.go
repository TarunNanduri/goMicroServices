package data

// AirLines is an structure to hold aeroplane related data
type AirLines struct {
	Name        string
	Source      string
	Destination string
}

// Options struct to fill in airlines
type Options func(*AirLines)

// WithName is used to initialize name of airline
func WithName(name string) Options {
	return func(o *AirLines) {
		o.Name = name
	}
}

// WithSource is used to initialize name of airline
func WithSource(source string) Options {
	return func(o *AirLines) {
		o.Source = source
	}
}

// WithDestination is used to initialize name of airline
func WithDestination(destination string) Options {
	return func(o *AirLines) {
		o.Destination = destination
	}
}

// NewAirlines creator
func NewAirlines(options ...Options) *AirLines {
	opts := &AirLines{}
	for _, opt := range options {
		opt(opts)
	}
	return opts
}
