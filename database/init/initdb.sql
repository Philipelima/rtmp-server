DROP TABLE IF EXISTS keys;

CREATE TABLE keys (
    id SERIAL PRIMARY KEY,
    streaming_key VARCHAR NOT NULL
);


INSERT INTO keys(id, streaming_key) VALUES(1, '3c2e8dd7-a54f-4442-8fb2-3f4f10bd21db'),
