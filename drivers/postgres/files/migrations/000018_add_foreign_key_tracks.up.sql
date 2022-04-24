ALTER Table
  tracks
ADD
  CONSTRAINT fk_tracks_current_agent FOREIGN KEY(current_agent_id) REFERENCES agents(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER Table
  tracks
ADD
  CONSTRAINT fk_tracks_destination_agent FOREIGN KEY(destination_agent_id) REFERENCES agents(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER Table
  tracks
ADD
  CONSTRAINT fk_tracks_start_agent FOREIGN KEY(start_agent_id) REFERENCES agents(id) ON DELETE CASCADE ON UPDATE CASCADE;