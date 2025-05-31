package postgres

import (
	"client-service/internal/entity"
	"context"
	"database/sql"
	"log"
)

type PostgresLocationRepository struct {
	db *sql.DB
}

func NewPostgresLocationRepository(path string) (*PostgresLocationRepository, error) {
	log.Printf("Creating PostgresLocationRepository with connection string %s", path)
	db, err := newDBconnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresLocationRepository{db: db}, nil
}

func (rep *PostgresLocationRepository) CreateNewLocation(ctx context.Context, location *entity.Location) (int64, error) {
	log.Printf("CreateNewLocation: parent_id=%v, location_type=%s, location_name=%s", location.ParentId, location.LocationType, location.LocationName)

	query := `INSERT INTO smartbuildingDB.location(parent_id,location_type,location_name,created_at)
	VALUES($1,$2,$3,NOW()) RETURNING location_id`
	log.Printf("Executing query: %s", query)

	var location_id int64
	err := rep.db.QueryRowContext(ctx, query, location.ParentId, location.LocationType, location.LocationName).Scan(&location_id)

	if err != nil {
		log.Printf("ERROR IN CreateNewLocation: %v", err)
		return -1, entity.ErrCreateLocation
	}

	log.Printf("CreateNewLocation: executed successfully")
	return location_id, nil
}

func (rep *PostgresLocationRepository) GetLocationById(ctx context.Context, location_id int64) (*entity.Location, error) {
	log.Printf("GetLocationById: location_id=%d", location_id)

	var location entity.Location
	query := "SELECT location_id, parent_id, location_type, location_name, created_at FROM smartbuildingDB.location WHERE location_id=$1"
	log.Printf("Executing query: %s", query)

	err := rep.db.QueryRowContext(ctx, query, location_id).Scan(&location.LocationId,
		&location.ParentId, &location.LocationType, &location.LocationName, &location.CreatedAt)

	if err != nil {
		log.Printf("ERROR IN GetLocationById: %v", err)
		return nil, entity.ErrGetLocationById
	}

	log.Printf("GetLocationById: executed successfully")
	return &location, nil
}

func (rep *PostgresLocationRepository) IsLocationExist(ctx context.Context, location_id int64) (bool, error) {
	log.Printf("IsLocationExist: location_id=%d", location_id)

	query := "SELECT COUNT(*) FROM smartbuildingDB.location WHERE location_id=$1"
	log.Printf("Executing query: %s", query)

	var countRows int64
	err := rep.db.QueryRowContext(ctx, query, location_id).Scan(&countRows)

	if err != nil {
		log.Printf("ERROR IN IsLocationExist: %v", err)
		return false, entity.ErrCheckLocationExistence
	}

	log.Printf("IsLocationExist: executed successfully")
	return countRows > 0, nil
}

func (rep *PostgresLocationRepository) GetLocationsList(ctx context.Context) ([]entity.Location, error) {
	log.Printf("GetLocationsList: starting")

	query := "SELECT location_id, parent_id, location_type, location_name, created_at FROM smartbuildingDB.location"
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("ERROR IN GetLocationsList: %v", err)
		return nil, entity.ErrGetAllLocations
	}
	defer rows.Close()

	locations := make([]entity.Location, 0)
	for rows.Next() {
		var location entity.Location
		err := rows.Scan(&location.LocationId, &location.ParentId, &location.LocationType, &location.LocationName, &location.CreatedAt)
		if err != nil {
			log.Printf("ERROR IN GetLocationsList: %v", err)
			return nil, entity.ErrScanLocationRow
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetLocationsList: %v", err)
		return nil, entity.ErrScanLocationRow
	}

	log.Printf("GetLocationsList: executed successfully")
	return locations, nil
}

func (rep *PostgresLocationRepository) DeleteLocation(ctx context.Context, location_id int64) error {
	log.Printf("DeleteLocation: location_id=%d", location_id)

	query := "DELETE FROM smartbuildingDB.location WHERE location_id=$1"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, location_id)

	if err != nil {
		log.Printf("ERROR IN DeleteLocation: %v", err)
		return entity.ErrDeleteLocation
	}

	log.Printf("DeleteLocation: executed successfully")
	return nil
}

func (rep *PostgresLocationRepository) GetLocationsListWithoutParent(ctx context.Context) ([]entity.Location, error) {
	log.Printf("GetLocationsListWithoutParent: starting")

	query := `SELECT location_id, parent_id, location_type, location_name, created_at FROM smartbuildingDB.location
	WHERE parent_id IS NULL`
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("ERROR IN GetLocationsListWithoutParent: %v", err)
		return nil, entity.ErrGetLocationsListWithoutParent
	}
	defer rows.Close()

	locations := make([]entity.Location, 0)
	for rows.Next() {
		var location entity.Location
		err := rows.Scan(&location.LocationId, &location.ParentId, &location.LocationType, &location.LocationName, &location.CreatedAt)
		if err != nil {
			log.Printf("ERROR IN GetLocationsListWithoutParent: %v", err)
			return nil, entity.ErrScanLocationRow
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetLocationsListWithoutParent: %v", err)
		return nil, entity.ErrScanLocationRow
	}

	log.Printf("GetLocationsListWithoutParent: executed successfully")
	return locations, nil
}

