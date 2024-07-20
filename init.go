package vector_inspector

import "github.com/koykov/dyntpl"

func init() {
	_ = dyntpl.RegisterPool("jsonvector", &ipool{fmtJSON})
	_ = dyntpl.RegisterPool("xmlvector", &ipool{fmtXML})
	_ = dyntpl.RegisterPool("yamlvector", &ipool{fmtYAML})
	_ = dyntpl.RegisterPool("urlvector", &ipool{fmtURL})
	_ = dyntpl.RegisterPool("halvector", &ipool{fmtHAL})

	dyntpl.RegisterModFnNS("vector", "parseJSON", "", modParseJSON).
		WithParam("data string", "JSON source to parse.").
		WithDescription("Parse `data` using jsonvector and return vector instance.").
		WithExample(`ctx.SetString("source", "{"x":{"y":{"z":"foobar"}}}")
---
{% ctx data = vector::parseJSON(source).(vector) %}
{%= data.x.y.z %} // foobar`)
	dyntpl.RegisterModFnNS("vector", "parseXML", "", modParseXML).
		WithParam("data string", "XML source to parse.").
		WithDescription("Parse `data` using xmlvector and return vector instance.").
		WithExample(`ctx.SetString("source", "<?xml version="1.0" encoding="UTF-8"?><x><y><z>foobar</z></y></x>")
---
{% ctx data = vector::parseXML(source).(vector) %}
{%= data.x.y.z %} // foobar`)
	dyntpl.RegisterModFnNS("vector", "parseYAML", "", modParseYAML).
		WithParam("data string", "YAML source to parse.").
		WithDescription("Parse `data` using yamlvector and return vector instance.").
		WithNote("CAUTION! Still not implement.")
	dyntpl.RegisterModFnNS("vector", "parseURL", "", modParseURL).
		WithParam("data string", "URL to parse.").
		WithDescription("Parse `data` using urlvector and return vector instance.").
		WithExample(`ctx.SetString("source", "http:://127.0.0.1:8080/post?xyz=foobar")
---
{% ctx data = vector::parseURL(source).(vector) %}
{%= data.query.xyz %} // foobar`)
	dyntpl.RegisterModFnNS("vector", "parseHAL", "", modParseHAL).
		WithParam("data string", "HAL string to parse.").
		WithDescription("Parse `data` using halvector and return vector instance.").
		WithExample(`ctx.SetString("source", "zh-Hant-cn;q=1, zh-cn;q=0.6, zh;q=0.4")
---
{% ctx data = source|vector::parseHAL().(vector) %}
{%= data.0.code %}-{%= data.0.script %}-{%= data.0.region %};q={%= data.0.quality %} // zh-Hant-cn;q=1`)

	dyntpl.RegisterModFnNS("vector", "coalesce", "", modCoalesce).
		WithParam("args ...string", "Keys to read.").
		WithDescription("Return value of first non-empty key in vector object.").
		WithExample(`// source: {"x":{"y":{"z":"foobar"}}}
---
{%= data|vector::coalesce("x.y.z.a.b.c", "x.y.z.a.b", "x.y.z.a", "x.y.z") %} // foobar`)
	dyntpl.RegisterModFnNS("vector", "marshal", "serialize", modMarshal).
		WithParam("data vector|node", "Vector or node object.").
		WithDescription("Serialize vector|node to string according format.").
		WithExample(`// source: {"x":{"y":{"z":"foobar"}}}
---
{%= data|vector::marshal() %}       // {"x":{"y":{"z":"foobar"}}}
{%= vector::marshal(data.x) %}      // {"y":{"z":"foobar"}}
{%= data.x.y|vector::marshal() %}   // {"z":"foobar"}
{%= data.x.y.z|vector::marshal() %} // "foobar"`)

	// todo: get back to NS version in further releases
	dyntpl.RegisterEmptyCheckFn("vector::node", VectorNodeEmptyCheck).
		WithDescription("Check given `value` is empty vector or node.")
}
