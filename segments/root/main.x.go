// Managed by GoX v0.2.1

//line main.gox:1
// Package root contains the route handler and page components for the "/" section.
// It is a "segment" — a self-contained group of related pages under one path model.
package root

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/doors-starter/path"
	"github.com/doors-dev/gox"
)

// Main is the router for the Root path model.
// It receives a writable Source[path.Root] from the parent RouteModel,
// then dispatches to the matching page component based on the Page field.
//
// RouteMatch selects a branch by predicate:
//   - RootLanding renders the Landing component
//   - RootAbout renders the About component
//
// The matched fragment only swaps when the active route changes.
//line main.gox:20
func Main(p doors.Source[path.Root]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line main.gox:21
		__e = __c.Any(p.Route(
		doors.RouteMatch(func(p path.Root) bool {
			return p.Page == path.RootLanding
		}).Source(Landing),
		doors.RouteMatch(func(p path.Root) bool {
			return p.Page == path.RootAbout
		}).Source(About),
	)); if __e != nil { return }
	return })
//line main.gox:29
}
