// Package assets manages embedded static files for the application.
// Files are embedded at compile time using Go's embed directive,
// then served via Doors middleware or used as managed resources.
package assets

import (
	"embed"
	"io/fs"
)

// Style is the embedded CSS content served as a Doors managed resource.
// Doors turns these bytes into a hash-based, cacheable URL at render time.
//
//go:embed style.css
var Style []byte

// static holds all files under the static/ directory (images, etc.).
// Subtree is exposed via Static() for the UseFS middleware.
//
//go:embed static/*
var static embed.FS

// Static returns an fs.FS rooted at the static/ subdirectory,
// suitable for doors.UseFS to serve under the /static URL prefix.
func Static() fs.FS {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	return sub
}
