ALTER TABLE
  manifest_receipts
ADD
  CONSTRAINT fk_manifest_receipts_manifest FOREIGN KEY(manifest_id) REFERENCES manifests(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE
  manifest_receipts
ADD
  CONSTRAINT fk_manifest_receipts_receipt FOREIGN KEY(receipt_id) REFERENCES receipts(id) ON DELETE CASCADE ON UPDATE CASCADE;