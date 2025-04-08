package postgres_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"hub_logging/external/infrastructure/repositories/postgres"
	"hub_logging/internal/domain/entities"
)

// setupLogMessageDB creates an in-memory SQLite DB and creates the log_messages table
// with a schema matching the production model.
func setupLogMessageDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test DB: %v", err)
	}
	createTableSQL := `
	CREATE TABLE log_messages (
		id TEXT PRIMARY KEY,
		status_code INTEGER NOT NULL,
		http_method TEXT NOT NULL,
		route_path TEXT NOT NULL,
		message TEXT NOT NULL,
		user_name TEXT NOT NULL,
		dest_hostname TEXT NOT NULL,
		source_ip TEXT NOT NULL,
		timestamp DATETIME
	);
	`
	if err := db.Exec(createTableSQL).Error; err != nil {
		t.Fatalf("failed to create table log_messages: %v", err)
	}
	return db
}

// createDummyLogMessage returns a dummy LogMessage entity with all required fields.
// Adjust the default values as needed to match your actual business logic.
func createDummyLogMessage(id uuid.UUID, message string) entities.LogMessage {
	return entities.LogMessage{
		ID:         id,
		StatusCode: 0, // default status code
		// HTTPMethod:   "", // empty or a valid HTTP method like "GET"
		RoutePath:    "", // empty or a valid route path
		Message:      message,
		UserName:     "", // empty or a test user name
		DestHostname: "", // empty or a valid hostname
		SourceIP:     "", // empty or a valid IP address
		// Timestamp can be left as zero value if not used in the test.
	}
}

func TestSaveAndFindByID(t *testing.T) {
	db := setupLogMessageDB(t)
	repo := postgres.NewLogMessageRepository(db)

	id := uuid.New()
	msg := createDummyLogMessage(id, "Test log message")

	// Test Save
	err := repo.Save(msg)
	assert.NoError(t, err, "Save should not return an error")

	// Test FindByID
	fetched, err := repo.FindByID(id.String())
	assert.NoError(t, err, "FindByID should not return an error")
	assert.Equal(t, msg.ID, fetched.ID, "IDs should match")
	assert.Equal(t, msg.Message, fetched.Message, "Messages should match")
}

func TestFindAllAndPagination(t *testing.T) {
	db := setupLogMessageDB(t)
	repo := postgres.NewLogMessageRepository(db)

	// Insert multiple messages.
	messages := []entities.LogMessage{
		createDummyLogMessage(uuid.New(), "Message 1"),
		createDummyLogMessage(uuid.New(), "Message 2"),
		createDummyLogMessage(uuid.New(), "Message 3"),
	}
	for _, m := range messages {
		err := repo.Save(m)
		assert.NoError(t, err)
	}

	// Test FindAll
	all, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, all, len(messages), "FindAll should return all messages")

	// Test Pagination (limit=2, offset=1)
	paged, err := repo.FindWithPagination(2, 1)
	assert.NoError(t, err)
	assert.Len(t, paged, 2, "FindWithPagination should return 2 messages")
}

func TestUpdateAndDelete(t *testing.T) {
	db := setupLogMessageDB(t)
	repo := postgres.NewLogMessageRepository(db)

	id := uuid.New()
	msg := createDummyLogMessage(id, "Original message")
	err := repo.Save(msg)
	assert.NoError(t, err)

	// Test Update: modify the message.
	msg.Message = "Updated message"
	err = repo.Update(msg)
	assert.NoError(t, err)

	updated, err := repo.FindByID(id.String())
	assert.NoError(t, err)
	assert.Equal(t, "Updated message", updated.Message, "Message should be updated")

	// Test Delete.
	err = repo.Delete(id.String())
	assert.NoError(t, err)

	_, err = repo.FindByID(id.String())
	assert.Error(t, err, "After deletion, FindByID should return an error")
}
