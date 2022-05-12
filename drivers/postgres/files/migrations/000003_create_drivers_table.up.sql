CREATE TABLE IF NOT EXISTS drivers (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  address text NOT NULL,
  truck_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL
);