// Managed by GoX v0.1.33-0.20260521150447-fa2a771ac3e0

//line landing.gox:1
// Package root contains the route handler and page components for the "/" section.
package root

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/doors-starter/path"
	"github.com/doors-dev/gox"
)

// Landing renders the home/landing page (matched on "/").
// Receives the writable path source so it could update the URL if needed.
//line landing.gox:12
func Landing(p doors.Source[path.Root]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Landing"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("h1"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Landing"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line landing.gox:16
		__e = __c.Any(counter{
		state: doors.NewSource(0),
	}); if __e != nil { return }
	return })
//line landing.gox:19
}
