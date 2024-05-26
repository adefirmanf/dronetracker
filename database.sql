-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- https://wiki.postgresql.org/wiki/Aggregate_Median#median.28numeric.29
CREATE OR REPLACE FUNCTION _final_median(numeric[])
   RETURNS numeric AS
$$
   SELECT AVG(val)
   FROM (
     SELECT val
     FROM unnest($1) val
     ORDER BY 1
     LIMIT  2 - MOD(array_upper($1, 1), 2)
     OFFSET CEIL(array_upper($1, 1) / 2.0) - 1
   ) sub;
$$
LANGUAGE 'sql' IMMUTABLE;

CREATE AGGREGATE median(numeric) (
  SFUNC=array_append,
  STYPE=numeric[],
  FINALFUNC=_final_median,
  INITCOND='{}'
);

CREATE TABLE estate (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	width int4 NOT NULL DEFAULT 0,
	length int4 NOT NULL DEFAULT 0,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	-- Add constraint for ID
	CONSTRAINT estate_id PRIMARY KEY (id)
);

CREATE TABLE tree (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	estate_id uuid NOT NULL,
	x_coordinate int NULL NULL DEFAULT 0,
	y_coordinate int NOT NULL DEFAULT 0,
	height int4 NOT NULL DEFAULT 0,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	-- Add constraint for ID
	CONSTRAINT tree_id PRIMARY KEY (id),
	CONSTRAINT fk_estate FOREIGN KEY (estate_id) REFERENCES estate(id)
);