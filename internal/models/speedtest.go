package models

import "time"

// SpeedTestResult represents the result of a speed test
type SpeedTestResult struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	UserID        string    `json:"user_id" bson:"user_id,omitempty"`
	DownloadSpeed float64   `json:"downloadSpeed" bson:"download_speed"`
	UploadSpeed   float64   `json:"uploadSpeed" bson:"upload_speed"`
	Ping          float64   `json:"ping" bson:"ping"`
	Jitter        float64   `json:"jitter" bson:"jitter"`
	ISP           string    `json:"isp" bson:"isp"`
	IPAddress     string    `json:"ipAddress" bson:"ip_address"`
	Country       string    `json:"country" bson:"country"`
	Region        string    `json:"region" bson:"region"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
}

// UserProfile represents a user's profile information
type UserProfile struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
