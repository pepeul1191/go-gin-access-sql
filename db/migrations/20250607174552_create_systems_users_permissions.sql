-- migrate:up

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

-- migrate:down

DROP TABLE system_user_permissions;