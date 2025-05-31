package postgres

import (
	"client-service/internal/entity"
	"context"
	"database/sql"
	"log"
)

type PostgresHubRepository struct {
	db *sql.DB
}

func NewPostgresHubRepository(path string) (*PostgresHubRepository, error) {
	log.Printf("Creating PostgresHubRepository with connection string %s", path)
	db, err := newDBconnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresHubRepository{db: db}, nil
}

func (rep *PostgresHubRepository) GetHubBySn(ctx context.Context, hubSn string) (*entity.Hub, error) {
	log.Printf("GetHubBySn: hub_sn=%s", hubSn)

	query := "SELECT hub_sn, location_id, uptime, last_ping_time, fw_version, created_at FROM smartbuildingDB.hubs WHERE hub_sn=$1"
	log.Printf("Executing query: %s", query)

	var hub entity.Hub

	err := rep.db.QueryRowContext(ctx, query, hubSn).Scan(&hub.HubSn, &hub.LocationId, &hub.Uptime, &hub.LastPingTime, &hub.FwVersion, &hub.CreatedAt)
	if err != nil {
		log.Printf("ERROR IN GetHubBySn: %v", err)
		return nil, entity.ErrGetHubBySn
	}

	log.Printf("GetHubBySn: executed successfully")
	return &hub, nil
}

func (rep *PostgresHubRepository) IsHubExist(ctx context.Context, hubSn string) (bool, error) {
	log.Printf("IsHubExist: hub_sn=%s", hubSn)

	var countRows int
	query := "SELECT COUNT(*) FROM smartbuildingDB.hubs WHERE hub_sn = $1"
	log.Printf("Executing query: %s", query)

	err := rep.db.QueryRowContext(ctx, query, hubSn).Scan(&countRows)
	if err != nil {
		log.Printf("ERROR IN IsHubExist: %v", err)
		return false, entity.ErrCheckHubExistence
	}

	log.Printf("IsHubExist: executed successfully")
	return countRows > 0, nil
}

func (rep *PostgresHubRepository) DeleteHub(ctx context.Context, hubSn string) error {
	log.Printf("DeleteHub: hubSn=%s", hubSn)

	query := "DELETE FROM smartbuildingDB.hubs WHERE hub_sn=$1"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, hubSn)
	if err != nil {
		return entity.ErrDeleteHub
	}

	log.Printf("DeleteHub: executed successfully")
	return nil
}

func (rep *PostgresHubRepository) GetHubList(ctx context.Context, locationId int64) ([]entity.Hub, error) {
	log.Printf("GetHubList: location_id=%d", locationId)

	query := "SELECT hub_sn, location_id, uptime, last_ping_time, fw_version, created_at FROM smartbuildingDB.hubs WHERE location_id=$1"
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query, locationId)
	if err != nil {
		log.Printf("ERROR IN GetHubList: %v", err)
		return nil, entity.ErrGetAllHubs
	}
	defer rows.Close()

	hubs := make([]entity.Hub, 0)
	for rows.Next() {
		var hub entity.Hub
		err := rows.Scan(&hub.HubSn, &hub.LocationId, &hub.Uptime, &hub.LastPingTime, &hub.FwVersion, &hub.CreatedAt)
		if err != nil {
			log.Printf("ERROR IN GetHubList: %v", err)
			return nil, entity.ErrScanHubRow
		}
		hubs = append(hubs, hub)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetHubList: %v", err)
		return nil, entity.ErrScanHubRow
	}

	log.Printf("GetHubList: executed successfully")
	return hubs, nil
}
