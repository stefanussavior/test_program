package controllers

import (
	"test_netwerk/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BuatInvestasiInput struct {
	gorm.Model
	Nama              string
	JenisKelamin      string
	Usia              int16
	Email             string
	Perokok           string
	Nominal           int32
	PeriodePembayaran string
	MetodeBayar       string
	LamaInvestasi     int16
}

type UbahInvestasiInput struct {
	gorm.Model
	Nama              string
	JenisKelamin      string
	Usia              int16
	Email             string
	Perokok           string
	Nominal           int32
	PeriodePembayaran string
	MetodeBayar       string
	LamaInvestasi     int16
}

func GetDataPengguna(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var investasi []models.Investasi

	if err := db.Find(&investasi).Error; err != nil {
		c.JSON(404, gin.H{"Message": "error"})
		return
	}
	c.JSON(200, gin.H{"Success": investasi})
}

func BuatDataPengguna(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Input BuatInvestasiInput
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(404, gin.H{"Message": "Error"})
		return
	}
	Investasi := models.Investasi{
		Nama:              Input.Nama,
		JenisKelamin:      Input.JenisKelamin,
		Usia:              Input.Usia,
		Email:             Input.Email,
		Perokok:           Input.Perokok,
		Nominal:           Input.Nominal,
		PeriodePembayaran: Input.PeriodePembayaran,
		MetodeBayar:       Input.MetodeBayar,
		LamaInvestasi:     Input.LamaInvestasi,
	}
	db.Create(&Investasi)
	c.JSON(200, gin.H{"Success": Investasi})
}

func CariDataInputNasabah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Investasi models.Investasi

	if err := db.Where("id = ?", c.Param("id")).First(&Investasi).Error; err != nil {
		c.JSON(404, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(202, gin.H{"Success": Investasi})
}

func UpdateInvestasiInput(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Investasi models.Investasi

	if err := db.Where("id = ?", c.Param("id")).First(&Investasi).Error; err != nil {
		c.JSON(404, gin.H{"Message": err.Error()})
		return
	}

	var UpdateInput UbahInvestasiInput
	if err := c.ShouldBindJSON(&UpdateInput); err != nil {
		c.JSON(404, gin.H{"Message": err})
		return
	}
	db.Model(&Investasi).Updates(UpdateInput)
	c.JSON(200, gin.H{"Success": Investasi})
}

func HapusDataPengguna(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Investasi models.Investasi

	if err := db.Where("id = ?", c.Param("id")).First(&Investasi).Error; err != nil {
		c.JSON(404, gin.H{"Message": "Error"})
		return
	}
	db.Delete(&Investasi)
	c.JSON(200, gin.H{"Message": "Data Berhasil Dihapus"})
}

func PertambahanInvestasi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Investasi models.Investasi

	if err := db.Raw("SELECT id, IF(perokok != 'ya', nominal, nominal + 1000) AS hasil_nominal FROM investasis").Find(&Investasi).Error; err != nil {
		c.JSON(404, gin.H{"Message": err.Error})
		return
	}
	c.JSON(202, gin.H{"Success": Investasi})
}
