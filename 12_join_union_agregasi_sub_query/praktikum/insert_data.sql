-- 1.a
INSERT INTO operators (`name`) VALUES
  ('Operator A'),
  ('Operator B'),
  ('Operator C'),
  ('Operator D'),
  ('Operator E');

-- 1.b
INSERT INTO product_types (`name`) VALUES
  ('Type 1'),
  ('Type 2'),
  ('Type 3');

-- 1.c
INSERT INTO products (`product_type_id`, `operator_id`, `code`, `name`, `status`) VALUES
  (1, 3, 'P1', 'Product 1', 1),
  (1, 3, 'P2', 'Product 2', 1);

-- 1.d
INSERT INTO products (`product_type_id`, `operator_id`, `code`, `name`, `status`) VALUES
  (2, 1, 'P3', 'Product 3', 1),
  (2, 1, 'P4', 'Product 4', 1),
  (2, 1, 'P5', 'Product 5', 1);

-- 1.e
INSERT INTO products (`product_type_id`, `operator_id`, `code`, `name`, `status`) VALUES
  (3, 4, 'P6', 'Product 6', 1),
  (3, 4, 'P7', 'Product 7', 1),
  (3, 4, 'P8', 'Product 8', 1);

-- 1.f
INSERT INTO product_descriptions (`product_id`, `description`) VALUES
  (1, 'Description for Product 1'),
  (2, 'Description for Product 2'),
  (3, 'Description for Product 3'),
  (4, 'Description for Product 4'),
  (5, 'Description for Product 5'),
  (6, 'Description for Product 6'),
  (7, 'Description for Product 7'),
  (8, 'Description for Product 8');

-- 1.g
INSERT INTO payment_methods (`name`, `status`) VALUES
  ('Payment Method 1', 1),
  ('Payment Method 2', 1),
  ('Payment Method 3', 1);

-- 1.h
INSERT INTO users (`id`, `status`, `birthdate`, `gender`) VALUES
  (1, 1, '1990-01-01', 'M'),
  (2, 1, '1995-05-15', 'F'),
  (3, 1, '1985-10-30', 'M'),
  (4, 1, '1980-03-20', 'F'),
  (5, 1, '1992-07-07', 'M');

-- 1.i
INSERT INTO transactions (`id`, `user_id`, `payment_method_id`, `status`, `total_quantity`, `total_price`) VALUES
  (1, 1, 1, 'Paid', 2, 100.00),
  (2, 1, 2, 'Unpaid', 3, 150.00),
  (3, 1, 3, 'Paid', 1, 50.00),
  (4, 2, 1, 'Paid', 2, 100.00),
  (5, 2, 2, 'Unpaid', 3, 150.00),
  (6, 2, 3, 'Paid', 1, 50.00),
  (7, 3, 1, 'Paid', 2, 100.00),
  (8, 3, 2, 'Unpaid', 3, 150.00),
  (9, 3, 3, 'Paid', 1, 50.00),
  (10, 4, 1, 'Paid', 2, 100.00),
  (11, 4, 2, 'Unpaid', 3, 150.00),
  (12, 4, 3, 'Paid', 1, 50.00),
  (13, 5, 1, 'Paid', 2, 100.00),
  (14, 5, 2, 'Unpaid', 3, 150.00),
  (15, 5, 3, 'Paid', 1, 50.00);

-- 1.j
INSERT INTO transaction_details (`transaction_id`, `product_id`, `status`, `quantity`, `price`) VALUES
  (1, 1, 'Paid', 1, 50.00),
  (1, 2, 'Paid', 1, 50.00),
  (2, 1, 'Unpaid', 1, 50.00),
  (2, 2, 'Unpaid', 1, 50.00),
  (2, 3, 'Unpaid', 1, 50.00),
  (3, 2, 'Paid', 1, 50.00),
  (4, 3, 'Paid', 1, 50.00),
  (5, 3, 'Unpaid', 1, 50.00),
  (6, 1, 'Paid', 1, 50.00),
  (6, 2, 'Paid', 1, 50.00),
  (6, 3, 'Paid', 1, 50.00),
  (7, 2, 'Paid', 1, 50.00),
  (8, 4, 'Paid', 1, 50.00),
  (9, 4, 'Unpaid', 1, 50.00),
  (10, 1, 'Paid', 1, 50.00),
  (11, 2, 'Paid', 1, 50.00),
  (12, 3, 'Paid', 1, 50.00),
  (13, 2, 'Paid', 1, 50.00),
  (14, 3, 'Paid', 1, 50.00),
  (15, 1, 'Paid', 1, 50.00);

