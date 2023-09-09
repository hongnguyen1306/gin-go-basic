-- +migrate Up
ALTER TABLE users
  ADD CONSTRAINT users_pk
    PRIMARY KEY (id);

ALTER TABLE users ALTER COLUMN id SET DATA TYPE uuid USING uuid_generate_v4 ();


ALTER TABLE users 
    ADD CONSTRAINT uniq UNIQUE (email);

-- +migrate Down
ALTER TABLE users
  DROP CONSTRAINT uniq;

ALTER TABLE users
ALTER COLUMN id SET DATA TYPE text;

ALTER TABLE users
  DROP CONSTRAINT users_pk;