CREATE TABLE partner (
  id              SERIAL PRIMARY KEY NOT NULL UNIQUE,
  trading_name    VARCHAR(50) NOT NULL,
  document        VARCHAR(18) NOT NULL UNIQUE,
  currency        VARCHAR(3) NOT NULL
);
