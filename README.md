# Forked YAML library with templating support

`flanksource/yaml` is forked from go-yaml and adds templating support


### Getting started

1. Replace `gopkg.in/yaml.v3` with `gopkg.in/flanksource/yaml.v3`

2. Use the custom `!!env` and `!!template` tags anywhere in your YAML and they will be replaced by the parser

```
 ami_name: !!template image-builder-{{ (time.Now).Format "2006-01-02-150405" }}
 eula: !!template {{file.ReadFile \"eula.txt\" }}
 access_key: !!env AWS_ACCESS_KEY_ID
 secret_key: !!env AWS_SECRET_ACCESS_KEY
```
#### !!env
Reads the value from an environment variable

#### !!template
Run a go text template with all the functions from [gomplate](https://docs.gomplate.ca/) added.

### Future Work

* Add support for all of gomplates datsources e.g. Vault, Consul, S3, AWS SSM, Git etc..
* Support emitting full YAML lists / dicts not just values.

### Comparison to other templating tools

#### Text Based Templating (Jinga,envsubst,etc)

* The resulting templates are still 100% valid YAML and editable / parseable by other YAML libraries / editors.

#### YTT / CUE

* Drop-in replacement for `gopkg.in/yaml.v3`
* Custom functions are written in Golang not Starlark/Python
* Gomplate already has a large library of functions for text, crypto, network, file handling, etc
