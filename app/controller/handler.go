package controller

import (
	"mvc-golang-1/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddAntrianHandler : Menambahkan data antrian
func AddAntrianHandler(c *gin.Context) {
	flag, err := model.AddAntrian()
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// GetAntrianHandler : Mendapatkan data antrian
func GetAntrianHandler(c *gin.Context) {
	flag, resp, err := model.GetAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   resp,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// UpdateAntrianHandler : Melakukan update data antrian
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.UpdateAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// DeleteAntrianHandler : Menghapus Data Antrian
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.DeleteAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "status",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// PageAntrianHandler : Menampilkan Antrian terakhir pada aplikasi web
func PageAntrianHandler(c *gin.Context) {
	flag, result, err := model.GetAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}
