-- +migrate Up
INSERT INTO `products` (`id`, `product_name`, `product_code`, `price`) VALUES ("a1b2c3d4e5f6", 'Tacos', 'classic-box', 49.99);
INSERT INTO `products` (`id`, `product_name`, `product_code`, `price`) VALUES ("b2c3d4e5f6g7", 'Bagel Bonanza', 'family-box', 249.99);
INSERT INTO `products` (`id`, `product_name`, `product_code`, `price`) VALUES ("d4e5f6g7h8i9", 'Roasted Veggie Medley', 'veggie-box', 99.99);

-- +migrate Down
DELETE FROM `products` WHERE id in ("a1b2c3d4e5f6", "b2c3d4e5f6g7", "d4e5f6g7h8i9");
