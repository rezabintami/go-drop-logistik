CREATE TABLE IF NOT EXISTS tracks (
  id bigserial PRIMARY KEY,
  start_agent_id integer NOT NULL,
  current_agent_id integer NOT NULL,
  destination_agent_id integer NOT NULL,
  message text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NULL,
  deleted_at timestamptz NULL
);