package v1

import (
	"Avito/internal/models"
	"Avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userSegmentRoutes struct {
	userSegmentService service.UserSegment
}

func newUserSegmentRoutes(handler *gin.RouterGroup, userSegmentService service.UserSegment) {
	r := &userSegmentRoutes{
		userSegmentService: userSegmentService,
	}

	handler.GET("/", r.getActiveSegmentsById)
	handler.PATCH("/", r.changeSegments)
}

type getActiveSegmentsInput struct {
	Id int `json:"id"`
}

// @Summary		Get Active Segments By ID
// @Description	Get Active Segments By ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]models.Segment
// @Failure		400	{object}	response
// @Failure		500	{object}	response
// @Router			/v1/ [get]
func (r *userSegmentRoutes) getActiveSegmentsById(c *gin.Context) {
	var input getActiveSegmentsInput

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	segments, err := r.userSegmentService.GetActiveSegmentsByUser(models.User{ID: input.Id})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, segments)
}

type changeSegmentsInput struct {
	Id       int      `json:"id"`
	ToAdd    []string `json:"toAdd"`
	ToRemove []string `json:"toRemove"`
}

// @Summary		Change Segments
// @Description	Change User`s Active Segments
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.User
// @Failure		400	{object}	response
// @Failure		500	{object}	response
// @Router			/v1/ [patch]
func (r *userSegmentRoutes) changeSegments(c *gin.Context) {
	var input changeSegmentsInput
	if err := c.Bind(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{ID: input.Id}
	err := r.userSegmentService.ChangeSegments(models.User{ID: input.Id}, input.ToAdd, input.ToRemove)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
