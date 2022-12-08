package html

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func HomePage() g.Node {
	return Page(
		H1(g.Text(`Does SQLite Scale?!`)),
		P(Class("lead"), g.Text(`Let's find out!`)),
		P(g.Raw(`<a href="https://github.com/maragudk/doessqlitescale.com">See the source for this page at github.com/maragudk/doessqlitescale.com</a>`)),
	)
}
