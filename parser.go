package vector_inspector

import (
	"github.com/koykov/dyntpl"
	"github.com/koykov/halvector"
	"github.com/koykov/jsonvector"
	"github.com/koykov/urlvector"
	"github.com/koykov/vector"
	"github.com/koykov/xmlvector"
	"github.com/koykov/yamlvector"
)

func modParseJSON(ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	return parse(fmtJSON, ctx, buf, val, args)
}

func modParseXML(ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	return parse(fmtXML, ctx, buf, val, args)
}

func modParseYAML(ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	return parse(fmtYAML, ctx, buf, val, args)
}

func modParseURL(ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	return parse(fmtURL, ctx, buf, val, args)
}

func modParseHAL(ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	return parse(fmtHAL, ctx, buf, val, args)
}

func parse(fmt_ fmt_, ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	var (
		vec  vector.Interface
		pool string
	)
	switch fmt_ {
	case fmtXML:
		pool = "xmlvector"
	case fmtYAML:
		pool = "yamlvector"
	case fmtURL:
		pool = "urlvector"
	case fmtHAL:
		pool = "halvector"
	case fmtJSON:
		fallthrough
	default:
		pool = "jsonvector"
	}
	vraw, err := ctx.AcquireFrom(pool)
	if err != nil {
		return err
	}
	switch fmt_ {
	case fmtXML:
		vec = vraw.(*xmlvector.Vector)
	case fmtYAML:
		vec = vraw.(*yamlvector.Vector)
	case fmtURL:
		vec = vraw.(*urlvector.Vector)
	case fmtHAL:
		vec = vraw.(*halvector.Vector)
	case fmtJSON:
		fallthrough
	default:
		var ok bool
		if vec, ok = vraw.(*jsonvector.Vector); !ok {
			return ErrUnknownType
		}
	}

	var data []byte
	if val != nil {
		data = ctx.BufAcc.StakeOut().WriteX(val).StakedBytes()
	} else {
		if len(args) == 0 {
			return dyntpl.ErrModPoorArgs
		}
		data = ctx.BufAcc.StakeOut().WriteX(args[0]).StakedBytes()
	}
	if ctx.BufAcc.Error() != nil {
		return err
	}
	if err = vec.Parse(data); err != nil {
		return err
	}
	*buf = vec.Root()

	return nil
}
