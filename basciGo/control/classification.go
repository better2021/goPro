package control

import (
	"basicGo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ClassList(c *gin.Context) {
	var classifications []model.Classification
	var count int
	name := c.Query("name")

	db.Where("name LIKE ?", "%"+name+"%").Find(&classifications).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"status":  http.StatusOK,
		"data":    classifications,
		"attr": gin.H{
			"total": count,
		},
	})
}

func ClasstCreat(c *gin.Context) {
	data := &model.Classification{}
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

func ClasstUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "--")

	data := &model.Classification{}
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

func ClasstDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "id")

	db.Where("id=?", id).Delete(model.Classification{})

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
