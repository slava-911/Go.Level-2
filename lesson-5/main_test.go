package main

import (
	"sync"
	"testing"
)

type SetMutex struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSetMutex() *SetMutex {
	return &SetMutex{
		mm: map[int]struct{}{},
	}
}

func (s *SetMutex) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *SetMutex) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type SetRWMutex struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewSetRWMutex() *SetRWMutex {
	return &SetRWMutex{
		mm: map[int]struct{}{},
	}
}

func (s *SetRWMutex) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *SetRWMutex) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkSetMutexAdd10(b *testing.B) {
	var set = NewSetMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}
func BenchmarkSetMutexHas90(b *testing.B) {
	var set = NewSetMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetMutexAdd50(b *testing.B) {
	var set = NewSetMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}
func BenchmarkSetMutexHas50(b *testing.B) {
	var set = NewSetMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetMutexAdd90(b *testing.B) {
	var set = NewSetMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}
func BenchmarkSetMutexHas10(b *testing.B) {
	var set = NewSetMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetRWMutexAdd10(b *testing.B) {
	var set = NewSetRWMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}
func BenchmarkSetRWMutexHas90(b *testing.B) {
	var set = NewSetRWMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetRWMutexAdd50(b *testing.B) {
	var set = NewSetRWMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}
func BenchmarkSetRWMutexHas50(b *testing.B) {
	var set = NewSetRWMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetRWMutexAdd90(b *testing.B) {
	var set = NewSetRWMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}
func BenchmarkSetRWMutexHas10(b *testing.B) {
	var set = NewSetRWMutex()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}
