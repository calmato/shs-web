CREATE TABLE `lessons`.`lessons` (
  `id`               BIGINT      NOT NULL AUTO_INCREMENT, -- 授業ID
  `shift_summary_id` BIGINT      NOT NULL,                -- 授業スケジュールサマリID
  `shift_id`         BIGINT      NOT NULL,                -- 授業スケジュールID
  `subject_id`       BIGINT      NOT NULL,                -- 授業科目ID
  `room_id`          INT         NOT NULL,                -- 教室ID
  `teacher_id`       VARCHAR(22) NOT NULL,                -- 講師ID
  `student_id`       VARCHAR(22) NOT NULL,                -- 生徒ID
  `notes`            TEXT        NOT NULL,                -- 備考
  `created_at`       DATETIME    NOT NULL,                -- 登録日時
  `updated_at`       DATETIME    NOT NULL,                -- 更新日時
  PRIMARY KEY(`id`),
  CONSTRAINT `fk_lessons_shift_summary_id`
    FOREIGN KEY (`shift_summary_id`) REFERENCES `lessons`.`shift_summaries` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_lessons_shift_id`
    FOREIGN KEY (`shift_id`) REFERENCES `lessons`.`shifts` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_lessons_shift_id_room_id` ON `lessons`.`lessons` (`shift_id` ASC, `room_id` ASC) VISIBLE;
CREATE INDEX `idx_lessons_teacher_id` ON `lessons`.`lessons` (`teacher_id` ASC) VISIBLE;
CREATE INDEX `idx_lessons_student_id` ON `lessons`.`lessons` (`student_id` ASC) VISIBLE;
