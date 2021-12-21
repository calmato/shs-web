DROP TABLE `classrooms`.`subjects_teachers`;

CREATE TABLE `classrooms`.`teacher_subjects` (
  `teacher_id` VARCHAR(22) NOT NULL, -- 講師ID
  `subject_id` BIGINT      NOT NULL, -- 授業科目ID
  `created_at` DATETIME    NOT NULL, -- 登録日時
  `updated_at` DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`teacher_id`, `subject_id`),
  CONSTRAINT `fk_teacher_subjects_subjects_id`
    FOREIGN KEY (`subject_id`) REFERENCES `classrooms`.`subjects` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idx_teacher_subjects_subject_id_teacher_id`
  ON `classrooms`.`teacher_subjects` (`subject_id` ASC, `teacher_id` ASC) VISIBLE;
