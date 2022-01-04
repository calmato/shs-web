CREATE TABLE `classrooms`.`student_subjects` (
  `student_id` VARCHAR(22) NOT NULL, -- 生徒ID
  `subject_id` BIGINT      NOT NULL, -- 授業科目ID
  `created_at` DATETIME    NOT NULL, -- 登録日時
  `updated_at` DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`student_id`, `subject_id`),
  CONSTRAINT `fk_student_subjects_subjects_id`
    FOREIGN KEY (`subject_id`) REFERENCES `classrooms`.`subjects` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idx_student_subjects_subject_id_student_id`
  ON `classrooms`.`student_subjects` (`subject_id` ASC, `student_id` ASC) VISIBLE;
