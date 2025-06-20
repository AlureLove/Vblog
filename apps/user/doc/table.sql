CREATE TABLE `users`
(
    `id`         bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NOT NULL,
    `created_by` varchar(100),
    `updated_at` datetime(3) NULL,
    `username`   varchar(191),
    `password`   varchar(255),
    `avatar`     varchar(255),
    `nickname`   varchar(100),
    `email`      varchar(100),
    PRIMARY KEY (`id`),
    INDEX        `idx_users_username` (`username`),
    CONSTRAINT `uni_users_username` UNIQUE (`username`)
) component:DATASOURCE