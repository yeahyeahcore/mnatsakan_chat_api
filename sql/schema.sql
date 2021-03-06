DROP DATABASE IF EXISTS mnatsakan_chat_api_db;
CREATE DATABASE mnatsakan_chat_api_db;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR NOT NULL UNIQUE,
  username VARCHAR NOT NULL,
  login    VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL
)