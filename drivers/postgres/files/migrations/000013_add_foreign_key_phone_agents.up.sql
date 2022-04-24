ALTER TABLE
  phone_agents
ADD
  CONSTRAINT fk_phone_agents_agent FOREIGN KEY(agent_id) REFERENCES agents(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE
  phone_agents
ADD
  CONSTRAINT fk_phone_agents_phone FOREIGN KEY(phone_id) REFERENCES phones(id) ON DELETE CASCADE ON UPDATE CASCADE;