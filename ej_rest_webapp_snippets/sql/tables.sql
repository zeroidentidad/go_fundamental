-- Create a `snippets` table.
CREATE TABLE snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);

-- Index on created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Dummy records.
INSERT INTO snippets (title, content, created, expires) VALUES (
'Un viejo estanque silencioso',
'Un viejo estanque silencioso... Una rana salta al estanque,splash! Silencio de nuevo.',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

-- Dummy records.
INSERT INTO snippets (title, content, created, expires) VALUES (
'Sobre el bosque invernal',
'Sobre el bosque invernal, los vientos aúllan de rabia, sin hojas para soplar.',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

-- Dummy records.
INSERT INTO snippets (title, content, created, expires) VALUES (
'Primera mañana de otoño',
'Primera mañana de otoño, el espejo en el que me miro muestra el rostro de mi padre...',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

-- Create a `users` table.
CREATE TABLE users (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
hashed_password CHAR(60) NOT NULL,
created DATETIME NOT NULL
);

-- Constraint on email column.
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);