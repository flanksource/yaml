package yaml_test

import (
	"fmt"

	. "gopkg.in/check.v1"
	"gopkg.in/flanksource/yaml.v3"
)

func init() {
	yaml.AddTemplateFunction("hello", helloFunc)
}

func helloFunc(person string) string {
	return fmt.Sprintf("Hello %s", person)
}

func (s *S) TestTemplateFunctionAlreadyDefined(c *C) {
	err := yaml.AddTemplateFunction("hello", helloFunc)
	c.Assert(err, Not(Equals), nil)
	c.Assert(err.Error(), DeepEquals, "Function hello already defined")
}
