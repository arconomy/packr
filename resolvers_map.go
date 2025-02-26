//go:generate mapgen -name "resolvers" -zero "nil" -go-type "resolver.Resolver" -pkg "" -a "nil" -b "nil" -c "nil" -bb "nil" -destination "packr"
// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package packr

import (
	"sort"
	"sync"

	"github.com/arconomy/packr/file/resolver"
)

// resolversMap wraps sync.Map and uses the following types:
// key:   string
// value: resolver.Resolver
type resolversMap struct {
	data sync.Map
}

// Delete the key from the map
func (m *resolversMap) Delete(key string) {
	m.data.Delete(key)
}

// Load the key from the map.
// Returns resolver.Resolver or bool.
// A false return indicates either the key was not found
// or the value is not of type resolver.Resolver
func (m *resolversMap) Load(key string) (resolver.Resolver, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return nil, false
	}
	s, ok := i.(resolver.Resolver)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *resolversMap) LoadOrStore(key string, value resolver.Resolver) (resolver.Resolver, bool) {
	i, _ := m.data.LoadOrStore(key, value)
	s, ok := i.(resolver.Resolver)
	return s, ok
}

// Range over the resolver.Resolver values in the map
func (m *resolversMap) Range(f func(key string, value resolver.Resolver) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(resolver.Resolver)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a resolver.Resolver in the map
func (m *resolversMap) Store(key string, value resolver.Resolver) {
	m.data.Store(key, value)
}

// Keys returns a list of keys in the map
func (m *resolversMap) Keys() []string {
	var keys []string
	m.Range(func(key string, value resolver.Resolver) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
