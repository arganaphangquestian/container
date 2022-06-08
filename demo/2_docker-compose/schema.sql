CREATE TABLE IF NOT EXISTS todos (
  id   BIGSERIAL PRIMARY KEY,
  title text    NOT NULL,
  done  boolean  NOT NULL default false,
  created_at timestamp NOT NULL default now()
);