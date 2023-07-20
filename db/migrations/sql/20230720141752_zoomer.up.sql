CREATE TABLE IF NOT EXISTS users(
   id TEXT PRIMARY KEY,
   username VARCHAR (16) UNIQUE NOT NULL,
   password VARCHAR (20) NOT NULL,
   limit bigint NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS rooms(
   id text PRIMARY KEY,
   name text,
   description text,
   category text,
   created_at timestamptz NOT NULL DEFAULT (now()),
   updated_at timestamptz NOT NULL DEFAULT (now()),
   created_by text REFERENCES users (id)
);
