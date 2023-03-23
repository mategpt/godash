package godash

import "github.com/imdario/mergo"

func Merge(dst, src interface{}) interface{} {
	return mergo.Merge(dst, src, mergo.WithOverride)
}
