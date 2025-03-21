package repositories

import (
	"context"
	"github.com/hizim-ne/backend/internal/models"
)

// SpeedTestRepository defines the interface for speed test data operations
type SpeedTestRepository interface {
	// SaveResult saves a speed test result to the database
	SaveResult(ctx context.Context, result *models.SpeedTestResult) error

	// GetResultsByUserID retrieves all speed test results for a specific user
	GetResultsByUserID(ctx context.Context, userID string) ([]*models.SpeedTestResult, error)

	// GetResultByID retrieves a specific speed test result by its ID
	GetResultByID(ctx context.Context, id string) (*models.SpeedTestResult, error)

	// DeleteResult deletes a speed test result from the database
	DeleteResult(ctx context.Context, id string) error
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// SaveUser saves a user profile to the database
	SaveUser(ctx context.Context, user *models.UserProfile) error

	// GetUserByID retrieves a user profile by its ID
	GetUserByID(ctx context.Context, id string) (*models.UserProfile, error)

	// GetUserByEmail retrieves a user profile by email
	GetUserByEmail(ctx context.Context, email string) (*models.UserProfile, error)
}