ALTER Table
  track_manifests
ADD
  CONSTRAINT fk_track_manifests_manifest FOREIGN KEY(manifest_id) REFERENCES manifests(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER Table
  track_manifests
ADD
  CONSTRAINT fk_track_manifests_track FOREIGN KEY(track_id) REFERENCES tracks(id) ON DELETE CASCADE ON UPDATE CASCADE;