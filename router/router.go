package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/dao"
	"test/mysql"
	"test/service"
)

var keyword string

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/js", "vue/frontend/dist/js")
	router.Static("/css", "vue/frontend/dist/css")
	router.Static("/fonts", "vue/frontend/dist/fonts")
	router.LoadHTMLFiles("vue/frontend/dist/index.html")
	router.Static("/img", "vue/frontend/dist/img")

	router.GET("/search", func(c *gin.Context) {
		//option := c.Query("option")
		option := "title"
		keyword, _ = c.GetQuery("keyword")
		//if !has {
		//	print("error")
		//}
		result := dao.SearchByKeyword(option, keyword)
		c.JSON(http.StatusOK, result)

	})

	router.GET("/solrSearch", func(c *gin.Context) {
		keyword, _ = c.GetQuery("keyword")
		//if !has {
		//	print("error")
		//}
		result := service.SolrSearchByKeyword(keyword)
		c.JSON(http.StatusOK, result)
	})

	router.GET("/elasticSearch", func(c *gin.Context) {
		keyword, _ = c.GetQuery("keyword")
		result := service.ElasticSearchByKeyword(keyword)
		c.JSON(http.StatusOK, result)
	})

	router.GET("/geo", func(c *gin.Context) {
		result := dao.GeoMap()
		c.JSON(200, result)
	})

	router.GET("/bar", func(c *gin.Context) {
		bar := dao.JournalNumberGraph(keyword)
		c.JSON(200, bar)
	})

	router.GET("/authorbar", func(c *gin.Context) {
		bar := dao.AuthorPaperGraph(keyword)
		c.JSON(200, bar)
	})

	router.GET("/geoMap", func(c *gin.Context) {
		geoMap := dao.GeoMap()
		c.JSON(200, geoMap)
	})

	router.GET("/linefold", func(c *gin.Context) {
		linefold := dao.PaperNumberGraph(keyword)
		c.JSON(http.StatusOK, linefold)
	})

	router.GET("/topk", func(c *gin.Context) {
		radar := dao.RelatedTopicPie(keyword)
		c.JSON(http.StatusOK, radar)
	})

	router.GET("/wordcloud", func(c *gin.Context) {
		wordcloud := dao.WordCloud()
		c.JSON(http.StatusOK, wordcloud)
	})

	router.GET("/radar", func(c *gin.Context) {
		radar := dao.RadarGraph(keyword)
		c.JSON(http.StatusOK, radar)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"message": "sss",
		})
	})
	router.GET("/displayAuthorList/:id", func(c *gin.Context) {
		name := c.Param("id")
		//name = strings.Trim(name,)
		fmt.Println(name)
		authors := mysql.AuthorList(name)
		c.JSON(200, authors)
	})
	router.GET("/displayAuthor/:id", func(c *gin.Context) {
		name := c.Param("id")
		//name = strings.Trim(name,)
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
	router.GET("/paperTitles", func(c *gin.Context) {
		titleNames := mysql.AllTitleNames()
		c.JSON(200, titleNames)
	})
	return router
}
