package main

import (
	"fmt"

	"github.com/evcc-io/evcc/util/templates"
	"github.com/fatih/structs"
	"golang.org/x/exp/slices"
)

func matchParam(p templates.Param, pp []templates.Param) (templates.Param, bool) {
	i := slices.IndexFunc(pp, func(p2 templates.Param) bool {
		return p.Name == p2.Name
	})

	if i < 0 {
		return templates.Param{}, false
	}

	return pp[i], true
}

func fieldsEqual(pf, dpf []*structs.Field) bool {
	for _, dpField := range dpf {
		for _, pField := range pf {
			if dpField.Name() != pField.Name() {
				continue
			}

			if pField.IsZero() {
				continue
			}

			fmt.Println("nonzero", pField.Name())
		}
	}

	return true
}

func main() {
	def := templates.Defaults()

	for _, class := range []string{templates.Charger, templates.Meter, templates.Vehicle} {
		fmt.Println("-", class)

		for _, t := range templates.ByClass(class) {
			fmt.Println("--", t.TemplateDefinition.Template)

			for _, p := range t.Params {
				pFields := structs.Fields(p)

				if dp, ok := matchParam(p, def.Params); ok {
					dpFields := structs.Fields(dp)

					fmt.Println("---", p.Name)

					_ = fieldsEqual(pFields, dpFields)
				}
			}
		}
	}
}
