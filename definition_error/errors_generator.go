//go:build generate

package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/recative/recative-backend/public"
	"github.com/samber/lo"
	"golang.org/x/tools/imports"
	"os"
	"text/template"
)

type RawJsonError struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var RawJsonErrors = func() (res []RawJsonError) {
	err := json.Unmarshal(public.RawJsonError, &res)
	if err != nil {
		panic(err)
	}
	return
}()

type TemplateDataStruct struct {
	Code       int
	Name       string
	PublicName string
}
type TemplateData struct {
	Structs []TemplateDataStruct
}

//go:embed error.tmpl
var errorTemplate string

//go:generate go run -tags=generate .
func main() {
	tmpl, err := template.New("error.tmpl").Parse(errorTemplate)
	if err != nil {
		panic(err)
	}

	var templateData = TemplateData{
		Structs: lo.Map(RawJsonErrors, func(rawJsonError RawJsonError, _ int) TemplateDataStruct {
			return TemplateDataStruct{
				Code:       rawJsonError.Code,
				Name:       rawJsonError.Name,
				PublicName: strcase.ToCamel(rawJsonError.Name),
			}
		}),
	}

	buffer := new(bytes.Buffer)
	{

		err = tmpl.Execute(buffer, templateData)
		if err != nil {
			panic(err)
		}
	}

	processedRes, err := imports.Process("", buffer.Bytes(), nil)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("error_generated.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(processedRes)
	if err != nil {
		panic(err)
	}
}
