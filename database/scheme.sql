CREATE SCHEMA IF NOT EXISTS userDB;
CREATE SCHEMA IF NOT EXISTS smartbuildingDB;

CREATE TABLE IF NOT EXISTS userDB.roles(
    role_id BIGSERIAL PRIMARY KEY,
    role TEXT NOT NULL,
    permissions TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS userDB.users(
    id BIGSERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL, -- Алгоритм bcrypt
    role_id BIGINT REFERENCES userDB.roles (role_id) ON DELETE RESTRICT,
    registration_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP);



INSERT INTO userDB.roles (role,permissions)
VALUES ('admin', 'all_permissions'),
('user', 'limited_permissions'),
('hub', 'limited_permissions');

INSERT INTO userDB.users (firstname, lastname, login, password_hash, role_id)
VALUES ('Hub', '001', 'hub001', '$2a$10$pNshLwyaU.1c8j.AwOEAzem6C0O.dI3TsSOjTlGnKvw9ClzqF.Uoi', 
(SELECT role_id FROM userDB.roles WHERE role = 'hub'));

INSERT INTO userDB.users (firstname, lastname, login, password_hash, role_id)
VALUES ('ADMIN', 'ADMIN', 'admin', '$2a$10$aPCPCIB67iYo.kTXOKgZouBglNkzgbxOILM18zqSvj5FhitmDBf/C', 
(SELECT role_id FROM userDB.roles WHERE role = 'admin'));

INSERT INTO userDB.users (firstname, lastname, login, password_hash, role_id)
VALUES ('TEST', 'USER', 'testUser', '$2a$10$R2UYCoOE8MEvnG6vqBmRxO/dV5CLJPW8U1iBqcLHHlXMw7Aw6hQoe', 
(SELECT role_id FROM userDB.roles WHERE role = 'user'));


CREATE TABLE IF NOT EXISTS smartbuildingDB.location(
    location_id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT NULL REFERENCES smartbuildingDB.location(location_id) ON DELETE CASCADE,
    location_type TEXT NOT NULL,
    location_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS smartbuildingDB.hubs(
    hub_id BIGSERIAL PRIMARY KEY,
    hub_sn TEXT UNIQUE NOT NULL,
    location_id BIGINT REFERENCES smartbuildingDB.location(location_id) ON DELETE CASCADE,
    uptime BIGINT DEFAULT 0,
    last_ping_time TIMESTAMP,
    fw_version TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS smartbuildingDB.devices(
    device_id BIGSERIAL PRIMARY KEY,
    device_sn TEXT UNIQUE NOT NULL,
    hub_sn TEXT REFERENCES smartbuildingDB.hubs(hub_sn) ON DELETE CASCADE,
    device_type TEXT NOT NULL,
    device_name TEXT NOT NULL,
    last_ping_time TIMESTAMP,
    fw_version TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS smartbuildingDB.metrics (
    metric_id BIGSERIAL PRIMARY KEY,
    device_sn TEXT REFERENCES smartbuildingDB.devices(device_sn) ON DELETE CASCADE,
    data JSONB NOT NULL,
    send_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


--Index for userDB
CREATE INDEX idx_roles_role ON userDB.roles(role);

--Index for smartbuildgDB
CREATE INDEX idx_hubs_location_id ON smartbuildingDB.hubs(location_id);
CREATE INDEX idx_devices_hub_sn ON smartbuildingDB.devices(hub_sn);
CREATE INDEX idx_location_parent_id ON smartbuildingDB.location(parent_id);
CREATE INDEX idx_metrics_device_sn ON smartbuildingDB.metrics(device_sn);
CREATE INDEX idx_metrics_send_time ON smartbuildingDB.metrics(send_time);


INSERT INTO smartbuildingDB.location(parent_id,location_type,location_name) VALUES(NULL,'Floor','bedroom');
INSERT INTO smartbuildingDB.location(parent_id,location_type,location_name) VALUES(NULL,'Floor','kitchen');

INSERT INTO smartbuildingDB.hubs(hub_sn,location_id,uptime,last_ping_time,fw_version,created_at) VALUES ('HubSn1',2,3600,NOW(),'v1',NOW());
INSERT INTO smartbuildingDB.devices(device_sn, hub_sn,device_type,device_name,last_ping_time,fw_version,created_at) VALUES ('DeviceSn1','HubSn1','temperature_sensor','Temp Sensor 1',NOW(),'v1',NOW());

INSERT INTO smartbuildingDB.metrics (device_sn, data, send_time)
VALUES   ('DeviceSn1', '{"temperature": 21.0, "humidity": 55}', '2025-05-17T14:00:00Z'),
  ('DeviceSn1', '{"temperature": 21.5, "humidity": 56}', '2025-05-17T14:15:00Z'),
  ('DeviceSn1', '{"temperature": 22.5, "humidity": 60}', '2025-05-17T14:30:00Z'),
  ('DeviceSn1', '{"temperature": 23.0, "humidity": 62}', '2025-05-17T14:45:00Z'),
  ('DeviceSn1', '{"temperature": 23.5, "humidity": 63}', '2025-05-17T15:00:00Z'),
  ('DeviceSn1', '{"temperature": 24.0, "humidity": 65}', '2025-05-17T15:15:00Z'),
  ('DeviceSn1', '{"temperature": 24.5, "humidity": 66}', '2025-05-17T15:30:00Z');