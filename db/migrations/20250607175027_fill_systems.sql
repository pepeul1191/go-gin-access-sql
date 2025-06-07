-- migrate:up

INSERT INTO systems (name, description, repository, created, updated)
VALUES (
  'Sistema de Usuarios',
  'Gestión de usuarios, roles y permisos',
  'https://github.com/ejemplo/sistema-usuarios',
  '2025-06-06 12:00:00',
  '2025-06-06 12:00:00'
);
INSERT INTO systems (name, description, repository, created, updated)
VALUES (
  'Inventario',
  'Sistema de control de inventario y stock',
  'https://gitlab.com/empresa/inventario',
  '2025-06-06 13:15:00',
  '2025-06-06 13:15:00'
);

-- migrate:down

DROP TABLE systems;