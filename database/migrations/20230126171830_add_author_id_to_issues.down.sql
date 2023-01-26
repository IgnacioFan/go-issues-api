ALTER TABLE issues
DROP CONSTRAINT fk_author_id;

ALTER TABLE issues
DROP COLUMN author_id;
