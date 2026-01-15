package marksmanship

import "context"

type gun struct {
	// TODO: implement this type
	metricsClient any // use what i built at work as inspiration
}

type Gun interface {
	// Calibrate is used for automatically
	// creating metrics for a service (svc);
	// note: svc must be a struct, a pointer
	// to a struct, or an interface
	//
	// TODO: add note about HTTP status
	// TODO: errors and global metrics
	// TODO: here
	Calibrate(svc any)
	// Load is used to create a bullet (otel
	// tracer) with the provided name and
	// inject it into the context to be used
	Load(ctx context.Context, name string)
}

func GrabGun() Gun {
	// TODO: here is where the calls to build the metrics client
	// TODO: and set up the otel exporter will go
	return &gun{}
}

// Calibrate is used for automatically
// creating metrics for a service (svc);
// note: svc must be a struct, a pointer
// to a struct, or an interface
//
// TODO: add note about HTTP status
// TODO: errors and global metrics
// TODO: here
func (g *gun) Calibrate(svc any) {
	// TODO: here is where the method-specific metrics will be
	// TODO: created. Later on, an options pattern will be
	// TODO: implemented so that you can specify that svc
	// TODO: should also spawn HTTP status code metrics, or
	// TODO: be able to register any global metrics
}

// Load is used to create a bullet (otel
// tracer) with the provided name and
// inject it into the context to be used
func (g *gun) Load(ctx context.Context, name string) {
	// TODO: here is where you'll create the tracer and
	// TODO: inject it into the context to be used later
	// TODO: on
}
