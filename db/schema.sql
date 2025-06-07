CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username VARCHAR(20) NOT NULL,
  password VARCHAR(100) NOT NULL,
  activation_key VARCHAR(30),
  reset_key VARCHAR(30),
  email VARCHAR(50) UNIQUE NOT NULL,
  activated BOOLEAN NOT NULL DEFAULT 0,
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL
);
CREATE TABLE systems (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(40) NOT NULL,
  description TEXT,
  repository VARCHAR(100),
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL
);
CREATE TABLE roles (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(40) NOT NULL,
  description TEXT,
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL,
  system_id INTEGER NOT NULL,
  FOREIGN KEY(system_id) REFERENCES systems(id)
);
CREATE TABLE permissions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(40) NOT NULL,
  description TEXT,
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL,
  role_id INTEGER NOT NULL,
  FOREIGN KEY(role_id) REFERENCES roles(id)
);
CREATE TABLE system_users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  system_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  created DATETIME NOT NULL,
  FOREIGN KEY(system_id) REFERENCES systems(id),
  FOREIGN KEY(user_id) REFERENCES users(id)
);
CREATE TABLE system_user_permissions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  system_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  permission_id INTEGER NOT NULL,
  created DATETIME NOT NULL,
  FOREIGN KEY(system_id) REFERENCES systems(id),
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(permission_id) REFERENCES permissions(id)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20250607174507'),
  ('20250607174516'),
  ('20250607174522'),
  ('20250607174528'),
  ('20250607174542'),
  ('20250607174552'),
  ('20250607175027'),
  ('20250607175057');
