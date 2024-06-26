-- +migrate Up
CREATE TABLE `products`
(
    `id`           VARCHAR(36) PRIMARY KEY,
    `product_name` VARCHAR(191)   NOT NULL,
    `product_code` VARCHAR(191)   NOT NULL,
    `price`        DECIMAL(10, 2) NOT NULL,
    `created_at`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `products`;
