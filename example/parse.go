package example

import (
	"fmt"

	clienthint "github.com/hentaiOS-Infrastructure/fasthttp-go-client-hints"
	"github.com/valyala/fasthttp"
)

func Handler(ctx *fasthttp.RequestCtx) {
	clientHints, err := clienthint.Parse(ctx)
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
