package dbenclib

import (
	"context"

	"C"

	config "github.com/lanwenhong/lgobase/gconfig"
)

type DbConf struct {
	DbConfFile    string
	Dbconf        *config.Gconf
	DbConfFileBuf string
}

func DbConfNew(context.Context, string) *DbConf {
	return nil
}
func (dbc *DbConf) DbConfReadGroup(string) map[string]string {
	return nil
}
func (dbc *DbConf) DbConfReadGroupWithCtx(ctx context.Context, group string) map[string]string {
	return nil
}
