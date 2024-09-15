package example

import (
	clienthint "github.com/hentaiOS-Infrastructure/fasthttp-go-client-hints"
	"github.com/valyala/fasthttp"
)

func Handler2(ctx *fasthttp.RequestCtx) {
	isSupport := clienthint.IsSupportClientHints(ctx)

	if isSupport {
		// ...do something
	}
}
