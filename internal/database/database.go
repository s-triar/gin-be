package database

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"gin-be/internal/ent"
	"gin-be/internal/ent/migrate"
	"gin-be/internal/tool"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

// Service represents a service that interacts with a database.
type DBInterface interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	CloseEnt() error

	GetDBClientEntTx(ctx context.Context) (*ent.Tx, error)

	GetDBClientEnt() *ent.Client

	RollbackTransaction(tx *ent.Tx, err error) error

	Migrate()

	SeedDB()

	SwapDBEntClient(client *ent.Client)
}

type DBService struct {
	con   string
	db    *sql.DB
	entDB *ent.Client
}

var (
	database   string
	password   string
	username   string
	port       string
	host       string
	dbInstance *DBService
)

func New() DBInterface {

	if dbInstance == nil {
		envApp := tool.NewEnv(nil)
		database = envApp.DB_DATABASE
		password = envApp.DB_PASSWORD
		username = envApp.DB_USERNAME
		port = envApp.DB_PORT
		host = envApp.DB_HOST
	}

	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	log.Default().Printf("connecting to db %s\n", connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres using database/sql: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	dbInstance = &DBService{
		db:    db,
		entDB: client,
		con:   connStr,
	}
	return dbInstance
}

func GetDB() DBInterface {
	if dbInstance != nil {
		return dbInstance
	}
	return New()
}

// hook for testing
func (s *DBService) SwapDBEntClient(client *ent.Client) {
	s.entDB = client
}

func (s *DBService) Migrate() {
	client := s.entDB
	// ============CREATE MIGRATION SQL FILE=============
	today := time.Now()
	today_str := today.Format("2006-01-02--15_04_05")
	path := "migrations"
	path_migration_file := "./internal/database/" + path + "/migrate_" + today_str + ".sql"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalf("create migrate folder: %v", err)
		}
	}

	f, err := os.Create(path_migration_file)
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.WriteTo(ctx, f); err != nil {
		log.Fatalf("failed creating schema resources to migrate.sql: %v", err)
	}
	if err := f.Close(); err != nil {
		fmt.Println("FAILED: Close migration file")
		fmt.Println(err)
	}

	file, err := os.Open(path_migration_file)

	if err != nil {
		fmt.Printf("FAILED: Open migration file %s.sql\n", today_str)
		fmt.Println(err)
	}

	migration_content := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		migration_content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("FAILED: Scanner has error")
		fmt.Println(err)
	}

	if err := file.Close(); err != nil {
		fmt.Printf("FAILED: Close migration file %s.sql\n", today_str)
		fmt.Println(err)
	}

	if migration_content != "" {
		// Run migration.
		// https://entgo.io/docs/schema-fields
		if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(false)); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	if migration_content == "" {
		if err := os.Remove(path_migration_file); err != nil {
			fmt.Printf("FAILED: Remove migration file %s.sql\n", today_str)
			fmt.Println(err)
		}
	}

	// ============END CREATE MIGRATION SQL FILE==========
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *DBService) Health() map[string]string {
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

func (s *DBService) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}

func (s *DBService) CloseEnt() error {
	log.Printf("Disconnected from database via ent: %s", database)
	return s.entDB.Close()
}

func (s *DBService) GetDBClientEntTx(ctx context.Context) (*ent.Tx, error) {
	tx, err := s.entDB.Tx(ctx)
	if err != nil {
		log.Fatalf("database/database.go|GetDBClientTx|error on starting a transaction of ent: %s", database)
		return nil, fmt.Errorf("error on starting a transaction: %w", err)
	}
	return tx, nil
}

func (s *DBService) GetDBClientEnt() *ent.Client {
	return s.entDB
}

func (s *DBService) RollbackTransaction(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		log.Fatalf("database/database.go|rollbackTransaction|error on rollback a transaction of ent: %s", database)
		err = fmt.Errorf("failed Rollback = %w: %v", err, rerr)
	}
	return err
}

func (s *DBService) SeedDB() {
	// seed_shoptype()
}
