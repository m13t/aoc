package types

// type Stack []byte
type item struct {
	p *item
	v byte
}

type Stack struct {
	head *item
	len  int
}

func newStack() *Stack {
	return &Stack{
		head: nil,
		len:  0,
	}
}

func (s *Stack) Push(v byte) {
	x := &item{
		v: v,
		p: s.head,
	}

	s.head = x
	s.len += 1
}

func (s *Stack) Peek() (byte, bool) {
	if s.len == 0 {
		return 0, false
	}

	return s.head.v, true
}

func (s *Stack) Pop() (byte, bool) {
	if s.len == 0 {
		return 0, false
	}

	x := s.head
	s.head = x.p
	s.len -= 1

	return x.v, true
}

type Stacks []*Stack

func NewStacks() Stacks {
	return Stacks{}
}

func (s *Stacks) PushTo(col int, val byte) {
	for len(*s) < col+1 {
		*s = append(*s, newStack())
	}

	(*s)[col].Push(val)
}

func (s *Stacks) PopFrom(col int) (byte, bool) {
	for len(*s) < col+1 {
		*s = append(*s, newStack())
	}

	return (*s)[col].Pop()
}

func (s *Stacks) MoveOne(count, from, to int) {
	for x := 0; x < count; x++ {
		if v, ok := s.PopFrom(from - 1); ok {
			s.PushTo(to-1, v)
		}
	}
}

func (s *Stacks) MoveSet(count, from, to int) {
	tmp := make([]byte, count)

	for x := 0; x < count; x++ {
		if v, ok := s.PopFrom(from - 1); ok {
			tmp[x] = v
		}
	}

	for x := len(tmp) - 1; x >= 0; x-- {
		s.PushTo(to-1, tmp[x])
	}
}
