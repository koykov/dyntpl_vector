package vector_inspector

import (
	"github.com/koykov/dyntpl"
	"github.com/koykov/jsonvector"
)

func modParseJSON(ctx *dyntpl.Ctx, buf *any, val any, args []any) error {
	rvec, err := ctx.AcquireFrom("jsonvector")
	if err != nil {
		return err
	}
	vec := rvec.(*jsonvector.Vector)
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
