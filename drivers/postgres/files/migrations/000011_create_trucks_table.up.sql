CREATE TABLE IF NOT EXISTS trucks (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  type varchar(255) NOT NULL,
  license_plate varchar(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL
);