# Notes

## Dependencies
```bash
go get github.com/go-sql-driver/mysql@v1
go get github.com/justinas/alice@v1
go get github.com/go-playground/form/v4@v4
go get github.com/alexedwards/scs/v2@v2
go get github.com/alexedwards/scs/mysqlstore@latest
go get golang.org/x/crypto/bcrypt@latest
go get github.com/justinas/nosurf@v1
```

## Creating Database
```sql
-- Create a new UTF-8 "snippetbox" database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

## Creating Snippets Tables
```sql
-- Create a "snippets" table.
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);
```

## Inserting some dummy data
```sql
-- Add some dummy records
INSERT INTO snippets (title, content, created, expires) VALUES (
        "An old silent pond",
        "An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Basho",
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
        "Over the wintry forest",
        "Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n- Natsume Soseki",
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
        "Mount Fuji",
        "Mount Fuji is the tallest\nmountain in Japan,\nstanding at 3,776 meters (12,380 feet).\n\n- Japanese",
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
```

## Create a user
```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'summer2024';
```

## Creating sessions Table
```sql
CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);
```

## Generate a trusted self-signed certificate
```bash
mkdir tls && cd tls
mkcert -install
mkcert -cert-file cert.pem -key-file key.pem localhost
```

## Create a users Table
```sql
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```
