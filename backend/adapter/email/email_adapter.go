package email

import (
	"corporation-site/domain"
	"corporation-site/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailAdapter struct {
	emailService service.EmailService
	validator    domain.Validator
}

func NewEmailAdapter(emailService service.EmailService, validator domain.Validator) *EmailAdapter {
	return &EmailAdapter{
		emailService: emailService,
		validator:    validator,
	}
}

func (a *EmailAdapter) RegisterRoutes(router interface{}) {
	switch r := router.(type) {
	case *gin.Engine:
		a.registerGinRoutes(r)
	default:
		panic("Unsupported router type")
	}
}

func (a *EmailAdapter) registerGinRoutes(router *gin.Engine) {
	router.POST("/email/send", a.handleSendEmail)
}

func (a *EmailAdapter) handleSendEmail(c *gin.Context) {
	var req domain.EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := a.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed"})
		return
	}

	response, err := a.emailService.SendTemplatedEmail(
		c.Request.Context(),
		&req,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to send email",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
