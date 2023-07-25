-- Remove the unique constraint from the 'field_name' and 'field_farm_id' columns
ALTER TABLE fields DROP CONSTRAINT unique_field_farm_names;