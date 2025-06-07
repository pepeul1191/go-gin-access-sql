-- migrate:up

CREATE TABLE foots (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(10) NOT NULL
);

-- migrate:down 

DROP TABLE foots;