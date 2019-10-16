package resk

import (
	"github.com/slomo0808/infra"
	"github.com/slomo0808/infra/base"
	_ "imooc.com/resk/apis/web"
	_ "imooc.com/resk/core/accounts"
)

func init() {

	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
