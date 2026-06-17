// Managed by GoX v0.2.1

//line nav.gox:1
// Package components holds reusable UI components shared across pages.
package components

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/doors-starter/path"
	"github.com/doors-dev/gox"
)

// TopNav renders the top navigation bar with links to each page section.
// Uses doors.ALink for client-side navigation via the Home path model.
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
//line nav.gox:39
}

// topNavLink is a reusable nav link component.
// It wraps an anchor with doors.ALink and highlights
// the link when its model matches the current URL.
type topNavLink struct {
	model any
	text any
}

//line nav.gox:49
func (l topNavLink) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line nav.gox:50
		__e = (doors.ALink{
		Active: doors.Active{
			Indicator: doors.IndicateClass("active"),
		},
		Model: l.model,
	}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("a"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line nav.gox:55
				__e = __c.Any(l.text); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line nav.gox:56
}
