-- Database: pinterest

-- DROP DATABASE pinterest;

CREATE DATABASE pinterest
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- DROP TABLE IF EXISTS users CASCADE;
-- DROP TABLE IF EXISTS pins CASCADE;
-- DROP TABLE IF EXISTS boards CASCADE;
-- DROP TABLE IF EXISTS subscriptions CASCADE;
-- DROP TABLE IF EXISTS commentaries CASCADE;

CREATE TABLE users(
	id serial PRIMARY KEY,
	email text NOT NULL,
	login text UNIQUE NOT NULL,
	encrypted_password text NOT NULL,
	about text,
	avatar text,
	subscriptions int,
	subscribers int,
	created_at TIMESTAMP NOT NULL
)

CREATE TABLE boards(
	id serial PRIMARY KEY,
	user_id int REFERENCES users(id),
	name text NOT NULL,
	description text,
	created_at timestamp
)

CREATE TABLE pins (
	id serial PRIMARY KEY,
	user_id int REFERENCES users(id),
	name text,
	description text,
	image text NOT NULL,
	board_id int REFERENCES boards(id),
	created_at timestamp
)

CREATE TABLE subscriptions (
	id serial PRIMARY KEY,
	user_id int REFERENCES users(id) NOT NULL,
	subscribed_at int REFERENCES users(id) NOT NULL,
	UNIQUE (user_id, subscribed_at)
)

CREATE TABLE commentaries (
	id serial PRIMARY KEY,
	user_id int REFERENCES users(id) NOT NULL,
	pin_id int REFERENCES pins(id) NOT NULL,
	comment text NOT NULL,
	created_at timestamp
)