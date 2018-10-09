package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/satori/go.uuid"
)

var (
	db *gorm.DB
)

func init() {

	//get config
	cnf := NewConfig()
	fmt.Printf("Server started %v - %v\n", cnf.DatabaseType, cnf.ConnectionString)

	//setup log
	flag.Parse()
	NewLog(cnf.LogFile)
	Log.Println(fmt.Sprintf("Server started %v - %v\n", cnf.DatabaseType, cnf.ConnectionString))

	//open a onnection
	var err error
	db, err = gorm.Open(cnf.DatabaseType, cnf.ConnectionString)
	if err != nil {
		panic("failed to connect database")
	}

}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v2")
	{
		v1.GET("/ping", fetchPing)
		v1.GET("/cinema-chain", fetchAllCinemaChains)
		v1.GET("/cinema-chain/:cinemachain/site", fetchSitesByCinemaChain)
		v1.GET("/cinema-chain/:cinemachain/site/:site", fetchShowtimesBySite)
		v1.GET("/site/:site", fetchShowtimesBySite)
	}
	router.Run()
}

type (
	cinemaChain struct {
		ID                         uuid.UUID `json:"id" gorm:"column:Id"`
		CinemaChainConfigurationID uuid.UUID `json:"cinema-chain-configuration-id" gorm:"column:CinemaChainConfigurationId"`
		Name                       string    `json:"name" gorm:"column:Name"`
		OrganisationCode           string    `json:"organisation-code" gorm:"column:OrganisationCode"`
		RegionID                   uuid.UUID `json:"region-id" gorm:"column:RegionId"`
	}

	cinemaChainConfiguration struct {
		ID               uuid.UUID `json:"id" gorm:"column:Id"`
		ClientID         uuid.UUID `json:"client-id" gorm:"column:ClientId"`
		SalesChannel     string    `json:"sales-channel" gorm:"column:SalesChannel"`
		AllowTicketTypes bool      `json:"allow-ticket-types" gorm:"column:AllowTicketTypes"`
	}

	joinResults struct {
		Name             string    `json:"name" gorm:"column:Name"`
		CinemaChainID    uuid.UUID `json:"cinema-chain-id" gorm:"column:Id"`
		SalesChannel     string    `json:"sales-channel" gorm:"column:SalesChannel"`
		AllowTicketTypes bool      `json:"allow-ticket-types" gorm:"column:AllowTicketTypes"`
	}

	site struct {
		ID            uuid.UUID `json:"id" gorm:"column:Id"`
		CinemaChainID uuid.UUID `json:"cinema-chain-id" gorm:"column:CinemaChainId"`
		Name          string    `json:"name" gorm:"column:Name"`
		Status        string    `json:"status" gorm:"column:Status"`
	}

	showtime struct {
		ID                  uuid.UUID `json:"id" gorm:"column:Id"`
		SeatsAvailable      string    `json:"name" gorm:"column:Name"`
		StartTime           string    `json:"start-time" gorm:"column:StartTime"`
		ScreenName          string    `json:"screen-name" gorm:"column:ScreenName"`
		AttributeShortNames string    `json:"attribute-short-names" gorm:"column:AttributeShortNames"`
		SalesChannels       string    `json:"sales-channels" gorm:"column:SalesChannels"`
	}
)

func (showtime) TableName() string {
	return "Showtime"
}

func (site) TableName() string {
	return "Site"
}

func (cinemaChain) TableName() string {
	return "CinemaChain"
}

func (cinemaChainConfiguration) TableName() string {
	return "CinemaChainConfiguration"
}

func fetchPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func fetchAllCinemaChains(c *gin.Context) {
	var res []joinResults
	db.Table("CinemaChainConfiguration").Select("CinemaChainConfiguration.SalesChannel, CinemaChain.Name, CinemaChain.Id").Joins("LEFT JOIN CinemaChain ON CinemaChain.CinemaChainConfigurationId = CinemaChainConfiguration.Id").Scan(&res)

	if len(res) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No cinemachains found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

func fetchSitesByCinemaChain(c *gin.Context) {
	cinemachain := c.Param("cinemachain")

	var res []site
	db.Where("CinemaChainId = ?", cinemachain).Find(&res)

	if len(res) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No sites found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

func fetchShowtimesBySite(c *gin.Context) {
	site := c.Param("site")

	var res []showtime
	db.Where("SiteId = ?", site).Find(&res)

	if len(res) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No showtimes found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}
