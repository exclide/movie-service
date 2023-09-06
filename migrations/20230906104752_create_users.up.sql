CREATE TABLE IF NOT EXISTS users
(
	login text NOT NULL,
	password text NOT NULL,
	PRIMARY KEY (login)
);