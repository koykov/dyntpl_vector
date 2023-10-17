# Dyntpl vector bindings

Provide [vector](https://github.com/koykov/vector) and [vector_inspector](https://github.com/koykov/vector_inspector)
features to use in [dyntpl](https://github.com/koykov/dyntpl) templates.

### Usage

```go
package main

import (
	"github.com/koykov/dyntpl"
	_ "github.com/koykov/dyntpl_vector"    // register vector bindings
	_ "github.com/koykov/vector_inspector" // register vector inspector
)

const (
	tpl = `{% ctx data = vector::parseJSON(source).(vector) %}
output: {%= data.x.y.z %}`
	json = `{"x":{"y":{"z":"foobar"}}}`
)

func main() {
	tree, _ := dyntpl.Parse([]byte(tpl), false)
	dyntpl.RegisterTplKey("example", tree)
	ctx := dyntpl.NewCtx()
	ctx.SetString("source", json)
	result, _ := dyntpl.Render("example", ctx)
	println(string(result)) // output: foobar
}
```

