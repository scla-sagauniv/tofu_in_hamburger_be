ALTER TABLE recipes
  rename column recipe_title to title;
ALTER TABLE recipes
  rename column foodImage_url to image_url;
ALTER TABLE recipes
  rename column recipe_materials to materials;
ALTER TABLE recipes
  rename column recipe_material_ids to material_ids;
ALTER TABLE recipes
  rename column recipe_publishday to publishday;