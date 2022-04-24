ALTER Table
  manifests
ADD
  CONSTRAINT fk_manifests_driver FOREIGN KEY(driver_id) REFERENCES drivers(id) ON DELETE CASCADE ON UPDATE CASCADE;