package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hizim-ne/backend/internal/services"
)

// SpeedTestController handles HTTP requests for speed testing
type SpeedTestController struct {
	speedTestService *services.SpeedTestService
}

// NewSpeedTestController creates a new instance of SpeedTestController
func NewSpeedTestController(speedTestService *services.SpeedTestService) *SpeedTestController {
	return &SpeedTestController{
		speedTestService: speedTestService,
	}
}

// RunTest handles the request to run a speed test
func (c *SpeedTestController) RunTest(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, we would extract the user ID from the authenticated session
	userID := "anonymous" // Default for unauthenticated users

	// Get connection type from query parameters
	isMultiConnection := false
	multiConnParam := r.URL.Query().Get("isMultiConnection")
	if multiConnParam == "true" {
		isMultiConnection = true
	}

	// Get IP information (in a real implementation, this would come from a geolocation service)
	ipInfo := map[string]string{
		"ip":      r.RemoteAddr,
		"isp":     "Example ISP",
		"country": "Turkey",
		"region":  "Istanbul",
	}

	// Run the speed test with connection type parameter
	result, err := c.speedTestService.RunSpeedTest(r.Context(), userID, ipInfo, isMultiConnection)
	if err != nil {
		http.Error(w, "Failed to run speed test: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the result as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetHistory handles the request to get a user's test history
func (c *SpeedTestController) GetHistory(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, we would extract the user ID from the authenticated session
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Get the user's test history
	results, err := c.speedTestService.GetUserTestHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to get test history", http.StatusInternalServerError)
		return
	}

	// Return the results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// DeleteResult handles the request to delete a test result
func (c *SpeedTestController) DeleteResult(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, we would extract the user ID from the authenticated session
	userID := "authenticated_user_id"

	// Get the result ID from the request
	resultID := r.URL.Query().Get("result_id")
	if resultID == "" {
		http.Error(w, "Result ID is required", http.StatusBadRequest)
		return
	}

	// Delete the result
	err := c.speedTestService.DeleteTestResult(r.Context(), resultID, userID)
	if err != nil {
		http.Error(w, "Failed to delete test result", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
