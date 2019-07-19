package my_package

type Type string

const (
	IMAGE Type = "IMAGE"
	TEXT  Type = "TEXT"
	VIDEO Type = "VIDEO"
)

var Type_name = map[Type]string{
	IMAGE: "IMAGE",
	TEXT: "TEXT",
	VIDEO: "VIDEO",
}
var Type_value = map[string]Type{
	"IMAGE": IMAGE,
	"TEXT":  TEXT,
	"VIDEO": VIDEO,
}
