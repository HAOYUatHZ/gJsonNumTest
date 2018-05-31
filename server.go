package main

import (
        "fmt"
        "net/http"
        "runtime"
        "time"
        
       "github.com/gin-gonic/gin"
)

type ReqJson struct {
    I    uint64     `json:"i"`
}

func feedback (c *gin.Context) {
    var reqJson ReqJson
    c.BindJSON(&reqJson)
    out := reqJson.I
    c.JSON(http.StatusOK, gin.H{ "i": out })
}

func main() {
    fmt.Println("Initiating...")

    fmt.Println("Setting cores number as:", runtime.NumCPU())
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    // gin.SetMode(gin.ReleaseMode)
    // gin.SetMode(gin.DebugMode)
    router := gin.Default()

    rootHandler := router.Group("/")
    {
        rootHandler.POST("/", feedback)
    }

    // router.Run(":8080")

    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    3 * time.Minute,
        WriteTimeout:   3 * time.Minute,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}