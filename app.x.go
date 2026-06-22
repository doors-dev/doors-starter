// Managed by GoX v0.2.1

//line app.gox:1
// App is the root component for every page instance.
// It defines the HTML shell: doctype, head elements, and body layout.
// Routing happens inside Main(), not in the page factory.
package main

import (
	"github.com/doors-dev/doors-starter/assets"
	"github.com/doors-dev/doors-starter/components"
	"github.com/doors-dev/doors-starter/segments/root"
	
	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

type App struct{}

//line app.gox:17
func (a App) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Raw("<!doctype html>"); if __e != nil { return }
		__e = __c.Init("html"); if __e != nil { return }
		{
//line app.gox:19
			__e = __c.Set("lang", "en"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("head"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:21
					__e = __c.Set("charset", "utf-8"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("meta"); if __e != nil { return }
				{
//line app.gox:22
					__e = __c.Set("name", "viewport"); if __e != nil { return }
//line app.gox:22
					__e = __c.Set("content", "width=device-width, initial-scale=1"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:25
					__e = __c.Set("rel", "icon"); if __e != nil { return }
//line app.gox:26
					__e = __c.Set("type", "image/png"); if __e != nil { return }
//line app.gox:27
					__e = __c.Set("href", "/static/ico.png"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.InitVoid("link"); if __e != nil { return }
				{
//line app.gox:31
					__e = __c.Set("rel", "stylesheet"); if __e != nil { return }
//line app.gox:32
					__e = __c.Set("href", assets.Style); if __e != nil { return }
//line app.gox:33
					__e = __c.Set("name", "main.css"); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("body"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line app.gox:36
				__e = __c.Any(components.TopNav()); if __e != nil { return }
//line app.gox:39
				__e = __c.Any(doors.Route(
				doors.RouteModel(root.Main),
				doors.RouteDefaultBind(components.NotFound),
			)); if __e != nil { return }
//line app.gox:43
				__e = __c.Any(components.Footer()); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line app.gox:46
}
