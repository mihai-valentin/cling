package cling

type Flag struct {
	Name    string
	Enabled bool
}

func NewFlag(name string) *Flag {
	return &Flag{
		Name: name,
	}
}

func (f *Flag) Set(enabled bool) {
	f.Enabled = enabled
}
