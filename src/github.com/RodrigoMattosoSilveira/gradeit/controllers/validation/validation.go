package validation

import "github.com/gin-gonic/gin"

// Bind the HTTP request id parameter
//
// Input:   *gin.Context
//
// Output:  (true, id) able to bind it,  (false, 0) otherwise
//
// TODO Figure out a way to unit test it
func ParseIdParm(ctx *gin.Context) (string, bool) {
	idParm := ctx.Param("id")
	if idParm == "" {
		return "", false
	} 
	return idParm, true
}
