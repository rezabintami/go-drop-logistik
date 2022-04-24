CREATE Table drivers (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  address text NOT NULL,
  truck_id integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL
);