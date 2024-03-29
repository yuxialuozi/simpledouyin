package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
)

func FollowerList(c *gin.Context) {
	var followerList []role.Author

	userId := c.Query("user_id")
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		var followerid []uint

		service.Db.Model(&role.Relation{}).Where("to_user_id = ?", userId).Pluck("user_id", &followerid)

		for _, id := range followerid {
			var follower role.Author

			service.Db.Model(&role.Author{}).Where("id = ?", id).First(&follower)

			followerList = append(followerList, follower)
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "找到了",
			"user_list":   followerList,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "你还没有登录",
		})

	}
}
