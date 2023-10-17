package vector_inspector

import "github.com/koykov/dyntpl"

func init() {
	dyntpl.RegisterModFnNS("vector", "parseJSON", "", modParseJSON)

	_ = dyntpl.RegisterPool("jsonvector", &ipool{fmtJSON})
	_ = dyntpl.RegisterPool("xmlvector", &ipool{fmtXML})
	_ = dyntpl.RegisterPool("yamlvector", &ipool{fmtYAML})
	_ = dyntpl.RegisterPool("urlvector", &ipool{fmtURL})
	_ = dyntpl.RegisterPool("halvector", &ipool{fmtHAL})
}
