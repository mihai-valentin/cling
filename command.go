package cling

type Command struct {
	Name        string
	Description string
}

func NewCommand(name string, options ...func(*Command)) *Command {
	c := &Command{
		Name: name,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *Command) GetName() string {
	return c.Name
}

func WithDescription(d string) func(c *Command) {
	return func(c *Command) {
		c.Description = d
	}
}
