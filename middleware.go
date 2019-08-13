package funcenter

import "context"

type Middleware interface {
	OnInit(context.Context) (context.Context, error)
	OnEnter(context.Context, ...interface{}) context.Context
	OnLeave(context.Context)
	OnClose(context.Context)
}
