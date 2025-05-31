package postgres

import (
	"context"
	"database/sql"
	"hub-service/internal/entity"
	"log"
)

// PostgreHubRepository provides access to the PostgreSQL database
type PostgresHubRepository struct {
	db *sql.DB
}

func NewPostgresHubRepository(path string) (*PostgresHubRepository, error) {
	log.Printf("Creating Postgres Hub Repository with connection string: %s", path)
	db, err := newDBConnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresHubRepository{db: db}, nil
}

func (rep *PostgresHubRepository) CreateOrUpdateHub(ctx context.Context, hub *entity.Hub) error {
	log.Printf("CreateOrUpdateHub: hub_sn=%s, location_id=%d, uptime=%d, fw_version=%s", hub.HubSn, hub.LocationId, hub.Uptime, hub.FwVersion)

	query := `INSERT INTO smartbuildingDB.hubs(hub_sn, location_id, uptime, last_ping_time, fw_version, created_at)
	VALUES ($1, $2, $3, NOW(), $4, NOW()) ON CONFLICT(hub_sn) DO UPDATE SET
	location_id = EXCLUDED.location_id,
	uptime = EXCLUDED.uptime,
	last_ping_time = NOW(),
	fw_version = EXCLUDED.fw_version`
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, hub.HubSn, hub.LocationId, hub.Uptime, hub.FwVersion)
	if err != nil {
		log.Printf("ERROR IN CreateOrUpdateHub: %v", err)
		return entity.ErrRegisterOrUpdateHub
	}

	log.Printf("CreateOrUpdateHub: executed successfully")
	return nil
}

func (rep *PostgresHubRepository) IsHubExist(ctx context.Context, sn string) (bool, error) {
	log.Printf("IsHUbExist: hub_sn=%s", sn)

	var countRows int
	query := "SELECT COUNT(*) FROM smartbuildingDB.hubs WHERE hub_sn = $1"
	log.Printf("Executing query: %s", query)

	err := rep.db.QueryRowContext(ctx, query, sn).Scan(&countRows)
	if err != nil {
		log.Printf("ERROR IN IsHubExist: %v", err)
		return false, entity.ErrCheckHubExistence
	}

	log.Printf("IsHubExist: executed successfully")
	return countRows > 0, nil
}

func (rep *PostgresHubRepository) UpdateHubUptime(ctx context.Context, hub_sn string, uptime int64) error {
	log.Printf("UpdateHupUptime: hub_sn=%s, uptume=%d", hub_sn, uptime)

	query := "UPDATE smartbuildingDB.hubs SET uptime = $1, last_ping_time = NOW() WHERE hub_sn = $2"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, uptime, hub_sn)
	if err != nil {
		log.Printf("ERROR IN UpdateHubUptime: %v", err)
		return entity.ErrUpdateHubUptime
	}

	log.Printf("UpdateHupUptime: executed successfully")
	return nil
}

func (rep *PostgresHubRepository) IsLocationExist(ctx context.Context, location_id int64) (bool, error) {
	log.Printf("IsLocationExist: location_id=%d", location_id)

	query := "SELECT COUNT(*) FROM smartbuildingDB.location WHERE location_id=$1"
	log.Printf("Executing queery: %s", query)

	var countRows int64
	err := rep.db.QueryRowContext(ctx, query, location_id).Scan(&countRows)

	if err != nil {
		log.Printf("ERROR IN IsLocationExist: %v", err)
		return false, entity.ErrCheckLocationExistence
	}

	log.Printf("IsLocationExist: executed successfully")
	return countRows > 0, nil
}
