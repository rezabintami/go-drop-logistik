CREATE TABLE IF NOT EXISTS manifests (
  id bigserial PRIMARY KEY,
  code varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  driver_id integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);