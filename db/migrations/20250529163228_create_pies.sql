-- migrate:up

CREATE TABLE pies (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  nombre VARCHAR(10) NOT NULL
);

-- migrate:down 

DROP TABLE pies;