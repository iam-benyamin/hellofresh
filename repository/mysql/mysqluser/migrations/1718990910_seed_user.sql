-- +migrate Up
INSERT INTO `users` (`id`, `first_name`, `last_name`, `password`) VALUES ("a1b2c3d4e5f6", 'Benyamin', 'Mahmoudyan', 'password123');
INSERT INTO `users` (`id`, `first_name`, `last_name`, `password`) VALUES ("c3d4e5f6a7b8", 'Ali', 'Rezaei', 'password789');
INSERT INTO `users` (`id`, `first_name`, `last_name`, `password`) VALUES ("e5f6a7b8c9d0", 'Hossain', 'Nazari', 'password101');
INSERT INTO `users` (`id`, `first_name`, `last_name`, `password`) VALUES ("b7c8d9e0f1a2", 'Sara', 'Karimi', 'password456');
INSERT INTO `users` (`id`, `first_name`, `last_name`, `password`) VALUES ("d9e0f1a2b3c4", 'Fatemeh', 'Ahmadi', 'password102');


-- +migrate Down
DELETE FROM `users` WHERE id in ("a1b2c3d4e5f6", "b7c8d9e0f1a2", "c3d4e5f6a7b8", "e5f6a7b8c9d0", d9e0f1a2b3c4);
