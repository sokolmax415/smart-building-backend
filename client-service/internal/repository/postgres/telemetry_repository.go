package postgres

import (
	"client-service/internal/entity"
	"context"
	"database/sql"
	"log"
	"time"
)

type PostgresTelemetryRepository struct {
	db *sql.DB
}

func NewPostgresTelemetryRepository(path string) (*PostgresTelemetryRepository, error) {
	log.Printf("Creating PostgresTelemetryRepository with connection string %s", path)
	db, err := newDBconnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresTelemetryRepository{db: db}, nil
}

func (rep *PostgresTelemetryRepository) GetLatestTelemetry(ctx context.Context, deviceSn string) (*entity.Telemetry, error) {
	log.Printf("GetLatestTelemetry: deviceSn=%s", deviceSn)

	query := "SELECT device_sn, data, send_time FROM smartbuildingDB.metrics WHERE device_sn=$1 ORDER BY send_time DESC LIMIT 1"
	log.Printf("Executing query: %s", query)

	var telemetry entity.Telemetry
	err := rep.db.QueryRowContext(ctx, query, deviceSn).Scan(&telemetry.DeviceSn, &telemetry.Data, &telemetry.SendTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrTelemetryNotFound
		}
		log.Printf("ERROR IN GetLatestTelemetry: %v", err)
		return nil, entity.ErrGetLatestTelemetry
	}

	log.Printf("GetLatestTelemetry: executed successfully")
	return &telemetry, nil
}

func (rep *PostgresTelemetryRepository) GetTelemetryInRange(ctx context.Context, deviceSn string, from, till time.Time) ([]entity.Telemetry, error) {
	log.Printf("GetTelemetryInRange: starting")

	query := "SELECT device_sn, data, send_time FROM smartbuildingDB.metrics WHERE device_sn=$1 AND send_time BETWEEN $2 AND $3 ORDER BY send_time ASC"
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query, deviceSn, from, till)
	if err != nil {
		log.Printf("ERROR IN GetTelemetryInRange: %v", err)
		return nil, entity.ErrGetTelemetryInRange
	}
	defer rows.Close()

	telemetries := make([]entity.Telemetry, 0)
	for rows.Next() {
		var telemetry entity.Telemetry
		err := rows.Scan(&telemetry.DeviceSn, &telemetry.Data, &telemetry.SendTime)
		if err != nil {
			log.Printf("ERROR IN GetTelemetryInRanget: %v", err)
			return nil, entity.ErrScanTelemetryRow
		}
		telemetries = append(telemetries, telemetry)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetTelemetryInRange: %v", err)
		return nil, entity.ErrScanTelemetryRow
	}

	log.Printf("GetTelemetryInRange: executed successfully")
	return telemetries, nil
}
