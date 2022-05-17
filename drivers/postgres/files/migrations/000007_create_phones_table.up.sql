CREATE TABLE IF NOT EXISTS phones (
  id bigserial PRIMARY KEY,
  phone varchar(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL
);