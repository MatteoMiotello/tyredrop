package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pillowww/titw/internal/email/mailer"
	"pillowww/titw/pkg/api/responses"
	"pillowww/titw/pkg/log"
)

type SupportController Controller

func NewSupportController() *SupportController {
	return &SupportController{}
}

type SupportEmailPayload struct {
	Email   string `json:"email" binding:"required" validate:"email"`
	Phone   string `json:"phone,omitempty"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

func (s SupportController) SendSupportEmail(ctx *gin.Context) {
	payload := new(SupportEmailPayload)

	err := ctx.BindJSON(payload)

	if err != nil {
		log.Error("Error binding payload on support email", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Internal server error"})
		return
	}

	sm := mailer.NewSupportMailer()
	err = sm.SendResetEmail(payload.Email, payload.Phone, payload.Name, payload.Message)

	if err != nil {
		log.Error("Error sending support email", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.ErrorResponse{Error: "Internal server error"})
		return
	}
}
