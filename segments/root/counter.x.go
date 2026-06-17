// Managed by GoX v0.2.1

//line counter.gox:1
// Package root contains the route handler and page components for the "/" section.
package root

import (
	"context"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// counter is an interactive demo showing Doors reactive state.
// It holds a Source[int] and wires a click event to Mutate it.
type counter struct {
	state doors.Source[int]
}

//line counter.gox:17
func (c counter) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("button"); if __e != nil { return }
		{
//line counter.gox:20
			__e = __c.Modify(doors.AClick{
			On: func(ctx context.Context, rp doors.RequestPointer) bool {
				c.state.Mutate(ctx, func(i int) int {
					return i + 1
				})
				return false
			},
		}); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Text("Add"); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line counter.gox:32
		__e = __c.Any(c.state.Bind(c.count)); if __e != nil { return }
	return })
//line counter.gox:33
}

// count renders the current value.
// It shows " ← Click!" at zero, then the count after the first click.
//line counter.gox:37
func (c counter) count(i int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line counter.gox:38
		if i == 0 {
			__e = __c.Text(" ← Click!"); if __e != nil { return }
		} else  {
//line counter.gox:41
			__e = __c.Many(" ", i); if __e != nil { return }
		}
	return })
//line counter.gox:43
}
