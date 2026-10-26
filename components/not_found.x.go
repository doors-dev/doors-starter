// Managed by GoX v0.3.3

//line not_found.gox:1
// Package components holds reusable UI components shared across pages.
package components

import (
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// NotFound renders the 404 fallback page.
// The app uses it as a doors.RouteBind leaf, so it receives the current
// doors.Location value and re-renders whenever the unmatched URL changes. It:
//   - Sets the HTTP response status to 404 (initial render only)
//   - Sets the page title to "Not Found"
//   - Displays the unmatched path
//
// Note: title and status tags can be rendered anywhere in the tree;
// Doors moves <title> into <head> and applies Status to the response.
//line not_found.gox:18
func NotFound(l doors.Location) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line not_found.gox:19
		__e = __c.Any(doors.Status(404)); if __e != nil { return }
		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Not Found"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("h1"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Not Found"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("p"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Page `"); if __e != nil { return }
//line not_found.gox:23
			__e = __c.Any(l.Path()); if __e != nil { return }
			__e = __c.Text("` not found. "); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line not_found.gox:25
}
