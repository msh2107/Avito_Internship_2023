package v1

import (
	"Avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type segmentRoutes struct {
	segmentService service.Segment
}

func newSegmentRoutes(handler *gin.RouterGroup, segmentService service.Segment) {
	r := &segmentRoutes{
		segmentService: segmentService,
	}
	handler.POST("/", r.create)
	handler.DELETE("/", r.delete)
}

type segmentCreateInput struct {
	Slug string `json:"slug"`
}

//	@Summary		Create segment
//	@Description	Create segment
//	@Tags			segments
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	models.Segment
//	@Failure		400	{object}	response
//	@Failure		500	{object}	response
//	@Router			/v1/segment/ [post]
func (r *segmentRoutes) create(c *gin.Context) {
	var input segmentCreateInput
	if err := c.Bind(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	segment, err := r.segmentService.CreateSegment(input.Slug)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, segment)
	return
}

type segmentDeleteInput struct {
	Slug string `json:"slug"`
}

//	@Summary		Delete segment
//	@Description	Delete segment
//	@Tags			segments
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Segment
//	@Failure		400	{object}	response
//	@Failure		500	{object}	response
//	@Router			/v1/segment/ [delete]
func (r *segmentRoutes) delete(c *gin.Context) {
	var input segmentDeleteInput
	if err := c.Bind(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	segment, err := r.segmentService.DeleteSegment(input.Slug)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, segment)
	return
}
