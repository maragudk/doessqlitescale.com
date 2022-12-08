package html

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func HomePage() g.Node {
	return Page(
		H1(g.Text(`Does SQLite Scale?!`)),
		P(Class("lead"), g.Text(`Let's find out!`)),

		Footer(Class("mt-32 prose-sm"),
			P(g.Raw(`<a href="https://github.com/maragudk/doessqlitescale.com">Source on Github</a>`)),
			P(g.Raw(`Made in 🇩🇰 by <a href="https://www.maragu.dk/">maragu</a>, maker of <a href="https://www.golang.dk/">online Go courses</a>.`)),
		),
	)
}
