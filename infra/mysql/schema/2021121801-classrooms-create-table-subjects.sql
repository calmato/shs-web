-- -----------------------------------------------
-- Database: classrooms
-- -----------------------------------------------
CREATE SCHEMA IF NOT EXISTS `classrooms` DEFAULT CHARACTER SET utf8mb4;

-- -----------------------------------------------
-- Table: subjects
-- -----------------------------------------------
CREATE TABLE `classrooms`.`subjects` (
  `id`         BIGINT      NOT NULL AUTO_INCREMENT, -- 授業科目ID
  `name`       VARCHAR(32) NOT NULL,                -- 授業科目名
  `color`      VARCHAR(7)  NOT NULL,                -- 表示色 (#rrggbb)
  `created_at` DATETIME    NOT NULL,                -- 登録日時
  `updated_at` DATETIME    NOT NULL,                -- 更新日時
  PRIMARY KEY(`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_subjects_name` ON `classrooms`.`subjects` (`name`) VISIBLE;
