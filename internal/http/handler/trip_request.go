package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
	"github.com/javierMorales9/rideshare-go/internal/service"
)

// Body: { "rider_id":1, "start_address":"X", "end_address":"Y" }
func CreateTripRequest(c *gin.Context) {
	var req struct {
		RiderID      uint   `json:"rider_id"      binding:"required"`
		StartAddress string `json:"start_address" binding:"required"`
		EndAddress   string `json:"end_address"   binding:"required"`
		State        string `json:"state"         binding:"required,len=2"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find-or-create locations (simple, sin geocoding todavía)
	startLoc, err := service.CreateOrGetLocation(c, req.StartAddress, req.State)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	endLoc, err := service.CreateOrGetLocation(c, req.EndAddress, req.State)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// ensure rider exists
	var rider model.User
	if err := db.DB.First(&rider, req.RiderID).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "rider not found"})
		return
	}

	var tripRequest model.TripRequest
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		tripRequest = model.TripRequest{
			RiderID:         rider.ID,
			StartLocationID: startLoc.ID,
			EndLocationID:   endLoc.ID,
		}
		if err := tx.Create(&tripRequest).Error; err != nil {
			return err
		}

		tc := service.TripCreator{TripRequestID: tripRequest.ID}
		if _, err := tc.CreateTrip(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"trip_request_id": tripRequest.ID})
}

// GET /trip_requests/:id
func ShowTripRequest(c *gin.Context) {
	print("Que cojones pasa aqui")
	var tr model.TripRequest
	if err := db.DB.First(&tr, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "not found"})
		return
	}

	var trip model.Trip
	db.DB.Where("trip_request_id = ?", tr.ID).First(&trip)

	c.JSON(http.StatusOK, gin.H{
		"trip_request_id": tr.ID,
		"trip_id":         trip.ID,
	})
}
