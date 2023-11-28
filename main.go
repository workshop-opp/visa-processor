package main

import (
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type appIdRequest struct {
	NumberOfIdentities int64 `json:"numberOfIdentities"`
}

func getApplicationIdentity(c *gin.Context) {
	var req appIdRequest = appIdRequest{NumberOfIdentities: 1}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Writer.Header().Add("X-Warning", "cannot decode body")
	}

	var result map[string]string = make(map[string]string)

	env := os.Environ()
	sort.Strings(env)
	var n int64 = 0
	for _, e := range env {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "APP_ID" {
			n = req.NumberOfIdentities
		}
		if n > 0 {
			result[pair[0]] = pair[1]
			n--
		}
	}

	if len(result) == 0 {
		c.IndentedJSON(http.StatusNotFound, result)
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func main() {
	router := gin.Default()
	router.POST("/getApplicationIdentity", getApplicationIdentity)

	router.Run("localhost:8080")
}
