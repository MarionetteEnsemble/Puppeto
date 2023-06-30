package engine

import (
	"fmt"
)

type Scope struct {
	List        map[string]PValue
	ParentScope *Scope
}

func NewScope(parentScope *Scope) *Scope {
	return &Scope{
		List:        make(map[string]PValue),
		ParentScope: parentScope,
	}
}

func (s *Scope) Define(key string, value PValue) error {
	if s.Has(key) {
		return fmt.Errorf("Variable %s already defined", key)
	}

	s.List[key] = value
	return nil
}

func (s *Scope) Get(key string) (PValue, error) {
	value, ok := s.List[key]
	if !ok {
		if s.ParentScope != nil {
			return s.ParentScope.Get(key)
		}
		return nil, fmt.Errorf("Variable %s is not defined", key)
	}
	return value, nil
}

func (s *Scope) Has(key string) bool {
	_, ok := s.List[key]
	if !ok && s.ParentScope != nil {
		return s.ParentScope.Has(key)
	}
	return ok
}

func (s *Scope) Set(key string, value PValue) error {
	if _, ok := s.List[key]; !ok {
		if s.ParentScope != nil {
			return s.ParentScope.Set(key, value)
		}
		return fmt.Errorf("Variable %s is not defined", key)
	}
	s.List[key] = value
	return nil
}
