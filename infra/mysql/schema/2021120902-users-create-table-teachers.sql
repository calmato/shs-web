CREATE SCHEMA IF NOT EXISTS `users` DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE `users`.`teachers` (
  `id`              VARCHAR(21)  NOT NULL,
  `mail`            VARCHAR(256) NOT NULL,
  `role`            INT          NOT NULL,
  `last_name`       VARCHAR(16)  NOT NULL,
  `first_name`      VARCHAR(16)  NOT NULL,
  `last_name_kana`  VARCHAR(32)  NOT NULL,
  `first_name_kana` VARCHAR(32)  NOT NULL,
  `created_at`      DATETIME     NOT NULL,
  `updated_at`      DATETIME     NOT NULL,
  `deleted_at`      DATETIME     NULL     DEFAULT NULL
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_teachers_mail` ON `users`.`teachers` (`mail` ASC) VISIBLE;
