package router

import (
	"fmt"
	"test/mysql"

	"github.com/gin-gonic/gin"
)

var keyword string

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/js", "vue/frontend/dist/js")
	router.Static("/css", "vue/frontend/dist/css")
	router.Static("/fonts", "vue/frontend/dist/fonts")
	router.LoadHTMLFiles("vue/frontend/dist/index.html")
	router.Static("/img", "vue/frontend/dist/img")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"message": "sss",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})

	router.GET("/displayAuthorList/:id", func(c *gin.Context) {
		name := c.Param("id")
		fmt.Println(name)
		authors := mysql.AuthorList(name)
		c.JSON(200, authors)
	})
	router.GET("/displayAuthor/:id", func(c *gin.Context) {
		name := c.Param("id")
		fmt.Println(name)
		author := mysql.AuthorTotalData(name)
		c.JSON(200, author)
	})

	router.GET("/displayBioentity/:id", func(c *gin.Context) {
		name := c.Param("id")
		fmt.Println(name)
		bioentity := mysql.BioentityTotalData(name)
		c.JSON(200, bioentity)
	})

	router.GET("/displayInstitution/:id", func(c *gin.Context) {
		name := c.Param("id")
		fmt.Println(name)
		instituion := mysql.InstitutionTotalData(name)
		c.JSON(200, instituion)
	})
	router.GET("/institutionNames", func(c *gin.Context) {
		institutionNames := mysql.AllInstitutionNames()
		c.JSON(200, institutionNames)
	})
	router.GET("/authorNames", func(c *gin.Context) {
		authorNames := mysql.AllAuthorNames()
		c.JSON(200, authorNames)
	})
	router.GET("/bioentityNames", func(c *gin.Context) {
		bioentityNames := mysql.AllBiontityNames()
		c.JSON(200, bioentityNames)
	})
	return router
}
