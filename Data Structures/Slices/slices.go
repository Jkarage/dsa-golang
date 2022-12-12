package slices

import "errors"

type Slice struct {
	Len  int
	Cap  int
	Data []interface{}
}

var defaultCapacity = 16

func (s *Slice) Length() int {
	return s.Len
}

func (s *Slice) Capacity() int {
	return s.Cap
}

func (s *Slice) ValueAt(i int) (interface{}, error) {
	if i > s.Len {
		return nil, errors.New("index out of range")
	}
	return s.Data[i], nil
}

func (s *Slice) Insert(v interface{}) {
	if (s.Len + 1) >= s.Cap {
		if s.Len == 0 {
			s.Cap = defaultCapacity
		} else {
			// bit shift the capacity
			s.Cap = s.Cap << 1
		}
		newArray := make([]interface{}, s.Cap)
		if s.Len > 0 {
			copy(newArray, s.Data)
		}
		s.Data = newArray
	}
	s.Data[s.Len] = v
	s.Len++
}

func (s *Slice) RemoveAt(i int) (interface{}, error) {
	if i > s.Len {
		return nil, errors.New("index out of bounds")
	}

	removed := s.Data[i]
	copy(s.Data[i:], s.Data[i+1:])
	s.Len--
	return removed, nil
}

func (s *Slice) InsertAt(i int, v interface{}) (interface{}, error) {
	if i > s.Len {
		return nil, errors.New("index out of Bound")
	}
	if (s.Len + 1) >= s.Cap {
		if s.Len == 0 {
			s.Cap = defaultCapacity
		} else {
			// bit shift capacity
			s.Cap = s.Cap << 1
		}
		newArray := make([]interface{}, s.Cap)
		copy(newArray, s.Data)
		s.Data = newArray
	}
	copy(s.Data[i+1:], s.Data[i:])
	s.Data[i] = v
	s.Len++
	return v, nil
}

func (s *Slice) IndexOf(p interface{}) int {
	for i, v := range s.Data {
		if v == p {
			return i
		}
	}
	return -1
}

func (s *Slice) Reverse() []interface{} {
	for i, j := 0, s.Length()-1; i < j; i, j = i+1, j-1 {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	return s.Data
}

func (s *Slice) Contains(a interface{}) bool {
	return s.IndexOf(a) != -1
}

func Intersect(a, b []interface{}) []interface{} {
	var m []interface{}
	for _, v := range a {
		for _, h := range b {
			if v == h {
				m = append(m, v)
			}
		}
	}
	return m
}
