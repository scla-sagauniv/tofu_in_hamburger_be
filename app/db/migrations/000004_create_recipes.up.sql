CREATE TABLE `recipes` (
  `id` int NOT NULL,
  `recipe_title` varchar(256) NOT NULL,
  `recipe_url` varchar(256) NOT NULL,
  `foodImage_url` varchar(256) NOT NULL,
  `pickup` boolean NOT NULL,
  `nickname` varchar(256) NOT NULL,
  `recipe_materials` varchar(256) NOT NULL,
  `recipe_material_ids` varchar(256) NOT NULL,
  `recipe_publishday` datetime NOT NULL,
  `rank` int NOT NULL,
  `recipe_indication_id` int NOT NULL,
  `recipe_cost_id` int NOT NULL,
  `created_at` datetime NOT NULL NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_1` (`recipe_indication_id`),
  CONSTRAINT `FK_1` FOREIGN KEY `FK_1` (`recipe_indication_id`) REFERENCES `recipe_indications` (`id`),
  KEY `FK_2` (`recipe_cost_id`),
  CONSTRAINT `FK_2` FOREIGN KEY `FK_2` (`recipe_cost_id`) REFERENCES `recipe_costs` (`id`)
);