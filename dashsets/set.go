package dashsets

import "github.com/rbtyang/godash/dashdefine"

var empty dashdefine.Empty

// Set 集合类型，不要直接构造，通过NewSet
type Set[T comparable] struct {
	m map[T]dashdefine.Empty
}

/*
@Editor robotyang at 2023

NewSet 构造Set

集合set（存储的元素是唯一的）

为减少空间占用，其值可用空结构（struct{}，go会优化为不占用空间）。
*/
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]dashdefine.Empty),
	}
}

/*
@Editor robotyang at 2023

NewSet 构造有元素的Set

集合set（存储的元素是唯一的）

为减少空间占用，其值可用空结构（struct{}，go会优化为不占用空间）。
*/
func NewSetWith[T comparable](items ...T) *Set[T] {
	s := &Set[T]{
		m: make(map[T]dashdefine.Empty, len(items)),
	}
	for _, v := range items {
		s.m[v] = empty
	}
	return s
}

// Add 添加元素
func (s *Set[T]) Add(val T) {
	s.m[val] = empty
}

// Remove 删除元素
func (s *Set[T]) Remove(val T) {
	delete(s.m, val)
}

// Contains 检测是否包含元素
func (s *Set[T]) Contains(val T) bool {
	_, ok := s.m[val]
	return ok
}

// Equals 判断两集合是否相等
func (s *Set[T]) Equals(other *Set[T]) bool {
	if other == nil || s.Len() != other.Len() {
		return false
	}

	for k := range s.m {
		if _, ok := other.m[k]; !ok {
			return false
		}
	}
	return true
}

// Len 获取长度
func (s *Set[T]) Len() int {
	return len(s.m)
}

// Clear 清空set
func (s *Set[T]) Clear() {
	s.m = make(map[T]dashdefine.Empty)
}

// Items 遍历set
func (s *Set[T]) Items() []T {
	item := make([]T, 0)
	for k := range s.m {
		item = append(item, k)
	}
	return item
}
