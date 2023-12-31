INSERT INTO `recipe_indications` (`id`, `value`)
VALUES (0, '指定なし'),
  (1, '5分以内'),
  (2, '約10分'),
  (3, '約15分'),
  (4, '約30分'),
  (5, '約1時間'),
  (6, '1時間以上');
INSERT INTO `recipe_costs` (`id`, `value`)
VALUES (0, '指定なし'),
  (1, '100円以下'),
  (2, '300円前後'),
  (3, '500円前後'),
  (4, '1,000円前後'),
  (5, '2,000円前後'),
  (6, '3,000円前後'),
  (7, '5,000円前後'),
  (8, '10,000円以上');
INSERT INTO materials (title, description, image_url)
VALUES (
    'Ingredient 1',
    'Description 1',
    'image1.jpg'
  ),
  (
    'Ingredient 2',
    'Description 2',
    'image2.jpg'
  ),
  (
    'Ingredient 3',
    'Description 3',
    'image3.jpg'
  ),
  (
    'Ingredient 4',
    'Description 4',
    'image4.jpg'
  ),
  (
    'Ingredient 5',
    'Description 5',
    'image5.jpg'
  );