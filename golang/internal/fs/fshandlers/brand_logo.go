package fshandlers

type BrandLogoHandler struct {
	*fsHandler
}

func NewBrandLogoHandler() *BrandLogoHandler {
	return &BrandLogoHandler{
		&fsHandler{
			BasePath:   "./assets/images",
			PublicPath: "/assets/img",
		},
	}
}
