CREATE TABLE IF NOT EXISTS trucks (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  type varchar(255) NOT NULL,
  license_plate varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);