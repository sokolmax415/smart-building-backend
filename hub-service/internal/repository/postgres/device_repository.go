package postgres

import (
	"context"
	"database/sql"
	"hub-service/internal/entity"
	"log"
)

type PostgresDeviceRepository struct {
	db *sql.DB
}

func NewPostgresDeviceRepository(path string) (*PostgresDeviceRepository, error) {
	log.Printf("Creating Postgres Device Repository with connection string: %s", path)

	db, err := newDBConnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresDeviceRepository{db: db}, nil
}

func (rep *PostgresDeviceRepository) RegisterDevice(ctx context.Context, device *entity.Device) error {
	log.Printf("RegisterDevice: device_sn=%s, hub_sn=%s, device_type=%s, device_name=%s, fw_version=%s", device.DeviceSn, device.HubSn, device.DeviceType, device.DeviceName, device.FwVersion)

	query := `INSERT INTO smartbuildingDB.devices(device_sn, hub_sn, device_type, device_name, last_ping_time, fw_version, created_at) 
	VALUES($1, $2, $3, $4, NOW(), $5, NOW()) ON CONFLICT(device_sn) DO UPDATE SET
	hub_sn = EXCLUDED.hub_sn,
	device_type = EXCLUDED.device_type,
	device_name = EXCLUDED.device_name,
	last_ping_time = NOW(),
	fw_version = EXCLUDED.fw_version`
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, device.DeviceSn, device.HubSn, device.DeviceType, device.DeviceName, device.FwVersion)
	if err != nil {
		log.Printf("ERROR IN RegisterDevice: %v", err)
		return entity.ErrRegisterDevice
	}

	log.Printf("RegisterDevice: executed successfully")
	return nil
}

func (rep *PostgresDeviceRepository) SaveTelemetry(ctx context.Context, telemetry *entity.Telemetry) error {
	log.Printf("SaveTelemetry: device_sn=%s", telemetry.DeviceSn)

	query := "INSERT INTO smartbuildingDB.metrics(device_sn, data, send_time) VALUES ($1, $2, $3)"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, telemetry.DeviceSn, telemetry.Data, telemetry.SendTime)
	if err != nil {
		log.Printf("ERROR IN SaveTelemetry: %v", err)
		return entity.ErrSaveTelemetry
	}

	log.Printf("SaveTelemetry: executed successfully")
	return nil
}

func (rep *PostgresDeviceRepository) IsDeviceExist(ctx context.Context, sn string) (bool, error) {
	log.Printf("IsDeviceExist: device_sn=%s", sn)

	var countRows int
	query := "SELECT COUNT(*) FROM smartbuildingDB.devices WHERE device_sn = $1"
	log.Printf("Exectuing query: %s", query)

	err := rep.db.QueryRowContext(ctx, query, sn).Scan(&countRows)
	if err != nil {
		log.Printf("ERROR IN IsDeviceExist: %v", err)
		return false, entity.ErrCheckHubExistence
	}

	log.Printf("IsDeviceExist: executed successfully")
	return countRows > 0, nil
}
