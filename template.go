package yaml

import (
	"bytes"
	"fmt"
	"strings"
	gotemplate "text/template"

	"github.com/hairyhenderson/gomplate"
	"github.com/pkg/errors"
)

var templateFunctions = gomplate.Funcs(nil)

func RenderTemplate(templ string) (string, error) {
	tpl := gotemplate.New("")

	tpl, err := tpl.Funcs(templateFunctions).Parse(templ)
	if err != nil {
		return "", fmt.Errorf("invalid template %s: %v", strings.Split(templ, "\n")[0], err)
	}

	unstructured := make(map[string]interface{})

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, unstructured); err != nil {
		return "", fmt.Errorf("error executing template %s: %v", strings.Split(templ, "\n")[0], err)
	}
	return buf.String(), nil
}

func AddTemplateFunction(funcName string, fn interface{}) error {
	_, found := templateFunctions[funcName]
	if found {
		return errors.Errorf("Function %s already defined", funcName)
	}
	templateFunctions[funcName] = fn
	return nil
}