func (rep *PostgresLocationRepository) GetLocationChildren(ctx context.Context, location_id int64) ([]entity.Location, error) {
	log.Printf("GetLocationChildren: location_id=%d", location_id)

	query := `SELECT location_id, parent_id, location_type, location_name, created_at FROM smartbuildingDB.location
	WHERE parent_id=$1`
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query, location_id)

	if err != nil {
		log.Printf("ERROR IN GetLocationChildren: %v", err)
		return nil, entity.ErrGetLocationChildren
	}
	defer rows.Close()

	locations := make([]entity.Location, 0)
	for rows.Next() {
		var location entity.Location
		err := rows.Scan(&location.LocationId, &location.ParentId, &location.LocationType, &location.LocationName, &location.CreatedAt)
		if err != nil {
			log.Printf("ERROR IN GetLocationChildren: %v", err)
			return nil, entity.ErrScanLocationRow
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetLocationChildren: %v", err)
		return nil, entity.ErrScanLocationRow
	}

	log.Printf("GetLocationChildren: executed successfully")
	return locations, nil

}

func (rep *PostgresLocationRepository) GetPathToLocation(ctx context.Context, location_id int64) ([]entity.Location, error) {
	log.Printf("GetPathToLocation: location_id=%d", location_id)

	query := `WITH RECURSIVE location_path AS (
		SELECT location_id, parent_id, location_type, location_name, created_at, 0 AS depth
		FROM smartbuildingDB.location
		WHERE location_id = $1
		UNION ALL
		SELECT l.location_id, l.parent_id, l.location_type, l.location_name, l.created_at, lp.depth+1
		FROM smartbuildingDB.location l
		INNER JOIN location_path lp ON l.location_id = lp.parent_id)
		SELECT location_id, parent_id, location_type, location_name, created_at FROM location_path ORDER BY depth DESC`
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query, location_id)
	if err != nil {
		log.Printf("ERROR IN GetPathToLocation: %v", err)
		return nil, entity.ErrGetLocationChildren
	}
	defer rows.Close()

	locations := make([]entity.Location, 0)
	for rows.Next() {
		var location entity.Location
		err := rows.Scan(&location.LocationId, &location.ParentId, &location.LocationType, &location.LocationName, &location.CreatedAt)
		if err != nil {
			log.Printf("ERROR IN GetLocationChildren: %v", err)
			return nil, entity.ErrScanLocationRow
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetPathToLocation: %v", err)
		return nil, entity.ErrScanLocationRow
	}

	log.Printf("GetPathToLocation executed successfully")
	return locations, nil
}

func (rep *PostgresLocationRepository) UpdateLocationType(ctx context.Context, locationType string, locationId int64) error {
	log.Printf("UpdateLocationType: locationType=%s, locationId=%d", locationType, locationId)

	query := "UPDATE smartbuildingDB.location SET location_type=$1 WHERE location_id=$2"

	_, err := rep.db.ExecContext(ctx, query, locationType, locationId)

	if err != nil {
		log.Printf("ERROR IN UpdateLocationType: %v", err)
		return entity.ErrUpdateLocationType
	}
	log.Printf("UpdateLocationType: executed successfully")
	return nil
}

func (rep *PostgresLocationRepository) UpdateLocationName(ctx context.Context, locationName string, locationId int64) error {
	log.Printf("UpdateLocationName: locationName=%s, locationId=%d", locationName, locationId)

	query := "UPDATE smartbuildingDB.location SET location_name=$1 WHERE location_id=$2"

	_, err := rep.db.ExecContext(ctx, query, locationName, locationId)

	if err != nil {
		log.Printf("ERROR IN UpdateLocationname: %v", err)
		return entity.ErrUpdateLocationName
	}
	log.Printf("UpdateLocationName: executed successfully")
	return nil
}

func (rep *PostgresLocationRepository) UpdateLocationParentId(ctx context.Context, locationParentId *int64, locationId int64) error {
	log.Printf("UpdateLocationPrentId: locationParentId=%v, locationId=%d", locationParentId, locationId)

	query := "UPDATE smartbuildingDB.location SET parent_id=$1 WHERE location_id=$2"

	_, err := rep.db.ExecContext(ctx, query, locationParentId, locationId)

	if err != nil {
		log.Printf("ERROR IN UpdateLocationPrentId: %v", err)
		return entity.ErrUpdateLocationParentId
	}
	log.Printf("UpdateLocationPrentId: executed successfully")
	return nil
}
