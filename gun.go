package marksmanship

import (
	"context"
	"time"
)

type gun struct {
	// TODO: implement this type
	metricsClient any               // use what i built at work as inspiration
	impacts       map[string]func() // method-keyed span.End() calls
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
	// Load is used to create a magazine (otel
	// tracer) with the provided name and
	// inject it into the context to be used
	Load(ctx context.Context, name string)
	// Fire is used to start a span and also begin
	// the latency histogram metric. It is intended
	// to be used inside a method that was registered
	// by a call to Load(). To end the span, you simply
	// defer a call to Impact() immediately after you
	// call this method
	Fire(ctx context.Context, now time.Time) Round
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

// Load is used to create a magazine (otel
// tracer) with the provided name and
// inject it into the context to be used
func (g *gun) Load(ctx context.Context, name string) {
	// TODO: here is where you'll create the tracer and
	// TODO: inject it into the context to be used later
	// TODO: on
}

// Fire is used to start a span and also begin
// the latency histogram metric. It is intended
// to be used inside a method that was registered
// by a call to Load(). To end the span, you simply
// defer a call to Impact() immediately after you
// call this method
func (g *gun) Fire(ctx context.Context, now time.Time) Round {
	// TODO: here is where you'll grab the name of
	// TODO: the method that called Fire()

	// TODO: here is where the span is started and the
	// TODO: latency histogram is started
	return &round{
		impact: func() {
			// TODO: span.End() will go here
		},
	}
}
