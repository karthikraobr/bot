CREATE TABLE IF NOT EXISTS Customers (
  id SERIAL PRIMARY KEY,  -- Unique identifier for each customer (auto-increment)
  first_name VARCHAR(50) NOT NULL,  -- Customer's first name
  last_name VARCHAR(50) NOT NULL,  -- Customer's last name
  email VARCHAR(100) UNIQUE,  -- Customer's email address (unique)
  phone_number VARCHAR(20),  -- Customer's phone number
  created_at       timestamp without time zone,
  last_modified_at timestamp without time zone,
  deleted_at       timestamp without time zone
);

CREATE TABLE IF NOT EXISTS Products (
  id SERIAL PRIMARY KEY,  -- Unique identifier for each product (auto-increment)
  name VARCHAR(255) NOT NULL,  -- Name of the product
  description TEXT,  -- Detailed description of the product
  price DECIMAL(10,2) NOT NULL,  -- Price of the product
  category VARCHAR(50),  -- Category the product belongs to
  brand VARCHAR(100),  -- Brand of the product
  stock INT NOT NULL DEFAULT 0,  -- Number of units currently in stock (defaults to 0)
  image_url VARCHAR(255),  -- URL of the product image
  created_at       timestamp without time zone,
  last_modified_at timestamp without time zone,
  deleted_at       timestamp without time zone
);

CREATE TABLE IF NOT EXISTS Reviews (
  id SERIAL PRIMARY KEY,  -- Unique identifier for each review (auto-increment)
  product_id INT REFERENCES Products(id) ON DELETE SET NULL,  -- Foreign key referencing Products table
  customer_id INT REFERENCES Customers(id) ON DELETE SET NULL,  -- Foreign key referencing Customers table (optional)
  rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),  -- Rating of the product by the customer (1-5)
  review_text TEXT,  -- Text of the customer review
  created_at       timestamp without time zone,
  last_modified_at timestamp without time zone,
  deleted_at       timestamp without time zone
);

ALTER TABLE Reviews
ADD CONSTRAINT reviews_product_customer
UNIQUE (product_id, customer_id);

CREATE TABLE IF NOT EXISTS Chats (
  id SERIAL PRIMARY KEY,  -- Unique identifier for each review (auto-increment)
  customer_id INT REFERENCES Customers(id) ON DELETE SET NULL,
  chat TEXT,  -- Text of the customer review
  created_at       timestamp without time zone,
  deleted_at       timestamp without time zone
);
