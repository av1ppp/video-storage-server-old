CREATE DATABASE video_storage_server;
\c video_storage_server;

CREATE TABLE archives (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    path VARCHAR NOT NULL
);

CREATE TABLE videos (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    archive_id INTEGER REFERENCES archives(id)
);