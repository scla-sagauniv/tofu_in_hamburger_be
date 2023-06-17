ALTER TABLE recipes
  rename column title to recipe_title;
ALTER TABLE recipes
  rename image_url to column foodImage_url;
ALTER TABLE recipes
  rename column materials to recipe_materials;
ALTER TABLE recipes
  rename column material_ids to recipe_material_ids;
ALTER TABLE recipes
  rename column publishday to recipe_publishday;