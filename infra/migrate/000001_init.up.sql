CREATE TABLE IF NOT EXISTS users (
	user_id  SERIAL PRIMARY KEY,
	username VARCHAR(64)         UNIQUE NOT NULL,
	password VARCHAR(64)                NOT NULL,
	pow      VARCHAR(32)                NOT NULL,
	email    VARCHAR(256)        UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens (
	user_id       INTEGER            NOT NULL REFERENCES users(user_id),
	device_id     VARCHAR(64)        UNIQUE NOT NULL,
	refresh_token VARCHAR(36)        UNIQUE NOT NULL,
	token_date    BIGINT             NOT NULL
);
