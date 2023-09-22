-- 1
SELECT *
FROM `transactions`
WHERE `user_id` IN (1, 2);

-- 2
SELECT SUM(`total_price`) AS `total_price`
FROM `transactions`
WHERE `user_id` = 1;

-- 3
SELECT COUNT(*) AS `total_transactions`
FROM `transactions`
WHERE `id` IN (
  SELECT `transaction_id`
  FROM `transaction_details`
  WHERE `product_id` IN (
    SELECT `id`
    FROM `products`
    WHERE `product_type_id` = 2
  )
);

-- 4
SELECT `products`.*, `product_types`.`name` AS `product_type_name`
FROM `products`
JOIN `product_types` ON `products`.`product_type_id` = `product_types`.`id`;

-- 5
SELECT `transactions`.*, `products`.`name` AS `product_name`, `users`.`name` AS `user_name`
FROM `transactions`
JOIN `products` ON `transactions`.`product_id` = `products`.`id`
JOIN `users` ON `transactions`.`user_id` = `users`.`id`;

-- 6
DELIMITER //

CREATE FUNCTION delete_transaction_and_details(transaction_id INT)
RETURNS BOOLEAN
BEGIN
  DELETE FROM `transaction_details` WHERE `transaction_id` = transaction_id;
  DELETE FROM `transactions` WHERE `id` = transaction_id;
  RETURN TRUE;
END //

DELIMITER ;

-- 7
DELIMITER //

CREATE FUNCTION update_total_qty(transaction_id INT)
RETURNS BOOLEAN
BEGIN
  DECLARE total_quantity INT;
  SELECT SUM(`quantity`) INTO total_quantity
  FROM `transaction_details`
  WHERE `transaction_id` = transaction_id;
  UPDATE `transactions` SET `total_quantity` = total_quantity WHERE `id` = transaction_id;
  RETURN TRUE;
END //

DELIMITER ;

-- 8
SELECT *
FROM `products`
WHERE `id` NOT IN (
  SELECT DISTINCT `product_id`
  FROM `transaction_details`
);
