CREATE TABLE IF NOT EXISTS users (
  id bigserial NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  roles varchar(255) NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL
);