CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE customers (
  id VARCHAR(50) PRIMARY KEY NOT NULL,
  NAME VARCHAR(255) NOT NULL,
  address TEXT NOT NULL,
  date_of_birth DATE NOT NULL,
  status_user ENUM('active', 'suspend'),
  gender ENUM('Laki-Laki', 'Perempuan'),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
  id VARCHAR(50) PRIMARY KEY NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  product_type VARCHAR(255) NOT NULL,
  product_description VARCHAR(255) NOT NULL,
  operator VARCHAR(255) NOT NULL,
  payment_method VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  customer_id VARCHAR(50) NOT NULL,
  total_price DOUBLE NOT NULL,
  STATUS ENUM('Paid', 'Unpaid') NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE detail_transactions (
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  transaction_id VARCHAR(50) NOT NULL,
  product_id VARCHAR(50) NOT NULL,
  quantity INT NOT NULL,
  price DOUBLE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE kurir (
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  NAME VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE kurir ADD COLUMN ongkos_dasar DOUBLE NOT NULL BEFORE created_at;

ALTER TABLE kurir RENAME TO shipping;

DROP TABLE shipping
