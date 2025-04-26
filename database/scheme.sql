
CREATE SCHEMA IF NOT EXISTS userDB;


CREATE TABLE IF NOT EXISTS userDB.roles(
    role_id BIGSERIAL PRIMARY KEY,
    role TEXT NOT NULL,
    permissions TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS userDB.users(
    id BIGSERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role_id BIGINT REFERENCES userDB.roles (role_id) ON DELETE SET NULL,
    registration_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP);



INSERT INTO userDB.roles (role,permissions)
VALUES ('admin', 'all_permissions'),
('user', 'limited_permissions'),
('hub', 'limited_permissions');

INSERT INTO userDB.users (firstname, lastname, login, password_hash, role_id)
VALUES ('Hub', '001', 'hub001', '$2a$10$pNshLwyaU.1c8j.AwOEAzem6C0O.dI3TsSOjTlGnKvw9ClzqF.Uoi', 
(SELECT role_id FROM userDB.roles WHERE role = 'hub'));

INSERT INTO userDB.users (firstname, lastname, login, password_hash, role_id)
VALUES ('ADMIN', 'ADMIN', 'admin', '$2a$10$khcpwC6Y5Vt1g6OogKSGs.H2.8/h22bOsyxnqYLvWPN2aX0X5DFve', 
(SELECT role_id FROM userDB.roles WHERE role = 'admin'));


CREATE SCHEMA IF NOT EXISTS building;

CREATE TABLE IF NOT EXISTS building.location(
    location_id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT REFERENCES building.location(location_id) ON DELETE SET NULL,
    user_id BIGINT REFERENCES userDB.users(id) ON DELETE CASCADE,
    location_type TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS building.hubs(
    hub_id BIGSERIAL PRIMARY KEY,
    sn TEXT UNIQUE NOT NULL,
    location_id BIGINT REFERENCES building.location(location_id) ON DELETE CASCADE,
    device_count BIGINT,
    uptime BIGINT,
    last_ping_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fw_version TEXT
);

