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
			if len(k) == 0 {
				return nil, fmt.Errorf("grep: invalid key=value pair %q (empty key) in expression %q", kv, expr)
			}

			equal := true
			if k[len(k)-1] == '!' {
				k = k[:len(k)-1]
				k = strings.TrimSpace(k)
				equal = false
			}

			wildcardPrefix := false
			if len(v) > 0 && v[0] == '*' {
				v = v[1:]
				wildcardPrefix = true
			}
			wildcardSuffix := false
			if len(v) > 0 && v[len(v)-1] == '*' {
				v = v[:len(v)-1]
				wildcardSuffix = true
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
					K:              k,
					Vs:             []string{v},
					Equal:          equal,
					WildcardPrefix: wildcardPrefix,
					WildcardSuffix: wildcardSuffix,
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

func (g *Grep) Filter(m map[string]interface{}) bool {
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
	K              string
	Vs             []string
	Equal          bool
	WildcardPrefix bool
	WildcardSuffix bool
}

func (g *grepKV) Filter(m map[string]interface{}) bool {
	raw, ok := m[g.K]
	if !ok {
		return false
	}

	v, ok := raw.(string)
	if !ok {
		v = fmt.Sprintf("%v", raw)
	}

	for _, x := range g.Vs {
		switch {
		case v == x:
			return g.Equal
		case g.WildcardPrefix && g.WildcardSuffix:
			if strings.Contains(v, x) {
				return g.Equal
			}
			continue
		case g.WildcardPrefix:
			if strings.HasSuffix(v, x) {
				return g.Equal
			}
			continue
		case g.WildcardSuffix:
			if strings.HasPrefix(v, x) {
				return g.Equal
			}
			continue
		}
	}
	return !g.Equal
}
