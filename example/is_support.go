package example

import (
	"net/http"

	clienthint "github.com/hentaiOS-Infrastructure/fasthttp-go-client-hints"
)

func Handler2(w http.ResponseWriter, r *http.Request) {
	isSupport := clienthint.IsSupportClientHints(&r.Header)

	if isSupport {
		// ...do something
	}
}
