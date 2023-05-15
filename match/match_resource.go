package match

import (
	"github.com/1-samuel/hoot-cal/owl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
	usecase Usecase
}

func NewMatchResource(repository owl.Repository) Resource {
	return Resource{usecase: Usecase{repo: repository}}
}

func (r Resource) GetAll(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	matches, err := r.usecase.FindAll()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, matches)
}

func (r Resource) GetActive(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	match, err := r.usecase.FindActive()

	if err != nil {
		if err == owl.ErrNotFound {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, match)
}

func (r Resource) GetCalendar(c *gin.Context) {
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
