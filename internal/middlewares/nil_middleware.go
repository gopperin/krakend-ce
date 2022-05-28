package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

// InitNilMiddleware InitNilMiddleware
func InitNilMiddleware() *NilMiddleware {
	return &NilMiddleware{}
}

// NilMiddleware NilMiddleware
type NilMiddleware struct{}

// Apply Apply
func (nm *NilMiddleware) Apply(c *gin.Context) {
	log.Println("this is a nil valid middleware")
	c.Next()
}
