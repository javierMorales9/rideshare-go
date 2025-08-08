package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
)

// ---------- Búsqueda simple ----------

// GET /trips?start_location=foo&driver_name=bar&rider_name=baz
func ListTrips(c *gin.Context) {
	q := db.DB.Model(&model.Trip{}).
		Joins("JOIN trip_requests ON trips.trip_request_id = trip_requests.id").
		Preload("Driver").
		Preload("TripRequest.Rider").
		Preload("TripRequest.StartLocation")

	if txt := c.Query("start_location"); txt != "" {
		q = q.Where("trip_requests.start_location_id IN (?)",
			db.DB.
				Select("id").
				Table("locations").
				Where("address ILIKE ?", "%"+txt+"%"))
	}
	if txt := c.Query("driver_name"); txt != "" {
		q = q.Joins("JOIN users AS drivers ON drivers.id = trips.driver_id").
			Where("drivers.first_name ILIKE ?", "%"+txt+"%")
	}
	if txt := c.Query("rider_name"); txt != "" {
		q = q.Joins("JOIN users AS riders ON riders.id = trip_requests.rider_id").
			Where("riders.first_name ILIKE ?", "%"+txt+"%")
	}

	var trips []model.Trip
	if err := q.Find(&trips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// ---------- Show simple ----------

// GET /trips/:id
func ShowTrip(c *gin.Context) {
	var trip model.Trip
	if err := db.DB.Preload("Driver").
		Preload("TripRequest.Rider").
		First(&trip, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Header("Cache-Control", "public, max-age=60") // 1 min
	c.JSON(http.StatusOK, trip)
}

// ---------- Show “details” estilo JSON:API ----------

// GET /trips/:id/details?include=driver&fields[driver]=display_name,average_rating
func TripDetails(c *gin.Context) {
	var trip model.Trip
	if err := db.DB.Preload("Driver").
		Preload("TripRequest.Rider").
		First(&trip, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// Construimos manualmente el JSON según los query-params.
	// No es JSON:API al pie de la letra, pero replica lo importante.
	resp := gin.H{
		"id":   trip.ID,
		"type": "trip",
		"attributes": gin.H{
			"rider_name":  trip.TripRequest.Rider.DisplayName(),
			"driver_name": trip.Driver.DisplayName(),
		},
	}

	// include=driver
	if inc := c.Query("include"); strings.Contains(inc, "driver") {
		driverJSON := gin.H{
			"id":   trip.Driver.ID,
			"type": "driver",
		}
		// fields[driver]=foo,bar
		if f := c.Query("fields[driver]"); f != "" {
			fields := strings.Split(f, ",")
			attr := gin.H{}
			for _, field := range fields {
				switch strings.TrimSpace(field) {
				case "display_name":
					attr["display_name"] = trip.Driver.DisplayName()
				case "average_rating":
					attr["average_rating"] = trip.Driver.AverageRating(db.DB)
				}
			}
			driverJSON["attributes"] = attr
		}
		resp["included"] = []gin.H{driverJSON}
	}

	c.JSON(http.StatusOK, resp)
}

// ---------- Historial del rider ----------

// GET /trips/my?rider_id=1   (sólo completados)
func ListMyTrips(c *gin.Context) {
	riderID := c.Query("rider_id")
	if riderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rider_id required"})
		return
	}
	var trips []model.Trip
	db.DB.
		Preload("Driver").
		Joins("JOIN trip_requests ON trips.trip_request_id = trip_requests.id").
		Where("trips.completed_at IS NOT NULL AND trip_requests.rider_id = ?", riderID).
		Find(&trips)

	c.JSON(http.StatusOK, trips)
}
