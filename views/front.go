package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func FrontPage() g.Node {
	return Page(
		"Canvas",
		"/",
		H1(g.Text(`khalid & Sania Welcome You`)),
		P(g.Text(`Do you have problems? We also had problems.`)),
		P(g.Raw(`Then we created this page <em>canvas</em> SK, and now we don't! 😬`)),
	)
}
