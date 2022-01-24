ALTER TABLE `lessons`.`student_submissions` DROP COLUMN `suggested_classes`;
ALTER TABLE `lessons`.`student_submissions` ADD COLUMN `suggested_lessons` JSON NOT NULL AFTER `decided`;
