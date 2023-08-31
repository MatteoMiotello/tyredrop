package fshandlers

import (
	"mime"
	"pillowww/titw/models"
)

type InvoiceHandler struct {
	*fsHandler
}

func NewInvoiceHandler() *InvoiceHandler {
	return &InvoiceHandler{
		&fsHandler{
			BasePath:   rootPath + "/invoices",
			PublicPath: "/private/invoices",
		},
	}
}

func (u InvoiceHandler) StoreInvoice(billing *models.UserBilling, ct string, number string, stream []byte) (*string, error) {
	ext, err := mime.ExtensionsByType(ct)

	if err != nil {
		return nil, err
	}

	fileName := number + ext[len(ext)-1]

	err = u.WriteFile(fileName, stream)

	if err != nil {
		return nil, err
	}

	return &fileName, nil
}
