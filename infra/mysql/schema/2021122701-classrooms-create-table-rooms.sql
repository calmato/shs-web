CREATE TABLE `classrooms`.`rooms` (
  `id`         INT      NOT NULL DEFAULT 0, -- ルームID
  `created_at` DATETIME NOT NULL,           -- 登録日時
  `updated_at` DATETIME NOT NULL,           -- 更新日時
  PRIMARY KEY(`id`)
) ENGINE = InnoDB;
