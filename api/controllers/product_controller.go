package controllers

import (
	"app/config"
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (con *ProductController) Index(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func (con *ProductController) Get(c *gin.Context) {
	var product models.Product
	config.DB.First(&product, c.Params.ByName("id"))
	c.JSON(http.StatusOK, product)
}

func (con *ProductController) Create(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	if err := config.DB.Create(&product).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, product)
}

func (con *ProductController) Update(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	product.Id, _ = strconv.Atoi(c.Params.ByName("id"))
	if err := config.DB.Updates(&product).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, product)
}

func (con *ProductController) Delete(c *gin.Context) {
	var product models.Product
	if err := config.DB.Delete(&product, c.Params.ByName("id")).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusOK)
}
