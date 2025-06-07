-- migrate:up

INSERT INTO pies (id,nombre) VALUES (1, 'Derecho');
INSERT INTO pies (id,nombre) VALUES (2, 'Izquierdo');

-- migrate:down

DELETE FROM pies;
