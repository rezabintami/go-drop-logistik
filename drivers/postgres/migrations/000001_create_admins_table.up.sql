CREATE TABLE IF NOT EXISTS admins (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  roles varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL
);