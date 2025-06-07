-- migrate:up

INSERT INTO foots (id,name) VALUES (1, 'Derecho');
INSERT INTO foots (id,name) VALUES (2, 'Izquierdo');

-- migrate:down

DELETE FROM foots;
