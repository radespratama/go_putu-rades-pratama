CREATE DATABASE 'praktikum'

USE 'praktikum'

CREATE TABLE operators (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_types (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `product_type_id` INT NOT NULL,
  `operator_id` INT NOT NULL,
  `code` VARCHAR(50),
  `name` VARCHAR(100),
  `status` SMALLINT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE products
  ADD CONSTRAINT `product_with_product_type` FOREIGN KEY (`product_type_id`) REFERENCES product_types(`id`),
  ADD CONSTRAINT `product_with_operator` FOREIGN KEY (`operator_id`) REFERENCES operators(`id`);

CREATE TABLE product_descriptions (
  `product_id` INT PRIMARY KEY NOT NULL,
  `description` TEXT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE product_descriptions
  ADD CONSTRAINT `product_detail_meets_product` FOREIGN KEY (`product_id`) REFERENCES products(`id`);

CREATE TABLE payment_methods (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `status` SMALLINT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users (
  `id` INT NOT NULL PRIMARY KEY,
  `status` SMALLINT,
  `birthdate` DATE,
  `gender` CHAR(1),
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  `id` INT NOT NULL PRIMARY KEY,
  `user_id` INT NOT NULL,
  `payment_method_id` INT NOT NULL,
  `status` VARCHAR(10),
  `total_quantity` INT NOT NULL,
  `total_price` INT NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE transactions 
  ADD CONSTRAINT `transaction_with_user` FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  ADD CONSTRAINT `transaction_with_payment_method` FOREIGN KEY (`payment_method_id`) REFERENCES payment_methods(`id`);

CREATE TABLE transaction_details (
  `transaction_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `status` VARCHAR(10),
  `quantity` INT NOT NULL,
  `price` FLOAT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE transaction_details 
  ADD CONSTRAINT `transaction_detail_primary_key` PRIMARY KEY (`transaction_id`, `product_id`),
  ADD CONSTRAINT `transaction_detail_with_transaction` FOREIGN KEY (`transaction_id`) REFERENCES transactions(`id`),
  ADD CONSTRAINT `transaction_detail_with_product` FOREIGN KEY (`product_id`) REFERENCES products(`id`);