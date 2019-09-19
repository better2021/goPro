package control

import (
	"basicGo/config"
	"basicGo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var db = config.Config()

func UserList(c *gin.Context) {
	var users []model.User
	// db.Select("user_name").Find(&users) // 只查询数据中的user_name字段
	// db.Where("sex=?","boy").Find(&users) // 查询sex为boy的数据
	// db.Where("id > ?",3).Find(&users) // 查询id大于3的数据 或者 SELECT * FROM users WHERE id>3;
	// db.Raw("select * from users").Scan(&users) // 原生的sql查询
	// SELECT * FROM users WHERE sex="girl" AND id > 3  从users表中查询sex为girl并且id大于3的数据

	// db.First(&users) // 获取第一条数据
	// db.Last(&users) // 获取最后一条数据

	var count int
	name := c.Query("userName")
	//  根据名字模糊查询数据
	db.Where("user_name LIKE ?", "%"+name+"%").Find(&users).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusOK,
		"data":    users,
		"attr": gin.H{
			"total": count,
		},
	})
}

func UserCreat(c *gin.Context) {
	data := &model.User{}
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	db.Create(data)
	// 原生SQL创建： INSERT INTO users VALUES(NULL,NULL,NULL,"coco",20,"gril",'xxx')
	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusOK,
		"data":    data,
	})
}

func UserUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "-id-")

	// 需要更新的元素
	data := &model.User{}
	err := c.Bind(data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// 根据id查询数据
	db.Model(data).Where("id=?", id).Update(data)
	// UPDATE users SET sex="boy01" WHERE id = 7 查找id为7的数据把sex修改为boy01
	fmt.Println(data, "data")
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"status":  http.StatusOK,
		"data":    data,
	})
}

func UserDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id, "--")

	db.Where("id=?", id).Delete(model.User{})
	// 原生SQL ：DELETE FROM users WHERE id = 5 删除id为5的数据
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"status":  http.StatusOK,
		"data":    id,
	})
}
