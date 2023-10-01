CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS urls (
    id uuid DEFAULT uuid_generate_v4 (),
    short varchar(100) NOT NULL,
    full varchar(100) NOT NULL
);

