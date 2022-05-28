package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

// InitSignatureMiddleware InitSignatureMiddleware
func InitSignatureMiddleware() *SignatureMiddleware {
	return &SignatureMiddleware{}
}

// SignatureMiddleware SignatureMiddleware
type SignatureMiddleware struct{}

// Apply Apply
func (sm *SignatureMiddleware) Apply(c *gin.Context) {
	log.Println("this is a signature valid middleware")
	c.Next()
}
