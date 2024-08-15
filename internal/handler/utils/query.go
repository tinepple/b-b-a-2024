package utils

import "github.com/gin-gonic/gin"

func GetQueryString(c *gin.Context, key string) string {
	if !c.Request.URL.Query().Has(key) {
		return ""
	}
	return c.Request.URL.Query().Get(key)
}
