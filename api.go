package dbenclib

import (
	"context"

	config "github.com/lanwenhong/lgobase/gconfig"
)

type DbConf struct {
	DbConfFile    string
	Dbconf        *config.Gconf
	DbConfFileBuf string
}

func DbConfNew(context.Context, string) *DbConf
func (dbc *DbConf) DbConfReadGroup(string) map[string]string
func (dbc *DbConf) DbConfReadGroupWithCtx(ctx context.Context, group string) map[string]string
