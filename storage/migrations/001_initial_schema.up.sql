BEGIN;

CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  username varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  birth_date date,
  bio text,
  hometown varchar(255)
);

CREATE TABLE posts (
  id SERIAL PRIMARY KEY NOT NULL,
  title varchar(255) NOT NULL UNIQUE,
  content text NOT NULL,
  date date NOT NULL,
  user_id INTEGER NOT NULL,
  CONSTRAINT posts_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE followers(
  follower_id INTEGER NOT NULL,
  target_id INTEGER NOT NULL,
  CONSTRAINT followers_pkey PRIMARY KEY (follower_id, target_id),
  CONSTRAINT followers_users_follower_fkey FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT followers_users_target_fkey FOREIGN KEY (target_id) REFERENCES users(id) ON DELETE CASCADE
);

COMMIT;