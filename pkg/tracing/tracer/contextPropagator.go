package tracer

import "context"

type contextPropagator struct {
	action  ContextPropagationAction
	context context.Context
}

func (c *contextPropagator) Action() ContextPropagationAction {
	return c.action
}

func (c *contextPropagator) SetAction(action ContextPropagationAction) {
	c.action = action
}

func (c *contextPropagator) Context() context.Context {
	return c.context
}

func (c *contextPropagator) SetContext(context context.Context) {
	c.context = context
}

func newContextPropagator() *contextPropagator {
	return &contextPropagator{}
}
