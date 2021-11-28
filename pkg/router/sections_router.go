package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/6156-DonaldDuck/sections/pkg/config"
	"github.com/6156-DonaldDuck/sections/pkg/model"
	"github.com/6156-DonaldDuck/sections/pkg/router/middleware"
	"github.com/6156-DonaldDuck/sections/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// users
	r.GET("/api/v1/sections", ListSections)
	r.GET("/api/v1/sections/:sectionId", GetSectionById)
	r.POST("/api/v1/sections", CreateSection)
	r.DELETE("/api/v1/sections/:sectionId", DeleteSectionById)

	r.Run(":" + config.Configuration.Port)
}

func ListSections(c *gin.Context) {
	sections, err := service.ListSections()
	if err != nil {
		err = fmt.Errorf("error occurred while listing sections, err=%v", err)
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, sections)
}

func GetSectionById(c *gin.Context) {
	sectionIdStr := c.Param("sectionId")
	sectionId, err := strconv.Atoi(sectionIdStr)
	if err != nil {
		log.Errorf("[router.GetSectionById] failed to parse section id %v, err=%v\n", sectionIdStr, err)
		c.JSON(http.StatusBadRequest, "invalid section id")
		return
	}
	section, err := service.GetSectionById(uint(sectionId))
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, section)
	}
}

func CreateSection(c *gin.Context) {
	section := model.Section{}
	if err := c.ShouldBind(&section); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid section, err=%v", err))
		return
	}
	if id, err := service.CreateSection(section); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error occurred while creating section, err=%v", err))
	} else {
		c.JSON(http.StatusCreated, id)
	}
}

func DeleteSectionById(c *gin.Context) {
	sectionIdStr := c.Param("sectionId")
	sectionId, err := strconv.Atoi(sectionIdStr)
	if err != nil {
		log.Errorf("[router.DeleteSectionById] failed to parse section id %v, err=%v\n", sectionIdStr, err)
		c.JSON(http.StatusBadRequest, "invalid section id")
		return
	}
	if err := service.DeleteSectionById(uint(sectionId)); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error occurred while deleting section, err=%v", err))
	} else {
		c.JSON(http.StatusOK, sectionId)
	}
}
