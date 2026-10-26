// Managed by GoX v0.3.3

//line nav.gox:1
// Package components holds reusable UI components shared across pages.
package components

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/doors-starter/path"
	"github.com/doors-dev/gox"
)

// TopNav renders the top navigation bar with links to each page section.
// Uses doors.ALink for in-app navigation via the Root path model.
// The third link demonstrates a raw Location link (triggers the 404 fallback).
//line nav.gox:13
func TopNav() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ul"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("li"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line nav.gox:16
				__e = __c.Any(topNavLink{
				model: path.Root{
					Page: path.RootLanding,
				},
				text: "Home",
			}); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("li"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line nav.gox:24
				__e = __c.Any(topNavLink{
				model: path.Root{
					Page: path.RootAbout,
				},
				text: "About",
			}); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("li"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line nav.gox:32
				__e = (doors.ALink{
				Batch: true,
				Model: doors.Location{
					Segments: []string{"any", "other", "path"},
				},
			}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("404"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line nav.gox:40
}

// topNavLink is a reusable nav link component.
// It wraps an anchor with doors.ALink and highlights
// the link when its model matches the current URL.
// Batch applies the navigation's page updates in one step, so the
// current page stays intact until the new one is ready.
type topNavLink struct {
	model any
	text any
}

//line nav.gox:52
func (l topNavLink) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line nav.gox:53
		__e = (doors.ALink{
		Batch: true,
		Active: doors.Active{
			Indicator: doors.IndicateClass("active"),
		},
		Model: l.model,
	}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("a"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line nav.gox:59
				__e = __c.Any(l.text); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line nav.gox:60
}
