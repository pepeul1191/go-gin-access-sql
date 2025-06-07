-- migrate:up

INSERT INTO users (username, password, activation_key, reset_key, email, activated, created, updated)
VALUES (
  'ana', '123', NULL, NULL, 'ana@example.com', 1, '2025-06-06 10:00:00', '2025-06-06 10:00:00'
);
INSERT INTO users (username, password, activation_key, reset_key, email, activated, created, updated)
VALUES (
  'luis', '123', NULL, NULL, 'luis@example.com', 1, '2025-06-06 10:05:00', '2025-06-06 10:05:00'
);
INSERT INTO users (username, password, activation_key, reset_key, email, activated, created, updated)
VALUES (
  'marta', '123', NULL, NULL, 'marta@example.com', 0, '2025-06-06 10:10:00', '2025-06-06 10:10:00'
);
INSERT INTO users (username, password, activation_key, reset_key, email, activated, created, updated)
VALUES (
  'pedro', '123', NULL, NULL, 'pedro@example.com', 1, '2025-06-06 10:20:00', '2025-06-06 10:20:00'
);
INSERT INTO users (username, password, activation_key, reset_key, email, activated, created, updated)
VALUES (
  'lucia', '123', NULL, NULL, 'lucia@example.com', 0, '2025-06-06 10:30:00', '2025-06-06 10:30:00'
);

-- migrate:down

DELETE FROM users;