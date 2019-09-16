package control

import (
	"basicGo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ProductList(c *gin.Context) {
	var products []model.Product
	var count int
	name := c.Query("productName")
	menuId, _ := strconv.Atoi(c.Query("menuId"))
	// fmt.Println(menuId, "----")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	fmt.Println(pageNum, pageSize, "----")

	db.Model(&products).Where("product_name LIKE ? AND  menu_id=?", "%"+name+"%", menuId).Count(&count)
	db.Offset((pageNum-1)*pageSize).Limit(pageSize).Order("created_at desc").Where("product_name LIKE ? AND  menu_id=?", "%"+name+"%", menuId).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    products,
		"attr": gin.H{
			"page":  pageNum,
			"total": count,
		},
	})
}

func ProductCreat(c *gin.Context) {
	data := &model.Product{}
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	db.Create(data)
	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusOK,
		"data":    data,
	})
}

func ProductUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "--")

	data := &model.Product{}
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	db.Model(data).Where("id=?", id).Update(data)

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

func ProductDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "id")

	db.Where("id=?", id).Delete(model.Product{})

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
