package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sergot/tibiago/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:mariadb@tcp(db:3306)/tibiago?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.World{}, &models.Boss{}, &models.Bosshunt{}, &models.Participant{})

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/worlds", func(c echo.Context) error {
		var worlds []models.World
		db.Find(&worlds)
		return c.JSON(http.StatusOK, worlds)
	})

	// TODO: this should actually be returned as part of /bosshunts object's boss field
	e.GET("/bosses", func(c echo.Context) error {
		var bosses []models.Boss
		db.Find(&bosses)
		return c.JSON(http.StatusOK, bosses)
	})

	e.GET("/bosshunts/:world", func(c echo.Context) error {
		var bosshunts []models.Bosshunt
		db.Where("world = ?", c.Param("world")).Find(&bosshunts)
		return c.JSON(http.StatusOK, bosshunts)
	})

	e.GET("/bosshunt/:id", func(c echo.Context) error {
		var bosshunt models.Bosshunt
		db.Where("id = ?", c.Param("id")).First(&bosshunt)
		return c.JSON(http.StatusOK, bosshunt)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
