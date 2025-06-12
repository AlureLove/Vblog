CREATE TABLE `blogs`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) NOT NULL,
    `created_by` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `title`      varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `summary`    text COLLATE utf8mb4_general_ci,
    `content`    text COLLATE utf8mb4_general_ci,
    `category`   varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `tags`       longtext COLLATE utf8mb4_general_ci,
    `stage`      tinyint(1) DEFAULT NULL,
    `changed_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY          `idx_blogs_category` (`category`),
    KEY          `idx_blogs_stage` (`stage`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci