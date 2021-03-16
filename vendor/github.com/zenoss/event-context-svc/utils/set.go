package utils

type StringSet struct {
	set map[string]bool
}

func NewStringSet() *StringSet {
	set := make(map[string]bool)
	return &StringSet{set: set}
}

func (s *StringSet) Values() []string {
	results := make([]string, 0, len(s.set))
	for k := range s.set {
		results = append(results, k)
	}
	return results
}

func (s *StringSet) Add(v string) bool {
	if !s.set[v] {
		s.set[v] = true
		return true
	}
	return false
}

func (s *StringSet) Remove(v string) bool {
	if !s.set[v] {
		return false
	}
	delete(s.set, v)
	return true
}

func (s *StringSet) Size() int {
	return len(s.set)
}
