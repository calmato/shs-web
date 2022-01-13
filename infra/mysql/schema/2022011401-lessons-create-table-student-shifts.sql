CREATE TABLE `lessons`.`student_submissions` (
  `student_id`       VARCHAR(22) NOT NULL, -- 生徒ID
  `shift_summary_id` BIGINT      NOT NULL, -- 授業スケジュールサマリID
  `decided`          TINYINT     NOT NULL, -- 確定フラグ
  `created_at`       DATETIME    NOT NULL, -- 登録日時
  `updated_at`       DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`student_id`, `shift_summary_id` DESC),
  CONSTRAINT `fk_student_submissions_shift_summaries_id`
    FOREIGN KEY (`shift_summary_id`) REFERENCES `lessons`.`shift_summaries` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE TABLE `lessons`.`student_shifts` (
  `student_id`       VARCHAR(22) NOT NULL, -- 生徒ID
  `shift_summary_id` BIGINT      NOT NULL, -- 授業スケジュールサマリID
  `shift_id`         BIGINT      NOT NULL, -- 授業スケジュールID
  `created_at`       DATETIME    NOT NULL, -- 登録日時
  `updated_at`       DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`student_id`, `shift_id` DESC),
  CONSTRAINT `fk_student_shifts_shift_summaries_id`
    FOREIGN KEY (`shift_summary_id`) REFERENCES `lessons`.`shift_summaries` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_student_shifts_shifts_id`
    FOREIGN KEY (`shift_id`) REFERENCES `lessons`.`shifts` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idex_student_shifts_shift_summary_id_student_id`
  ON `lessons`.`student_shifts` (`shift_summary_id` DESC, `student_id` ASC) VISIBLE;
CREATE INDEX `idex_student_shifts_shift_id_student_id`
  ON `lessons`.`student_shifts` (`shift_id` DESC, `student_id` ASC) VISIBLE;
