package main

import (
	"test_netwerk/controllers"
	"test_netwerk/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.ConnectDB()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "Berhasil"})
	})
	r.POST("/TambahDataNasabah", controllers.BuatDataPengguna)
	r.GET("/LihatDataNasabah", controllers.GetDataPengguna)
	r.GET("/LihatDataNasabah/:id", controllers.CariDataInputNasabah)
	r.PATCH("/UpdateDataNasabah/:id", controllers.UpdateInvestasiInput)
	r.DELETE("/HapusDataNasabah/:id", controllers.HapusDataPengguna)

	Soal1 := r.Group("/api")
	{
		Soal1.GET("/test", controllers.PertambahanInvestasi)
	}

	r.Run(":8080")
}
