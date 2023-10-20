package vector_inspector

import "github.com/koykov/dyntpl"

func init() {
	_ = dyntpl.RegisterPool("jsonvector", &ipool{fmtJSON})
	_ = dyntpl.RegisterPool("xmlvector", &ipool{fmtXML})
	_ = dyntpl.RegisterPool("yamlvector", &ipool{fmtYAML})
	_ = dyntpl.RegisterPool("urlvector", &ipool{fmtURL})
	_ = dyntpl.RegisterPool("halvector", &ipool{fmtHAL})

	dyntpl.RegisterModFnNS("vector", "parseJSON", "", modParseJSON)
	dyntpl.RegisterModFnNS("vector", "parseXML", "", modParseXML)
	dyntpl.RegisterModFnNS("vector", "parseYAML", "", modParseYAML)
	dyntpl.RegisterModFnNS("vector", "parseURL", "", modParseURL)
	dyntpl.RegisterModFnNS("vector", "parseHAL", "", modParseHAL)

	dyntpl.RegisterModFnNS("vector", "coalesce", "", modCoalesce)
}
