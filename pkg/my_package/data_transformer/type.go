package data_transformer

import (
	"github.com/gol4ng/skeleton/pkg/my_package"
	protos "github.com/gol4ng/skeleton/pkg/my_package/protos"
)

func TransformType(t my_package.Type) protos.Document_Type {
	if v, ok := protos.Document_Type_value[string(t)]; ok {
		return protos.Document_Type(v)
	}
	panic("cannot transform type into protos type")
}

func ReverseTransformType(t protos.Document_Type) my_package.Type {
	if v, ok := my_package.Type_value[protos.Document_Type_name[int32(t)]]; ok {
		return my_package.Type(v)
	}
	panic("cannot transform protos type into type")
}
