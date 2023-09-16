package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func AnswerGet(c *gin.Context) {
	answerId, err := strconv.Atoi(c.Param("answer_id"))
	if err != nil {
		fmt.Errorf("Cannot get user id")
	}
}

func AnswerPost(c *gin.Context) {
	title := c.PostForm("user_id")
	themeId := c.PostForm("theme_id")
	content := c.PostForm("content")
}

func AnswersGetByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Errorf("Cannot get user id")
	}
}
