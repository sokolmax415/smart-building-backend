# Smart Building Backend System

This project is a backend application for a building management automation system. It is used by an Android application for clients and hubs with devices. The system was originally developed as a university coursework project.

---

## Tech Stack

- **Language**: Go
- **Architecture**: Clean Architecture + Microservices
- **API style**: REST (via `go-chi`)
- **Database**: PostgreSQL 17
- **Auth**: JWT (access & refresh tokens)
- **Password hashing**: bcrypt with per-user salt
- **Containerization**: Docker + docker-compose
- **Testing**: Postman + manual integration testing

---

## Microservices Overview

## `auth-service`

### Description:
This service is responsible for user's authentication and admin's management

### Features:
- User registration and login
- Access and refresh token generation
- Role-based access control
- Admin-level user management

### Endpoints:
- `POST smartbuilding/v1/auth/register` - User registration
- `POST smartbuilding/v1/auth/login` - User authentication 
- `POST smartbuilding/v1/auth/refresh` - Refreshing an expired access token
- `GET smartbuilding/v1/users` - Get a list of all users (available only for the admin role)
- `GET smartbuilding/v1/users/{login}` - Get information about a user by login (available only for the admin role)
- `POST smartbuilding/v1/users` - Create a new user in the system (available only for the admin role)
- `PUT smartbuilding/v1/users/{login}` - Change a user role (available only for the admin role)
- `DELETE smartbuilding/v1/users/{login}` - Delete a user from the system (available only for the admin role)
- `PUT smartbuilding/v1/users/name/{login}` - Change the user's first and last name (available only for the admin role)

---

## `client-service`

### Description:
This service is responsible for managing hierarchical locations (such as buildings, floors, and rooms), assigning IoT devices to them, getting telemetry data from the devices and location's information.

### Features:
- Hierarchical location management (floors and rooms)
- Getting information about location
- Requesting telemetry from devices

### Endpoints:
- `POST smartbuilding/v1/client/locations` – Create a location
- `GET smartbuilding/v1/client/locations/tree` – Get tree of all locations
- `GET smartbuilding/v1/client/locations/{location_id}` - Get brief information about location
-	`PATCH smartbuilding/v1/client/locations/{location_id}` – Update information about location
-	`DELETE smartbuilding/v1/client/locations/{location_id}` – Delete location
-	`GET smartbuilding/v1/client/locations/root` – Get root locations
-	`GET smartbuilding/v1/client/locations/{location_id}/children` – Get list of children of a location
-	`GET smartbuilding/v1/client/locations/{location_id}/details` – Get detailed information about a location (including data about hubs and devices on it)
-	`GET smartbuilding/v1/client/locations/{location_id}/parent` – Get location parent
-	`GET smartbuilding/v1/client/locations/{location_id}/path` – Get path to location
-	`GET smartbuilding/v1/client/devices/{device_sn}/telemetry/latest` – Get latest telemetry from sensor
-	`GET smartbuilding/v1/client/devices/{device_sn}/telemetry?from=timestamp&till=timestamp` – Gettelemetry within time range
-	`GET smartbuilding/v1/client/hubs/{hub_sn}` –  Get hub information
-	`DELETE smartbuilding/v1/client/hubs/{hub_sn}` – Delete hub from system
-	`GET smartbuilding/v1/client/hubs/{hub_sn}/devices` – Get list of devices for a hub
-	`GET smartbuilding/v1/client/hubs/{hub_sn}/count` – Get the number of devices on a hub

---

## `hub-service`

### Description:
This service is responsible for managing devices and storing telemetry by hubs

### Features:
- Hub registration (hub sends the request itself)
- Device registration
- Saving telemetry data from devices
- Ping a hub

### Endpoints:
- `POST smartbuilding/v1/hubs/register` - Create a hub and binding it to the location
- `POST smartbuilding/v1/hubs/devices/register` - Create or updating a device
- `POST smartbuilding/v1/hubs/devices/telemetry` - Send telemetry from the sensor
- `POST smartbuilding/v1/hubs/ping` - Ping a hub (update its working time since the last launch)


## Database Structure

Indexes are used to speed up work with the database.

### Scheme: `userDB` 
- The `users` table (id, firstname, lastname, login, password_hash, role_id, registration_time) stores information about registered users (name, last name, login and hashed password) and their further authentication in the system
- The `roles` table (role_id, role, permissions) stores all roles and permissions to perform operations associated with these roles

### Scheme: `smartbuildingDB`
- The `location` table (location_id, parent_id, location_type, location_name, created_at) stores information about created locations
- The `hubs` table (hub_id, hub_sn, location_id, uptime, last_ping_time, fw_version, created_at) stores data about hubs
- The `devices` table (device_id, device_sn, hub_sn, device_type, device_name, last_ping_time, fw_version, created_at) stores information about devices
- The `metrics` table (metric_id, device_sn, data, send_time) stores data received from sensors and other devices(data in JSONB)


## Security

- **JWT-based authentication**
  - Access token: valid for 10 minutes
  - Refresh token: valid for 7 days
- **Role-based access control**
  - Admins can manage users
  - Users can only see their own data from the devices and get information about hubs
- **Secure password hashing**
  - bcrypt with per-password salt



## Testing

- API tested using Postman collections
- Manual integration testing of token lifecycle
- Simulated telemetry stream tested with:
  - 1 request per second per device
  - 3 hours of continuous input (≥ 10,000 records)
- Edge cases tested: expired tokens, missing auth, invalid payloads

## Environment Variables

The project uses the following environment variables. You can define them in a `.env` file in the root of the project.  
There is an example in the file .env.example

### PostgreSQL Configuration

- `POSTGRES_USER` — Username for the PostgreSQL database (e.g. `userHSE`)
- `POSTGRES_PASSWORD` — Password for the PostgreSQL database (e.g. `qwerty123`)
- `POSTGRES_DB` — Name of the PostgreSQL database (e.g. `buildingDB`)
- `POSTGRES_PORT` — Port for PostgreSQL (default: `5432`)

### Service Ports

- `AUTH_SERVICE_PORT` — Port for the auth-service (e.g. `8080`)
- `HUB_SERVICE_PORT` — Port for the hub-service (e.g. `8081`)
- `CLIENT_SERVICE_PORT` — Port for the client-service (e.g. `8082`)

### JWT Secrets

- `ACCESS_SECRET` — Secret key used to sign access tokens
- `REFRESH_SECRET` — Secret key used to sign refresh tokens


## Running the Project

Make sure you have Docker installed.  
- Clone the repository:

```bash
git clone https://github.com/sokolmax415/smart-building-backend.git
cd smart-building-backend
```
- Build and start the services:
```bash
docker-compose up --build
```

- Stop services
```bash 
docker-compose down
```

**Note:** On the first run, please wait a few minutes for the PostgreSQL database container to fully initialize before the other services start working properly.

## Author

Developed by **Max Sokolov**  
Originally built as a university coursework project.  
Email: sokolmax415@gmail.com  
GitHub: [sokolmax415](https://github.com/sokolmax415)
