CREATE TABLE IF NOT EXISTS manifests (
  id bigserial PRIMARY KEY,
  code varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  driver_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL
);