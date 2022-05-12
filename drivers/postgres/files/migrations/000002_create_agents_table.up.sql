CREATE TABLE IF NOT EXISTS agents (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  roles varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  address TEXT NOT NULL,
  balance INTEGER NOT NULL,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL
);