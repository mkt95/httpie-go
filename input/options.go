package input

type Options struct {
	JSON       bool
	Form       bool
	ReadStdin  bool
	TrackingID TrackingIDOptions
}

type TrackingIDOptions struct {
	Enabled bool
	Sender  string
}
