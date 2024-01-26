package mjml

import (
	"context"
	
	"github.com/Boostport/mjml-go"
	
	"github.com/creamsensation/gox"
)

func Render(nodes ...gox.Node) (string, error) {
	return mjml.ToHTML(context.Background(), gox.Render(nodes...))
}
