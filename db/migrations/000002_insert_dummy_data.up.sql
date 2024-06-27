INSERT INTO Customers (first_name, last_name, email, phone_number, created_at, last_modified_at, deleted_at)
VALUES
  ('John', 'Doe', 'john.doe@email.com', '555-123-4567', '2024-06-26 10:00:00', '2024-06-26 10:00:00', NULL),
  ('Jane', 'Smith', 'jane.smith@email.com', '555-987-6543', '2024-06-25 14:30:00', '2024-06-26 09:15:00', NULL),
  ('Michael', 'Johnson', 'michael.j@email.com', '555-246-8135', '2024-06-24 11:45:00', '2024-06-24 11:45:00', NULL),
  ('Emily', 'Brown', 'emily.brown@email.com', '555-369-2580', '2024-06-23 16:20:00', '2024-06-25 13:10:00', NULL),
  ('David', 'Wilson', 'david.wilson@email.com', '555-147-2589', '2024-06-22 09:30:00', '2024-06-22 09:30:00', '2024-06-26 11:00:00'),
  ('Sarah', 'Taylor', 'sarah.t@email.com', '555-753-9514', '2024-06-21 13:15:00', '2024-06-23 10:45:00', NULL),
  ('Robert', 'Anderson', 'robert.a@email.com', '555-951-7532', '2024-06-20 15:40:00', '2024-06-20 15:40:00', NULL),
  ('Jennifer', 'Martinez', 'jennifer.m@email.com', '555-357-1593', '2024-06-19 12:00:00', '2024-06-24 14:30:00', NULL),
  ('William', 'Garcia', 'william.g@email.com', '555-852-9630', '2024-06-18 10:30:00', '2024-06-18 10:30:00', '2024-06-25 16:45:00'),
  ('Lisa', 'Lopez', 'lisa.lopez@email.com', '555-741-8520', '2024-06-17 17:00:00', '2024-06-22 11:20:00', NULL);


INSERT INTO Products (name, description, price, category, brand, stock, image_url, created_at, last_modified_at, deleted_at)
VALUES
  ('Smartphone X', 'Latest model with advanced features', 799.99, 'Electronics', 'TechBrand', 50, 'http://example.com/images/smartphone_x.jpg', '2024-06-26 10:00:00', '2024-06-26 10:00:00', NULL),
  ('Laptop Pro', 'High-performance laptop for professionals', 1299.99, 'Computers', 'CompTech', 30, 'http://example.com/images/laptop_pro.jpg', '2024-06-25 14:30:00', '2024-06-26 09:15:00', NULL),
  ('Wireless Headphones', 'Noise-cancelling over-ear headphones', 199.99, 'Audio', 'SoundMaster', 100, 'http://example.com/images/wireless_headphones.jpg', '2024-06-24 11:45:00', '2024-06-24 11:45:00', NULL),
  ('4K TV', 'Ultra HD Smart TV with HDR', 999.99, 'Electronics', 'VisionPlus', 20, 'http://example.com/images/4k_tv.jpg', '2024-06-23 16:20:00', '2024-06-25 13:10:00', NULL),
  ('Gaming Console', 'Next-gen gaming console with 1TB storage', 499.99, 'Gaming', 'GameWorld', 75, 'http://example.com/images/gaming_console.jpg', '2024-06-22 09:30:00', '2024-06-22 09:30:00', '2024-06-26 11:00:00'),
  ('Smartwatch', 'Fitness tracker with heart rate monitor', 149.99, 'Wearables', 'FitTech', 200, 'http://example.com/images/smartwatch.jpg', '2024-06-21 13:15:00', '2024-06-23 10:45:00', NULL),
  ('Bluetooth Speaker', 'Portable speaker with deep bass', 79.99, 'Audio', 'SoundMaster', 150, 'http://example.com/images/bluetooth_speaker.jpg', '2024-06-20 15:40:00', '2024-06-20 15:40:00', NULL),
  ('Digital Camera', 'DSLR camera with 24MP sensor', 599.99, 'Cameras', 'PhotoPro', 40, 'http://example.com/images/digital_camera.jpg', '2024-06-19 12:00:00', '2024-06-24 14:30:00', NULL),
  ('Tablet', '10-inch tablet with 64GB storage', 299.99, 'Computers', 'CompTech', 60, 'http://example.com/images/tablet.jpg', '2024-06-18 10:30:00', '2024-06-18 10:30:00', '2024-06-25 16:45:00'),
  ('Fitness Tracker', 'Waterproof fitness tracker with GPS', 99.99, 'Wearables', 'FitTech', 120, 'http://example.com/images/fitness_tracker.jpg', '2024-06-17 17:00:00', '2024-06-22 11:20:00', NULL);
