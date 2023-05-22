package stack

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func (s *Stack) PushAtTheBack(v interface{}) {
	s.list.PushBack(v)
}

func (s *Stack) PopFromTheBack() interface{} {
	item := s.list.Back()
	s.list.Remove(item)
	return item
}
