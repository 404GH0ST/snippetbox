CREATE TABLE snippets (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_expires ON snippets (created);

CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (name, email, hashed_password, created) VALUES (
  'Alice Jones',
  'alice@example.com',
  '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HewTy.eS206TNfkGfr6HzGJSWG',
  '2022-01-01 09:18:24'
);