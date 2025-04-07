package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"hub_logging/external/infrastructure/repositories/postgres"
	"hub_logging/internal/domain/entities"
)

// setupMethodStatusStatisticsDB creates a unique in-memory SQLite DB and a table with the expected schema.
func setupMethodStatusStatisticsDB(t *testing.T) *gorm.DB {
	// Use a unique DSN to avoid interference from other tests.
	dsn := "file:" + uuid.NewString() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test DB: %v", err)
	}
	createTableSQL := `
	CREATE TABLE method_status_statistics (
		id TEXT PRIMARY KEY,
		http_method TEXT NOT NULL,
		period_start DATETIME NOT NULL,
		period_end DATETIME NOT NULL,
		total_requests INTEGER NOT NULL,
		success_count INTEGER NOT NULL,
		error_count INTEGER NOT NULL
	);
	`
	if err := db.Exec(createTableSQL).Error; err != nil {
		t.Fatalf("failed to create table method_status_statistics: %v", err)
	}
	return db
}

// createDummyMethodStatusStatistics returns a dummy MethodStatusStatistics entity.
func createDummyMethodStatusStatistics(id uuid.UUID, periodStart, periodEnd time.Time, totalRequests int) *entities.MethodStatusStatistics {
	return &entities.MethodStatusStatistics{
		ID:            id,
		HttpMethod:    "GET", // Default HTTP method.
		PeriodStart:   periodStart,
		PeriodEnd:     periodEnd,
		TotalRequests: totalRequests,
		SuccessCount:  0, // Default value.
		ErrorCount:    0, // Default value.
	}
}

// TestMethodStatusStatisticsCreateAndGetByID tests the Create and GetByID methods.
func TestMethodStatusStatisticsCreateAndGetByID(t *testing.T) {
	db := setupMethodStatusStatisticsDB(t)
	repo := postgres.NewMethodStatusStatisticsRepository(db)
	ctx := context.Background()

	id := uuid.New()
	now := time.Now()
	dummyStats := createDummyMethodStatusStatistics(id, now, now.Add(time.Hour), 10)

	// Test Create.
	err := repo.Create(ctx, dummyStats)
	assert.NoError(t, err, "Create should not return an error")

	// Test GetByID.
	fetched, err := repo.GetByID(ctx, id)
	assert.NoError(t, err, "GetByID should not return an error")
	assert.NotNil(t, fetched, "Fetched entity should not be nil")
	assert.Equal(t, dummyStats.ID, fetched.ID, "IDs should match")
	assert.Equal(t, dummyStats.HttpMethod, fetched.HttpMethod, "HttpMethod should match")
	assert.True(t, dummyStats.PeriodStart.Equal(fetched.PeriodStart), "PeriodStart should match")
	assert.True(t, dummyStats.PeriodEnd.Equal(fetched.PeriodEnd), "PeriodEnd should match")
	assert.Equal(t, dummyStats.TotalRequests, fetched.TotalRequests, "TotalRequests should match")
	assert.Equal(t, dummyStats.SuccessCount, fetched.SuccessCount, "SuccessCount should match")
	assert.Equal(t, dummyStats.ErrorCount, fetched.ErrorCount, "ErrorCount should match")
}

// TestMethodStatusStatisticsGetByPeriod tests the GetByPeriod method.
func TestMethodStatusStatisticsGetByPeriod(t *testing.T) {
	db := setupMethodStatusStatisticsDB(t)
	repo := postgres.NewMethodStatusStatisticsRepository(db)
	ctx := context.Background()

	now := time.Now()
	stats1 := createDummyMethodStatusStatistics(uuid.New(), now, now.Add(time.Hour), 5)
	stats2 := createDummyMethodStatusStatistics(uuid.New(), now.Add(2*time.Hour), now.Add(3*time.Hour), 8)

	// Insert both records.
	err := repo.Create(ctx, stats1)
	assert.NoError(t, err, "Creating stats1 should not return an error")
	err = repo.Create(ctx, stats2)
	assert.NoError(t, err, "Creating stats2 should not return an error")

	// Query for a period that should include only stats1.
	start := now.Add(-time.Minute)
	end := now.Add(90 * time.Minute)
	results, err := repo.GetByPeriod(ctx, start, end)
	assert.NoError(t, err, "GetByPeriod should not return an error")
	assert.Len(t, results, 1, "Should return exactly one record")
	assert.Equal(t, stats1.ID, results[0].ID, "The returned record should match stats1")
}

// TestMethodStatusStatisticsUpdate tests the Update method.
func TestMethodStatusStatisticsUpdate(t *testing.T) {
	db := setupMethodStatusStatisticsDB(t)
	repo := postgres.NewMethodStatusStatisticsRepository(db)
	ctx := context.Background()

	id := uuid.New()
	now := time.Now()
	stats := createDummyMethodStatusStatistics(id, now, now.Add(time.Hour), 10)
	err := repo.Create(ctx, stats)
	assert.NoError(t, err, "Create should not return an error")

	// Update TotalRequests.
	stats.TotalRequests = 20
	err = repo.Update(ctx, stats)
	assert.NoError(t, err, "Update should not return an error")

	updated, err := repo.GetByID(ctx, id)
	assert.NoError(t, err, "GetByID should not return an error")
	assert.Equal(t, 20, updated.TotalRequests, "TotalRequests should be updated")
}
