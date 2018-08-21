package pkg

import (
	"fmt"
	"strings"
)

type Grep struct {
	kvs []*grepKV
}

func NewGrep(expr string) (*Grep, error) {
	var kvs []*grepKV

	if len(expr) != 0 {
		for _, kv := range strings.Split(expr, ",") {
			pair := strings.Split(kv, "=")
			if len(pair) != 2 {
				return nil, fmt.Errorf("grep: invalid key=value pair %q in expression %q", kv, expr)
			}
			k := strings.TrimSpace(pair[0])
			v := strings.TrimSpace(pair[1])

			equal := true
			if k[len(k)-1] == '!' {
				k = k[:len(k)-1]
				k = strings.TrimSpace(k)
				equal = false
			}

			var found *grepKV
			for _, kv := range kvs {
				if kv.K == k && kv.Equal == equal {
					found = kv
					break
				}
			}
			if found != nil {
				found.Vs = append(found.Vs, v)
			} else {
				kv := &grepKV{
					K:     k,
					Vs:    []string{v},
					Equal: equal,
				}
				kvs = append(kvs, kv)
			}
		}
	}

	g := &Grep{
		kvs: kvs,
	}

	return g, nil
}

func (g *Grep) Filter(m map[string]string) bool {
	if len(g.kvs) == 0 {
		return true
	}

	for _, kv := range g.kvs {
		if !kv.Filter(m) {
			return false
		}
	}

	return true
}

type grepKV struct {
	K     string
	Vs    []string
	Equal bool
}

func (g *grepKV) Filter(m map[string]string) bool {
	v, ok := m[g.K]
	if !ok {
		return false
	}

	for _, x := range g.Vs {
		if v == x {
			return g.Equal
		}
	}
	return !g.Equal
}
