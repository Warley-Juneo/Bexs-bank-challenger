CREATE TABLE consumer (
    id          SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name        VARCHAR(50) NOT NULL,
    national_id VARCHAR(20) NOT NULL UNIQUE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE payment (
    id          SERIAL PRIMARY KEY NOT NULL UNIQUE,
    partner_id  INTEGER NOT NULL REFERENCES partner(id),
    amount      NUMERIC NOT NULL,
    consumer_id INTEGER NOT NULL REFERENCES consumer(id),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);