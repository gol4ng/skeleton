package config

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"text/tabwriter"
)

func ToString(iface interface{}) string {
	b := &bytes.Buffer{}
	w := tabwriter.NewWriter(b, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprint(w, "\n-----------------------------------\n---  Application configuration  ---\n")
	fprint(w, iface)
	fmt.Fprint(w, "-----------------------------------\n")
	w.Flush()

	return b.String()
}

func fprint(w io.Writer, iface interface{}) {
	value := reflect.ValueOf(iface)
	if value.Kind() != reflect.Struct {
		value = value.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Kind() == reflect.Struct {
			iface := field.Interface()
			fprint(w, iface)

			continue
		}

		typeField := value.Type().Field(i)
		fmt.Fprintf(w, "%s\t\033[0m%v\t\033[1;34m%s\033[0m \033[1;92m`%s`\033[0m\n", typeField.Name, field.Interface(), field.Type().String(), typeField.Tag)
	}
}
