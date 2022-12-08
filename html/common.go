package html

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Page(body ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Does SQLite Scale?!",
		Description: "Yes. Yes it does.",
		Language:    "en",
		Head: []g.Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("wasm_exec.js")),
			Script(Src("sqlite3.js")),
			Script(Src("app.js")),
		},
		Body: []g.Node{Class("dark:bg-gray-900"),
			Container(true,
				Prose(
					g.Group(body),
				),
			),
		},
	})
}

func Container(padY bool, children ...g.Node) g.Node {
	return Div(
		c.Classes{
			"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8": true,
			"py-4 sm:py-6 lg:py-8":                   padY,
		},
		g.Group(children),
	)
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose prose-lg lg:prose-xl xl:prose-2xl dark:prose-invert"), g.Group(children))
}
