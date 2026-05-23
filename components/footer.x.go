// Managed by GoX v0.1.33-0.20260521150447-fa2a771ac3e0

//line footer.gox:1
// Package components holds reusable UI components shared across pages.
package components

import "github.com/doors-dev/gox"

// Footer renders the page footer with a horizontal rule and copyright notice.
//line footer.gox:7
func Footer() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.InitVoid("hr"); if __e != nil { return }
		{
		}
		__e = __c.Submit(); if __e != nil { return }
		__e = __c.Text("No rights reserved"); if __e != nil { return }
	return })
//line footer.gox:10
}
