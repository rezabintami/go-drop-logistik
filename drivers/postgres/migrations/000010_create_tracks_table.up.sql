CREATE TABLE IF NOT EXISTS tracks (
  id bigserial PRIMARY KEY,
  start_agent_id integer NOT NULL,
  current_agent_id integer NOT NULL,
  destination_agent_id integer NOT NULL,
  message text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);