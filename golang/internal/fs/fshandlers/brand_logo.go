package fshandlers

type BrandLogoHandler struct {
	*fsHandler
}

func NewBrandLogoHandler() *BrandLogoHandler {
	return &BrandLogoHandler{
		&fsHandler{
			BasePath:   rootPath + "/images",
			PublicPath: "/assets/img",
		},
	}
}
