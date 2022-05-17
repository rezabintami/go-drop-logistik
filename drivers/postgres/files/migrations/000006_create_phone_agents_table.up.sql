CREATE TABLE IF NOT EXISTS phone_agents (
  id bigserial PRIMARY KEY,
  phone_id integer NOT NULL,
  agent_id integer NOT NULL
);