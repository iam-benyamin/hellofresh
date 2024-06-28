-- +migrate Up
CREATE TABLE `orders`
(
    `id`                 CHAR(12) PRIMARY KEY,
    `user_id`            CHAR(12)       NOT NULL,
    `product_code`       VARCHAR(191)   NOT NULL,
    `customer_full_name` VARCHAR(383)   NOT NULL,
    `product_name`       VARCHAR(191)   NOT NULL,
    `total_amount`       DECIMAL(10, 2) NOT NULL,
    `created_at`         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `orders`;
