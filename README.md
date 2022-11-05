# Go Client Hints

## Install

```bash
go get -u github.com/cateiru/go-client-hints
```

## Parse Client Hints

```go
import (
	"fmt"
	"net/http"

	goclienthints "github.com/cateiru/go-client-hints"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	clientHints, err := goclienthints.Parse(&r.Header)
	if err != nil {
		return
	}

	// Sec-CH-UA field
	fmt.Println("Brand: ", clientHints.Brand.Brand)
	fmt.Println("Brand Version: ", clientHints.BrandVersion)
	fmt.Println("Brands: ", clientHints.Brands)

	// Sec-Ch-Ua-Platform filed
	fmt.Println("Platform: ", clientHints.Platform)

	// Sec-CH-UA-Platform-Version filed
	fmt.Println("Platform Version: ", clientHints.PlatformVersion)

	// Sec-Ch-Ua-Mobile filed
	fmt.Println("IsMobile: ", clientHints.IsMobile)

	// Sec-CH-UA-Arch filed
	fmt.Println("Arch: ", clientHints.Architecture)

	// Sec-CH-UA-Bitness filed
	fmt.Println("Bitness: ", clientHints.Bitness)

	// Sec-CH-UA-Model filed
	fmt.Println("Model: ", clientHints.Model)

	// Sec-Ch-Ua-Full-Version filed
	fmt.Println("Full Version: ", clientHints.FullVersion)
}

```

## Check Support Client Hints

```go
import (
	"net/http"

	clienthint "github.com/cateiru/go-client-hints"
)

func Handler2(w http.ResponseWriter, r *http.Request) {
	isSupport := clienthint.IsSupportClientHints(&r.Header)

	if isSupport {
		// ...do something
	}
}
```

## LICENSE

[MIT](./LICENSE)
