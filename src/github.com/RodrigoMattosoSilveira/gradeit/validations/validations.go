package validations

import "github.com/gin-gonic/gin"

// Bind the HTTP request to the string parameter
//
// Input:   (*gin.Context, string)
//
// Output:  (true, value) able to bind it,  (false, 0) otherwise
//
// TODO Figure out a way to unit test it
func ParseParm(ctx *gin.Context, key string) (bool, string) {
	value := ctx.Param(key)
	if value == "" {
		return false, ""
	} 
	return true, value
}
