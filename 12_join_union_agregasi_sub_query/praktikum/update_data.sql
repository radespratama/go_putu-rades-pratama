-- 1
UPDATE `products`
SET `name` = 'product dummy'
WHERE `id` = 1;

-- 2
UPDATE `transaction_details`
SET `quantity` = 3
WHERE `product_id` = 1;