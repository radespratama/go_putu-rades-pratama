-- 1
SELECT `id`, `name`, `gender`
FROM `users`
WHERE `gender` = 'M';

-- 2
SELECT *
FROM `products`
WHERE `id` = 3;

-- 3
SELECT `id`, `name`, `created_at`
FROM `users`
WHERE `created_at` BETWEEN NOW() - INTERVAL 7 DAY AND NOW()
  AND `name` LIKE '%a%';

-- 4
SELECT COUNT(*) AS `total_female_users`
FROM `users`
WHERE `gender` = 'F';

-- 5
SELECT `id`, `name`, `gender`
FROM `users`
ORDER BY `name` ASC;

-- 6
SELECT *
FROM `products`
LIMIT 5;
