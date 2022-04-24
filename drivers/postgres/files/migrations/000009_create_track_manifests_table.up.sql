CREATE TABLE IF NOT EXISTS track_manifests (
  track_id integer UNIQUE NOT NULL,
  manifest_id integer UNIQUE NOT NULL
);