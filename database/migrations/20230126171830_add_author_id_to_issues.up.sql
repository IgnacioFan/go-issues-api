ALTER TABLE issues
ADD COLUMN author_id int;

ALTER TABLE issues
ADD CONSTRAINT fk_author_id
FOREIGN KEY (author_id) REFERENCES users(id);
