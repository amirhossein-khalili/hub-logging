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

// setupIpStatisticsDB initializes an in-memory SQLite database and creates the ip_statistics table manually.
func setupIpStatisticsDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test DB: %v", err)
	}

	// Create the ip_statistics table using SQLite‑compatible SQL.
	createTableSQL := `
	CREATE TABLE ip_statistics (
		id TEXT PRIMARY KEY,
		source_ip VARCHAR(45) NOT NULL,
		period_start DATETIME NOT NULL,
		period_end DATETIME NOT NULL,
		total_requests INTEGER NOT NULL
	);
	`
	err = db.Exec(createTableSQL).Error
	if err != nil {
		t.Fatalf("failed to create table ip_statistics: %v", err)
	}

	return db
}

// createDummyIPStatistics returns a dummy IPStatistics entity for testing.
// It sets values for all fields required by the table.
func createDummyIPStatistics(id uuid.UUID, start, end time.Time) *entities.IPStatistics {
	return &entities.IPStatistics{
		ID:            id,
		SourceIP:      "127.0.0.1",
		PeriodStart:   start,
		PeriodEnd:     end,
		TotalRequests: 1,
	}
}

// TestCreateAndGetByID tests the Create and GetByID methods of the repository.
func TestCreateAndGetByID(t *testing.T) {
	db := setupIpStatisticsDB(t)
	repo := postgres.NewIPStatisticsRepository(db)
	ctx := context.Background()

	id := uuid.New()
	now := time.Now()
	dummyStats := createDummyIPStatistics(id, now, now.Add(time.Hour))

	// Test Create
	err := repo.Create(ctx, dummyStats)
	assert.NoError(t, err, "Create should not return an error")

	// Test GetByID
	fetched, err := repo.GetByID(ctx, id)
	assert.NoError(t, err, "GetByID should not return an error")
	assert.NotNil(t, fetched, "Fetched entity should not be nil")
	assert.Equal(t, dummyStats.ID, fetched.ID, "Fetched ID should match")
	assert.True(t, dummyStats.PeriodStart.Equal(fetched.PeriodStart), "PeriodStart should match")
	assert.True(t, dummyStats.PeriodEnd.Equal(fetched.PeriodEnd), "PeriodEnd should match")
}

// TestGetByPeriod tests whether querying by a time period returns the correct records.
func TestGetByPeriod(t *testing.T) {
	db := setupIpStatisticsDB(t)
	repo := postgres.NewIPStatisticsRepository(db)
	ctx := context.Background()

	now := time.Now()
	stats1 := createDummyIPStatistics(uuid.New(), now, now.Add(time.Hour))
	stats2 := createDummyIPStatistics(uuid.New(), now.Add(2*time.Hour), now.Add(3*time.Hour))

	// Insert both entities.
	err := repo.Create(ctx, stats1)
	assert.NoError(t, err, "Create stats1 should not return an error")
	err = repo.Create(ctx, stats2)
	assert.NoError(t, err, "Create stats2 should not return an error")

	// Query for a period that should include only stats1.
	start := now.Add(-time.Minute)
	end := now.Add(90 * time.Minute)
	results, err := repo.GetByPeriod(ctx, start, end)
	assert.NoError(t, err, "GetByPeriod should not return an error")
	assert.Equal(t, 1, len(results), "Should return exactly one result")
	assert.Equal(t, stats1.ID, results[0].ID, "Result should match stats1")
}

// TestUpdate tests that an existing record is correctly updated.
func TestUpdate(t *testing.T) {
	db := setupIpStatisticsDB(t)
	repo := postgres.NewIPStatisticsRepository(db)
	ctx := context.Background()

	id := uuid.New()
	now := time.Now()
	dummyStats := createDummyIPStatistics(id, now, now.Add(time.Hour))

	// Create the entity.
	err := repo.Create(ctx, dummyStats)
	assert.NoError(t, err, "Create should not return an error")

	// Update the entity's PeriodEnd.
	newPeriodEnd := now.Add(2 * time.Hour)
	dummyStats.PeriodEnd = newPeriodEnd

	err = repo.Update(ctx, dummyStats)
	assert.NoError(t, err, "Update should not return an error")

	// Fetch the updated entity.
	updated, err := repo.GetByID(ctx, id)
	assert.NoError(t, err, "GetByID should not return an error")
	assert.True(t, updated.PeriodEnd.Equal(newPeriodEnd), "PeriodEnd should be updated")
}
