package main
import "github.com/gin-gonic/gin"

import "net/http"

func main() {
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })



    //router.GET("/someGet", getting)
    //router.POST("/somePost", posting)
    //router.PUT("/somePut", putting)
    //router.DELETE("/someDelete", deleting)
    //router.PATCH("/somePatch", patching)
    //router.HEAD("/someHead", head)
    //router.OPTIONS("/someOptions", options)

    // By default it serves on :8080 unless a
    // PORT environment variable was defined.
    //router.Run()
    router.Run(":8200")
}
