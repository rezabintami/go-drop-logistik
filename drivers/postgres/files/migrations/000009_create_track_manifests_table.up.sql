CREATE TABLE IF NOT EXISTS track_manifests (
  id bigserial PRIMARY KEY, 
  track_id integer NOT NULL,
  manifest_id integer NOT NULL
);