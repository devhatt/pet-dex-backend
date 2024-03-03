ALTER TABLE users DROP COLUMN deleted_at;

ALTER TABLE users DROP COLUMN updated_at;

ALTER TABLE users DROP COLUMN created_at;

ALTER TABLE users DROP COLUMN pass;

ALTER TABLE users CHANGE COLUMN document documentId VARCHAR(14); 



