// main is the entry point for the Doors web application.
// It creates the app, attaches middleware, and starts the HTTP server.
package main

import (
	"context"
	"net/http"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/doors-starter/assets"
	"github.com/doors-dev/gox"
)

func main() {
	// Create the Doors app with a page factory.
	// The factory runs once per new page instance (first load, reload, new tab,
	// external link — not for dynamic in-app navigation which updates the
	// existing instance's URL in place) and returns the root component.
	app := doors.NewApp(func(ctx context.Context, r doors.Request) gox.Comp {
		return App{}
	})

	// Attach middleware for static file serving before Doors handles requests.
	// Embedded static files in assets/static/ are served under the /static URL prefix.
	app.Use(
		doors.UseFS("/static", assets.Static(), doors.CacheControlStatic),
	)

	// Start the HTTP server. Doors.App implements http.Handler,
	// so it plugs directly into the standard Go HTTP server.
	if err := http.ListenAndServe(":8080", app); err != nil {
		panic(err)
	}
}
