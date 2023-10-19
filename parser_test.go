package vector_inspector

import (
	"testing"

	"github.com/koykov/dyntpl"
	_ "github.com/koykov/vector_inspector"
)

const (
	json_ = `{"x":{"y":{"z":"foobar"}}}`
	xml_  = `<?xml version="1.0" encoding="UTF-8"?><x><y><z>foobar</z></y></x>`
	yaml_ = `x:
	y:
		z: "foobar"`
	url_ = `http:://127.0.0.1:8080/post?xyz=foobar`
	hal_ = `zh-Hant-cn;q=1, zh-cn;q=0.6, zh;q=0.4`

	tplJSON = `{% ctx data = source|vector::parseJSON().(vector) %}{%= data.x.y.z %}`
	tplXML  = `{% ctx data = source|vector::parseXML().(vector) %}{%= data.x.y.z %}`
	tplYAML = `{% ctx data = source|vector::parseYAML().(vector) %}{%= data.x.y.z %}`
	tplURL  = `{% ctx data = source|vector::parseURL().(vector) %}{%= data.query.xyz %}`
	tplHAL  = `{% ctx data = source|vector::parseHAL().(vector) %}{%= data.0.code %}-{%= data.0.script %}-{%= data.0.region %};q={%= data.0.quality %}`
)

func init_() (err error) {
	var treeJSON, treeXML, treeYAML, treeURL, treeHAL *dyntpl.Tree

	if treeJSON, err = dyntpl.Parse([]byte(tplJSON), false); err != nil {
		return
	}
	dyntpl.RegisterTplKey("json", treeJSON)

	if treeXML, err = dyntpl.Parse([]byte(tplXML), false); err != nil {
		return
	}
	dyntpl.RegisterTplKey("xml", treeXML)

	if treeYAML, err = dyntpl.Parse([]byte(tplYAML), false); err != nil {
		return
	}
	dyntpl.RegisterTplKey("yaml", treeYAML)

	if treeURL, err = dyntpl.Parse([]byte(tplURL), false); err != nil {
		return
	}
	dyntpl.RegisterTplKey("url", treeURL)

	if treeHAL, err = dyntpl.Parse([]byte(tplHAL), false); err != nil {
		return
	}
	dyntpl.RegisterTplKey("hal", treeHAL)

	return
}

func TestParser(t *testing.T) {
	if err := init_(); err != nil {
		t.Error(err)
		return
	}
	t.Run("json", func(t *testing.T) {
		ctx := dyntpl.NewCtx()
		ctx.SetString("source", json_)
		r, err := dyntpl.Render("json", ctx)
		if err != nil {
			t.Error(err)
		}
		if string(r) != "foobar" {
			t.FailNow()
		}
	})
	t.Run("xml", func(t *testing.T) {
		ctx := dyntpl.NewCtx()
		ctx.SetString("source", xml_)
		r, err := dyntpl.Render("xml", ctx)
		if err != nil {
			t.Error(err)
		}
		if string(r) != "foobar" {
			t.FailNow()
		}
	})
	t.Run("yaml", func(t *testing.T) {
		// todo implement me
		_ = yaml_
	})
	t.Run("url", func(t *testing.T) {
		ctx := dyntpl.NewCtx()
		ctx.SetString("source", url_)
		r, err := dyntpl.Render("url", ctx)
		if err != nil {
			t.Error(err)
		}
		if string(r) != "foobar" {
			t.FailNow()
		}
	})
	t.Run("hal", func(t *testing.T) {
		ctx := dyntpl.NewCtx()
		ctx.SetString("source", hal_)
		r, err := dyntpl.Render("hal", ctx)
		if err != nil {
			t.Error(err)
		}
		if string(r) != "zh-Hant-cn;q=1" {
			t.FailNow()
		}
	})
}
