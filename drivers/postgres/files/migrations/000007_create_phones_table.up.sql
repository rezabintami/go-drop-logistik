CREATE TABLE IF NOT EXISTS phones (
  id bigserial PRIMARY KEY,
  phone varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);