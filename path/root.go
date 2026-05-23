// Package path defines the URL path models for the application.
// Each struct describes one URL shape: which path variants exist,
// which path segments to decode, and which query parameters matter.
package path

// Root is the path model for the root section of the site.
// It has two variants matched by path pattern:
//   - "/"      -> RootLanding
//   - "/about" -> RootAbout
//
// The int field stores the index of the matched variant,
// which pairs naturally with iota constants.
type Root struct {
	Page RootPage `path:"/ | /about"`
}

type RootPage int

const (
	RootLanding RootPage = iota
	RootAbout
)
