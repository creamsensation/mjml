package mjml

import "github.com/creamsensation/gox"

func Title(nodes ...gox.Node) gox.Node {
	return gox.CreateShared("title")(nodes...)
}
