package main

import (
  // import standard libraries
  "net/http"
  "os"

  // import the web framework
  "github.com/gin-gonic/gin"
)


// all go projects require a main function
// this is the function that runs the program
func main() {
  // initialize a gin server
  r := gin.Default()

  // load html files to be rendered
  r.LoadHTMLGlob("*.html")

  // parse through static files to be served
  // static files include our js and css
  r.Static("/public", "public")

  // define the port our server will be running on
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

  // define the route path and response
  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H {
      "HelloMessage": "Go is awesome!",
    })
  })

  // start the server
  r.Run(":" + port)
}
