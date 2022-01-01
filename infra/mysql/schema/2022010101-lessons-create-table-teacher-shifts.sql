CREATE TABLE `lessons`.`teacher_submissions` (
  `teacher_id`       VARCHAR(22) NOT NULL, -- 講師ID
  `shift_summary_id` BIGINT      NOT NULL, -- 授業スケジュールサマリID
  `decided`          TINYINT     NOT NULL, -- 確定フラグ
  `created_at`       DATETIME    NOT NULL, -- 登録日時
  `updated_at`       DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`teacher_id`, `shift_summary_id` DESC),
  CONSTRAINT `fk_teacher_submissions_shift_summaries_id`
    FOREIGN KEY (`shift_summary_id`) REFERENCES `lessons`.`shift_summaries` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `lessons`.`teacher_shifts` (
  `teacher_id`       VARCHAR(22) NOT NULL, -- 講師ID
  `shift_summary_id` BIGINT      NOT NULL, -- 授業スケジュールサマリID
  `shift_id`         BIGINT      NOT NULL, -- 授業スケジュールID
  `created_at`       DATETIME    NOT NULL, -- 登録日時
  `updated_at`       DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`teacher_id`, `shift_id` DESC),
  CONSTRAINT `fk_teacher_shifts_shift_summaries_id`
    FOREIGN KEY (`shift_summary_id`) REFERENCES `lessons`.`shift_summaries` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_teacher_shifts_shifts_id`
    FOREIGN KEY (`shift_id`) REFERENCES `lessons`.`shifts` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idex_teacher_shifts_shift_summary_id_teacher_id`
  ON `lessons`.`teacher_shifts` (`shift_summary_id` DESC, `teacher_id` ASC) VISIBLE;
CREATE INDEX `idex_teacher_shifts_shift_id_teacher_id`
  ON `lessons`.`teacher_shifts` (`shift_id` DESC, `teacher_id` ASC) VISIBLE;
