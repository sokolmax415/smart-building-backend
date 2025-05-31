package postgres

import (
	"client-service/internal/entity"
	"context"
	"database/sql"
	"log"
)

type PostgresDeviceRepository struct {
	db *sql.DB
}

func NewPostgresDeviceRepository(path string) (*PostgresDeviceRepository, error) {
	log.Printf("Creating PostgresDeviceRepository with connection string %s", path)
	db, err := newDBconnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresDeviceRepository{db: db}, nil
}

func (rep *PostgresDeviceRepository) IsDeviceExist(ctx context.Context, deviceSn string) (bool, error) {
	log.Printf("IsDeviceExist: hub_sn=%s", deviceSn)
	var countRows int
	query := "SELECT COUNT(*) FROM smartbuildingDB.devices WHERE device_sn = $1"
	log.Printf("Executing query: %s", query)

	err := rep.db.QueryRowContext(ctx, query, deviceSn).Scan(&countRows)
	if err != nil {
		log.Printf("ERROR IN IsDeviceExist: %v", err)
		return false, entity.ErrCheckDeviceExistence
	}

	log.Printf("IsDEviceExist: executed successfully")
	return countRows > 0, nil
}

func (rep *PostgresDeviceRepository) GetDeviceCount(ctx context.Context, hubSn string) (int64, error) {
	log.Printf("GetDeviceCount: %s", hubSn)

	query := "SELECT COUNT(*) FROM smartbuildingDB.devices WHERE hub_sn=$1"
	log.Printf("Executing query: %s", query)

	var deviceCount int64
	err := rep.db.QueryRowContext(ctx, query, hubSn).Scan(&deviceCount)

	if err != nil {
		log.Printf("ERROR IN GetDeviceCount: %v", err)
		return -1, entity.ErrGetDeviceCount
	}

	log.Printf("GetDeviceCount: executed successfully")
	return deviceCount, nil
}

func (rep *PostgresDeviceRepository) GetDeviceList(ctx context.Context, hubSn string) ([]entity.Device, error) {
	log.Printf("GetDeviceList: hubSn=%s", hubSn)

	query := "SELECT device_sn, hub_sn, device_type, device_name, last_ping_time, fw_version, created_at FROM smartbuildingDB.devices WHERE hub_sn=$1"
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query, hubSn)
	if err != nil {
		log.Printf("ERROR IN GetDeviceList: %v", err)
		return nil, entity.ErrGetAllDevices
	}
	defer rows.Close()

	devices := make([]entity.Device, 0)
	for rows.Next() {
		var device entity.Device
		err := rows.Scan(&device.DeviceSn, &device.HubSn, &device.DeviceType, &device.DeviceName, &device.LastPingTime, &device.FwVersion, &device.CreatedAt)
		if err != nil {
			log.Printf("ERROR IN GetDeviceList: %v", err)
			return nil, entity.ErrScanDeviceRow
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetDeviceList: %v", err)
		return nil, entity.ErrScanDeviceRow
	}

	log.Printf("GetDeviceList: executed successfully")
	return devices, nil
}
