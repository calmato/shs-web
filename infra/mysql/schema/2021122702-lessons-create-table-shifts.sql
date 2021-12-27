CREATE TABLE `lessons`.`shift_summaries` (
  `id`         BIGINT   NOT NULL AUTO_INCREMENT, -- 授業スケジュールサマリID
  `year_month` INT      NOT NULL,                -- 年月
  `open_at`    DATETIME NOT NULL,                -- 募集開始日時
  `end_at`     DATETIME NOT NULL,                -- 募集締切日時
  `created_at` DATETIME NOT NULL,                -- 登録日時
  `updated_at` DATETIME NOT NULL,                -- 更新日時
  PRIMARY KEY(`id`)
) ENGINE = InnoDB;

CREATE TABLE `lessons`.`shifts` (
  `id`               BIGINT     NOT NULL AUTO_INCREMENT, -- 授業スケジュールID
  `shift_summary_id` BIGINT     NOT NULL,                -- 授業スケジュールサマリID
  `date`             DATE       NOT NULL,                -- 授業日
  `start_time`       VARCHAR(4) NOT NULL,                -- 授業開始時間
  `end_time`         VARCHAR(4) NOT NULL,                -- 授業終了時間
  `created_at`       DATETIME   NOT NULL,                -- 登録日時
  `updated_at`       DATETIME   NOT NULL,                -- 更新日時
  PRIMARY KEY(`id`),
  CONSTRAINT `fk_shifts_shift_summaries_id`
    FOREIGN KEY (`shift_summary_id`) REFERENCES `lessons`.`shift_summaries` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_shift_summaries_year_month`
  ON `lessons`.`shift_summaries` (`year_month` DESC) VISIBLE;
