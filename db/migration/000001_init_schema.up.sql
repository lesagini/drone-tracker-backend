CREATE TYPE "pilot_statuses" AS ENUM (
  'active',
  'inactive',
  'suspended'
);

CREATE TYPE "pilot_classifications" AS ENUM (
  'A',
  'B',
  'C',
  'D'
);

CREATE TYPE "variety_types" AS ENUM (
  'continuous',
  'cyclic'
);

CREATE TABLE "pilots" (
  "id" bigserial NOT NULL,
  "pilot_id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "pilot_operator_id" varchar NOT NULL,
  "pilot_initials" varchar NOT NULL,
  "pilot_number" int NOT NULL,
  "pilot_full_name" varchar NOT NULL,
  "pilot_entry_date" timestamp DEFAULT 'now()' NOT NULL,
  "pilot_license_number" bigint UNIQUE NOT NULL,
  "pilot_farm_location_code" varchar NOT NULL,
  "pilot_farm_location" varchar NOT NULL,
  "pilot_status" pilot_statuses NOT NULL,
  "pilot_classification" pilot_classifications NOT NULL,
  "pilot_flight_hours" numeric NOT NULL,
  "pilot_covered_acreage" numeric NOT NULL
);

CREATE TABLE "farms" (
  "id" bigserial NOT NULL,
  "farm_code" varchar UNIQUE PRIMARY KEY NOT NULL,
  "farm_coordinates" varchar UNIQUE NOT NULL,
  "farm_airspace" varchar NOT NULL,
  "farm_location" varchar NOT NULL,
  "farm_geolocation" varchar NOT NULL,
  "farm_creation_date" timestamp DEFAULT 'now()' NOT NULL,
  "farm_contact" bigint NOT NULL
);

CREATE TABLE "flights" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "flight_date" timestamp DEFAULT 'now()' NOT NULL,
  "flight_farm_location" varchar NOT NULL,
  "flight_farm_id" varchar NOT NULL,
  "flight_farm_geolocation" varchar NOT NULL,
  "flight_duration" numeric NOT NULL,
  "flight_pilot" varchar NOT NULL,
  "flight_acreage" numeric NOT NULL
);

CREATE TABLE "operators" (
  "id" bigserial NOT NULL,
  "operator_id" varchar UNIQUE NOT NULL,
  "operator_name" varchar NOT NULL,
  "operator_headquater" varchar NOT NULL,
  "operator_number_pilots_deployed" int NOT NULL,
  "opertor_contact" bigint NOT NULL,
  PRIMARY KEY ("id", "operator_id")
);

CREATE TABLE "varieties" (
  "id" bigserial PRIMARY KEY,
  "variety_internal_identity" varchar UNIQUE NOT NULL,
  "variety_botanical_name" varchar NOT NULL,
  "variety_farm_id" varchar NOT NULL,
  "variety_creation_date" timestamp DEFAULT 'now()' NOT NULL,
  "variety_acreage" bigint NOT NULL,
  "variety_type" variety_types NOT NULL,
  "variety_iterval_code" varchar NOT NULL
);

COMMENT ON COLUMN "varieties"."variety_iterval_code" IS '7 day encoding e.g 1010000 for sunday and tuesday schedule';

ALTER TABLE "pilots" ADD FOREIGN KEY ("pilot_operator_id") REFERENCES "operators" ("operator_id");

ALTER TABLE "pilots" ADD FOREIGN KEY ("pilot_farm_location_code") REFERENCES "farms" ("farm_code");

ALTER TABLE "flights" ADD FOREIGN KEY ("flight_farm_id") REFERENCES "farms" ("farm_code");

ALTER TABLE "flights" ADD FOREIGN KEY ("flight_pilot") REFERENCES "pilots" ("pilot_id");

ALTER TABLE "varieties" ADD FOREIGN KEY ("variety_farm_id") REFERENCES "farms" ("farm_code");

ALTER SEQUENCE farms_id_seq RESTART WITH 1;

ALTER SEQUENCE operators_id_seq RESTART WITH 1;

ALTER SEQUENCE pilots_id_seq RESTART WITH 1;

ALTER SEQUENCE flights_id_seq RESTART WITH 1;

ALTER SEQUENCE varieties_id_seq RESTART WITH 1;