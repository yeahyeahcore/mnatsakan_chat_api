DROP DATABASE IF EXISTS social_network_api_db;
CREATE DATABASE social_network_api_db;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR NOT NULL UNIQUE,
  login VARCHAR NOT NULL UNIQUE,
  username VARCHAR NOT NULL,
)