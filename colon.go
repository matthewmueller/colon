package colon

import (
	"fmt"
	"regexp"
	"strings"
)

// Render function type
type Render func(map[string]interface{}) string

// Compile the template
func Compile(template string) Render {
	return func(obj map[string]interface{}) string {
		obj = flatten(obj)
		v := make([]string, 0, len(obj))

		for _, value := range obj {
			v = append(v, fmt.Sprintf("%v", value))
		}

		rtemplate := regexp.MustCompile("(\\\\?:|::)(" + strings.Join(v, "|") + "|[$A-Za-z_][[$A-Za-z_\\.0-9]+)")
		out := rtemplate.ReplaceAllFunc([]byte(template), func(m []byte) []byte {
			value := obj[string(m[1:])]
			if value == nil {
				return m
			}

			return []byte(fmt.Sprintf("%v", value))
		})

		return string(out)
	}
}

// Flatten the dependencies
func flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}
