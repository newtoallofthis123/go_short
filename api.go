package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/newtoallofthis123/ranhash"
)

type APIServer struct {
	listAddr string
	store    Store
}

func NewAPIServer(listAddr string, store Store) *APIServer {
	return &APIServer{
		listAddr: listAddr,
		store:    store,
	}
}

func (api *APIServer) handleHomePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (api *APIServer) handleCreateURL(c *gin.Context) {
	url := c.PostForm("url")
	var hash string

	if c.PostForm("hash") != "" {
		hash = c.PostForm("hash")
	} else {
		hash = ranhash.GenerateRandomString(8)
	}

	err := api.store.InsertURL(CreateURLRequest{
		Url:  url,
		Hash: hash,
	})
	if err != nil {
		if _, ok := err.(*HashExistsError); ok {
			c.JSON(400, gin.H{"error": "Hash already exists"})
			return
		}
		c.String(500, "Error inserting URL")
		return
	}

	c.String(200, "<a href='/%s'>/%s</a>", hash, hash)
}

func (api *APIServer) handleRedirect(c *gin.Context) {
	url, err := api.store.GetUrlByHash(c.Param("hash"))
	if err != nil {
		fmt.Println(err)
		c.String(500, "Error retrieving URL")
		return
	}

	c.Redirect(302, url)
}

func (api *APIServer) handleApiCreate(c *gin.Context) {
	var req CreateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	url := req.Url
	var hash string

	if req.Hash != "" {
		hash = req.Hash
	} else {
		hash = ranhash.GenerateRandomString(8)
	}

	err := api.store.InsertURL(CreateURLRequest{
		Url:  url,
		Hash: hash,
	})
	if err != nil {
		if _, ok := err.(*HashExistsError); ok {
			c.JSON(400, gin.H{"error": "Hash already exists"})
			return
		}
		c.JSON(500, gin.H{"error": "Error inserting URL"})
		return
	}

	c.JSON(200, gin.H{"url": url, "hash": hash})
}

func (api *APIServer) handleUrlInfo(c *gin.Context) {
	if c.Param("hash") == "" {
		c.JSON(400, gin.H{"error": "hash is required"})
		return
	}
	url, err := api.store.GetUrlInfoByHash(c.Param("hash"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retrieving URL"})
		return
	}

	c.JSON(200, url)
}

func (api *APIServer) Start() error {
	r := gin.Default()

	apiR := r.Group("/api")

	// CORS
	apiR.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	})

	r.SetTrustedProxies(nil)

	r.Static("/assets", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/version", func(c *gin.Context) {
		c.String(200, "Go Short Server Version 0.1")
	})

	r.GET("/source", func(c *gin.Context) {
		c.Redirect(302, "https://github.com/newtoallofthis123/go_short")
	})

	r.GET("/about", func(c *gin.Context) {
		c.String(200, "This is a simple URL shortener written in Go")
	})

	r.GET("/", api.handleHomePage)
	r.POST("/create", api.handleCreateURL)
	r.GET("/:hash", api.handleRedirect)

	apiR.POST("/create", api.handleApiCreate)
	apiR.GET("/info/:hash", api.handleUrlInfo)

	err := r.Run(api.listAddr)
	if err != nil {
		return err
	}

	return nil
}
