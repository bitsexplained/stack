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

func (s *Stack) PushAtTheFront(v interface{}) {
	s.list.PushFront(v)
}

func (s *Stack) PopFromTheFront() interface{} {
	item := s.list.Front()
	s.list.Remove(item)
	return item
}

func (s *Stack) InsertValueBefore(value interface{}, key interface{}) {
	s.list.InsertBefore(value, key.(*list.Element))
}

func (s *Stack) InsertValueAfter(value interface{}, key interface{}) {
	s.list.InsertAfter(value, key.(*list.Element))
}
