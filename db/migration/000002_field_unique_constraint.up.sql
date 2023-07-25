-- Add a unique constraint to the 'field_name' and 'field_farm_id' columns
ALTER TABLE fields ADD CONSTRAINT unique_field_farm_names UNIQUE (field_name, field_farm_id);