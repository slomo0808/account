package resk

import (
	_ "github.com/slomo0808/account/apis/web"
	_ "github.com/slomo0808/account/core/accounts"
	"github.com/slomo0808/infra"
	"github.com/slomo0808/infra/base"
)

func init() {

	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
