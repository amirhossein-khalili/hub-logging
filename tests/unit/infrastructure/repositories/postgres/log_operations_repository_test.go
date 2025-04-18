package postgres_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"hub_logging/external/infrastructure/repositories/postgres"
	"hub_logging/internal/domain/entities"
)

// setupLogOperationsDB creates an in-memory SQLite DB and creates the log_operations table.
func setupLogOperationsDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test DB: %v", err)
	}
	createTableSQL := `
    CREATE TABLE log_operations (
        id TEXT PRIMARY KEY,
        log_message_id TEXT NOT NULL,
        operation TEXT NOT NULL,
        created_at DATETIME NOT NULL
    );
    `
	if err := db.Exec(createTableSQL).Error; err != nil {
		t.Fatalf("failed to create table log_operations: %v", err)
	}
	return db
}

// createDummyLogOperation returns a dummy LogOperations entity.
func createDummyLogOperation(id uuid.UUID, logMessageID uuid.UUID, operation string) entities.LogOperations {
	return entities.LogOperations{
		ID:           id,
		LogMessageID: logMessageID,
		Operation:    operation,
	}
}

func TestSaveAndFindByLogMessageID(t *testing.T) {
	db := setupLogOperationsDB(t)
	repo := postgres.NewLogOperationsRepository(db)

	// Create a context
	ctx := context.Background()

	logMsgID := uuid.New()
	op1 := createDummyLogOperation(uuid.New(), logMsgID, "CREATE")
	op2 := createDummyLogOperation(uuid.New(), logMsgID, "UPDATE")

	// Test Save
	err := repo.Save(ctx, op1) // Pass context here
	assert.NoError(t, err)
	err = repo.Save(ctx, op2) // Pass context here
	assert.NoError(t, err)

	// Test FindByLogMessageID
	ops, err := repo.FindByLogMessageID(ctx, logMsgID.String()) // Pass context here
	assert.NoError(t, err)
	assert.Len(t, ops, 2, "Should return two log operations")
}
