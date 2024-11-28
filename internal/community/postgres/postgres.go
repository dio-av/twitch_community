package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	c "twitchApp/internal/community"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

type service struct {
	db *sql.DB
}

func NewCommunityRepository() c.Repository {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Create(ctx context.Context, p *c.Post) (sql.Result, error) {
	q := `INSERT INTO posts(title, content, reactions) VALUES($1, $2, $3);`
	r, err := s.db.Exec(q, p.Title, p.Content, p.Reactions)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (s *service) Get(ctx context.Context, id int) (*c.Post, error) {
	var p c.Post

	q := `SELECT * FROM community_posts WHERE id = $1;`
	r := s.db.QueryRow(q, id)

	if err := r.Scan(&p.Id, &p.Title, &p.Content, &p.Reactions); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, c.ErrNotExist
		}
		return nil, err
	}

	return &p, nil
}

func (s *service) GetByTitle(ctx context.Context, t string) (*c.Post, error) {
	var p c.Post

	q := `SELECT * FROM community_posts WHERE title = $1;`
	r := s.db.QueryRow(q, p.Title)

	if err := r.Scan(&p.Id, &p.Title, &p.Content, &p.Reactions); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, c.ErrNotExist
		}
		return nil, err
	}

	return &p, nil
}

func (s *service) All(ctx context.Context) ([]c.Post, error) {
	var pp []c.Post

	q := `SELECT * FROM community_posts;`

	r, err := s.db.Query(q)
	if err != nil {
		return []c.Post{}, nil
	}

	for r.Next() {
		var p c.Post
		if err := r.Scan(&p.Id, &p.Title, &p.Content, &p.Reactions); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, c.ErrNotExist
			}
			return nil, err
		}
		pp = append(pp, p)
	}
	return pp, nil
}

func (s *service) Update(ctx context.Context, p *c.Post) (sql.Result, error) {
	var np c.Post

	q := `SELECT post FROM community_posts WHERE title = $1;`
	r := s.db.QueryRow(q, p.Title)

	if err := r.Scan(&np.Id, &np.Title, &np.Content, &np.Reactions); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, c.ErrNotExist
		}
		return nil, err
	}

	q = `UPDATE community_posts SET content = $1 WHERE id = $2;`
	result, err := s.db.Exec(q, np.Id, np.Content)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
