package tracer

type constants string
type ContextPropagationAction string

const version constants = "v1.1.2"
const tracerName constants = "github.com/tomegathericon/go-utils/tracing/tracer"
const Extract ContextPropagationAction = "extract"
const Inject ContextPropagationAction = "inject"
