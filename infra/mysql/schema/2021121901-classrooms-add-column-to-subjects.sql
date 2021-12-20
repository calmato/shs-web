ALTER TABLE `classrooms`.`subjects` DROP INDEX `ui_subjects_name`;

ALTER TABLE `classrooms`.`subjects` ADD COLUMN `school_type` INT NOT NULL AFTER `name`;
CREATE UNIQUE INDEX `ui_subjects_name_school_type` ON `classrooms`.`subjects` (`name`, `school_type`) VISIBLE;
CREATE INDEX `idx_subjects_name_school_type` ON `classrooms`.`subjects` (`name` ASC, `school_type` ASC) VISIBLE;
CREATE INDEX `idx_subjects_school_type_name` ON `classrooms`.`subjects` (`school_type` ASC, `name` ASC) VISIBLE;
