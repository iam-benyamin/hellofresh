-- +migrate Up
CREATE TABLE `users`
(
    `id`         VARCHAR(36) PRIMARY KEY,
    `first_name` VARCHAR(191) NOT NULL,
    `last_name`  VARCHAR(191) NOT NULL,
    `password`   VARCHAR(191) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `users`;
