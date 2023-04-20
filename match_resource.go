package hoot_cal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MatchResource struct {
	usecase MatchUsecase
}

func (r MatchResource) GetAll(c *gin.Context) {
	matches, err := r.usecase.FindAll()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, matches)
}

func (r MatchResource) GetCalendar(c *gin.Context) {
	cal, err := r.usecase.FindAllCal()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Content-type", "text/calendar")
	c.Header("charset", "utf-8")
	c.Header("Content-Disposition", "inline")
	c.Header("filename", "owl.ics")
	c.Status(http.StatusOK)
	_, err = c.Writer.WriteString(cal.Serialize())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
