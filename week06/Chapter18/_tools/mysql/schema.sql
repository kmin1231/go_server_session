CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'user identifier',
    `name`     varchar(20) NOT NULL COMMENT 'user name -- testing',
    `password` VARCHAR(80) NOT NULL COMMENT 'password hash',
    `role`     VARCHAR(80) NOT NULL COMMENT 'role',
    `created`  DATETIME(6) NOT NULL COMMENT 'record creation time',
    `modified` DATETIME(6) NOT NULL COMMENT 'record update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user';

CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'task identifier',
    `title`    VARCHAR(128) NOT NULL COMMENT 'task title',
    `status`   VARCHAR(20)  NOT NULL COMMENT 'task status',
    `created`  DATETIME(6) NOT NULL COMMENT 'task creation time',
    `modified` DATETIME(6) NOT NULL COMMENT 'task update time',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='task';