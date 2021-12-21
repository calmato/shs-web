package set

import "sort"

type Set struct {
	values map[interface{}]struct{}
}

func New(cap int) *Set {
	return &Set{
		values: make(map[interface{}]struct{}, cap),
	}
}

func (s *Set) Reset(cap int) {
	s.values = make(map[interface{}]struct{}, cap)
}

func (s *Set) Len() int {
	return len(s.values)
}

func (s *Set) Contains(values ...interface{}) bool {
	for i := range values {
		if _, ok := s.values[values[i]]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) FindOrAdd(values ...interface{}) bool {
	isFind := false
	for i := range values {
		if _, ok := s.values[values[i]]; ok {
			continue
		}
		s.Add(values[i])
		isFind = true
	}
	return isFind
}

func (s *Set) Add(values ...interface{}) {
	for _, v := range values {
		s.values[v] = struct{}{}
	}
}

func (s *Set) AddStrings(values ...string) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *Set) AddInt64s(values ...int64) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *Set) Remove(v interface{}) {
	delete(s.values, v)
}

func (s *Set) Do(f func(interface{})) {
	for v := range s.values {
		f(v)
	}
}

func (s *Set) Strings() []string {
	res := make([]string, 0, s.Len())
	for v := range s.values {
		res = append(res, v.(string))
	}
	return res
}

func (s *Set) SortStrings() []string {
	res := s.Strings()
	sort.Strings(res)
	return res
}

func (s *Set) Int64s() []int64 {
	res := make([]int64, 0, s.Len())
	for v := range s.values {
		res = append(res, v.(int64))
	}
	return res
}

func (s *Set) SortInt64s() []int64 {
	res := s.Int64s()
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}
