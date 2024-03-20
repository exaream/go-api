TRUNCATE TABLE `comments`;

-- You can not use TRUNCATE due to foreign key constraints.
DELETE FROM `articles`;
ALTER TABLE `articles` AUTO_INCREMENT = 1;
