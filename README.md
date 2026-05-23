# Doors Starter

A starter project for [Doors](https://github.com/doors-dev/doors) — a reactive Go web framework built on [GoX](https://github.com/doors-dev/gox).

## Project Structure

```
.
├── main.go                       # Entry point: creates the Doors app, attaches middleware, starts HTTP server
├── app.gox                       # Root component: HTML shell with head elements, nav, router, and footer
├── go.mod / go.sum               # Go module definition and dependency checksums
│
├── path/                         # URL path models — Go structs describing URL shapes
│   └── root.go                   #   Root path model: "/" and "/about" variants
│
├── components/                   # Reusable UI components shared across pages
│   ├── nav.gox                   #   Top navigation bar with doors.ALink navigation links
│   ├── footer.gox                #   Page footer
│   └── not_found.gox             #   404 fallback page with reactive path display
│
├── segments/                     # Self-contained page groups (one per path model)
│   └── root/                     #   Segment for the "/" section
│       ├── main.gox              #     Routes the Root path model to Landing or About
│       ├── landing.gox           #     Landing page ("/") — includes the counter demo
│       ├── about.gox             #     About page ("/about")
│       └── counter.gox           #     Interactive counter (Source + Bind demo)
│
└── assets/                       # Embedded static files
    ├── embed.go                  #   Go embed directives exposing Style bytes and Static FS
    ├── style.css                 #   Application stylesheet
    └── static/                   #   Raw static files served via UseFS middleware
        └── ico.png               #     Favicon
```

> Generated `.x.go` files are auto-managed by GoX and omitted from this tree.

> The `segments/` layout above is a suggested convention, not a strict rule.
> For larger sections you may want to split further — e.g. separate Go packages per page, or shared `components/` and `drivers/` packages within a segment. Keep the codebase systematically organized in whatever way fits your project.

## How to Extend

### Add a New Page to an Existing Path Model

If you want a new page under the existing `Root` path model (e.g. `/contact`):

1. Add the variant to `path/root.go`:
   ```go
   const (
       RootLanding RootPage = iota
       RootAbout
       RootContact   // new
   )
   ```
   Update the struct tag: `` `path:"/ | /about | /contact"` ``

2. Create a new page component in `segments/root/`:
   ```gox
   elem Contact(p doors.Source[path.Root]) {
       <title>Contact</title>
       <h1>Contact</h1>
   }
   ```

3. Add a new route match in `segments/root/main.gox`:
   ```gox
   doors.RouteMatch(func(p path.Root) bool {
       return p.Page == path.RootContact
   }).Source(Contact),
   ```

4. Link to it in `components/nav.gox`:
   ```gox
   <li>~topNavLink{model: path.Root{Page: path.RootContact}, text: "Contact"}</li>
   ```

### Add a New Path Model (New Section)

To add an entirely new section like `/blog`, `/blog/:id`:

1. Create `path/blog.go`:
   ```go
   package path

   type Blog struct {
       Section BlogSection `path:"/blog | /blog/:ID"`
       ID      string
   }

   type BlogSection int

   const (
       BlogIndex BlogSection = iota
       BlogPost
   )
   ```

2. Create `segments/blog/main.gox`:
   ```gox
   package blog

   import (
       "github.com/doors-dev/doors"
       "github.com/doors-dev/doors-starter/path"
   )

   elem Main(p doors.Source[path.Blog]) {
       ~(p.Route(
           doors.RouteMatch(func(p path.Blog) bool {
               return p.Section == path.BlogIndex
           }).Source(Index),
           doors.RouteDerive(func(p path.Blog) (string, bool) {
               return p.ID, p.Section == path.BlogPost
           }).Beam(Post),
       ))
   }
   ```

    `RouteDerive` matches when `Section == BlogPost` and derives the ID into a `Beam[string]`. The `Post` component receives just the ID instead of the full Blog source, keeping updates narrower.

3. Create the page components:
   ```gox
   // segments/blog/index.gox
   elem Index(p doors.Source[path.Blog]) {
       <h1>Blog Index</h1>
   }

   // segments/blog/post.gox — receives Beam[string] (the ID)
   elem Post(id doors.Beam[string]) {
       <h1>Post: ~(id.Bind(elem(v string) { ~(v) }))</h1>
   }
   ```

4. Add the model to the router in `app.gox`:
   ```gox
   <>
       ~(doors.Route(
           doors.RouteModel(root.Main),
           doors.RouteModel(blog.Main),    // new — tried in order
           doors.RouteLocationDefault(components.NotFound),
       ))
   </>
   ```

### Add a New Static Asset

- **For a cacheable file** (favicon, images): place it in `assets/static/` — it's served via `UseFS` under `/static/`.
- **For a managed resource** (stylesheet, JS): embed it as `[]byte` in `assets/embed.go` and use it inline:
  ```gox
  <script src=(assets.MyScript)></script>
  ```

### Use Tailwind CSS

Create a `style.css` at the project root:

```css
@import "tailwindcss";
```

Build it to `assets/style.css` (which is embedded by `assets/embed.go`):

```bash
npx @tailwindcss/cli -i ./style.css -o ./assets/style.css
```

See [Tailwind CLI installation](https://tailwindcss.com/docs/installation/tailwind-cli) for setup instructions. If using npm, add `node_modules/` to `.gitignore`. A standalone binary is also available (slower but no npm needed).

### Inline Scripts and Styles

It is perfectly fine to use inline `<script>` and `<style>` tags directly. Doors converts them into loadable, cacheable resources at render time. Inline scripts are wrapped in an anonymous async function with access to `$data`, `$hook`, `$fetch`, `$on`, and `$sys`. See the [JavaScript docs](https://doors.dev/docs/15-javascript) for details.

### GoX Workflow

- **Write** GoX source in `.gox` files — this is where templates and components live.
- **Use** `.go` files for plain Go logic (path models, embed directives, helpers).
- **`.x.go` files are auto-generated** — the GoX language server keeps them up to date while you edit `.gox` files. Do not edit them manually.
- Run `gox fmt` then `gox gen` (fmt first, gen second) to format and regenerate from the command line.

### Running the Project

```bash
go run .
```

Open [http://localhost:8080](http://localhost:8080).

> **Safari on localhost**: Doors uses a `Secure` session cookie by default, which Safari rejects on plain HTTP. Use `doors.WithConf(doors.Conf{ServerSessionCookieNoSecure: true})` or set up local HTTPS.

> **Local HTTPS**: Use [mkcert](https://github.com/FiloSottile/mkcert) for trusted self-signed certs — `mkcert -install && mkcert localhost 127.0.0.1 ::1`.

### File Watching

Doors reloads the page automatically when it detects the server has restarted. For file watching during development, install [wgo](https://github.com/bokwoon95/wgo) and run:

```bash
wgo -file=.go -file=.css -file=.js -file=.ts go run .
```

This restarts the server on file changes and the browser reloads automatically.

> Less useful in agentic development — the agent typically restarts the server explicitly after changes.

## Learn More

- [https://doors.dev/docs](https://doors.dev/docs)
- [https://doors.dev/tutorial](https://doors.dev/tutorial)
- [doors repo docs](https://github.com/doors-dev/doors/tree/main/docs)
