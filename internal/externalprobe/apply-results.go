package externalprobe

type ApplyResults int

const (
	Unchanged ApplyResults = iota
	Configured
	Created
	Failed
)

func (a ApplyResults) String() string {
	return [...]string{"Unchanged", "Configured", "Created", "Failed"}[a]
}