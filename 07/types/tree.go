package types

type Entry struct {
	Name     string
	Size     int
	Parent   *Entry
	Children []*Entry
}

func (e *Entry) Child(s string) *Entry {
	for _, c := range e.Children {
		if c.Name == s {
			return c
		}
	}

	return nil
}

func (e *Entry) IsRoot() bool {
	return e.Parent == nil
}

func (e *Entry) IsFile() bool {
	return e.Children == nil
}

func (e *Entry) IsDir() bool {
	return e.Children != nil
}

func (e *Entry) UpdateSize(i int) {
	e.Size += i

	if e.Parent != nil {
		e.Parent.UpdateSize(i)
	}
}

func (e *Entry) Walk(fn func(x *Entry)) {
	for _, c := range e.Children {
		c.Walk(fn)
	}

	fn(e)
}
