CREATE TABLE `lessons`.`student_shift_templates` (
  `student_id`        VARCHAR(22) NOT NULL, -- 生徒ID
  `schedules`         JSON        NOT NULL, -- 希望授業日時テンプレート
  `suggested_lessons` JSON        NOT NULL, -- 希望授業科目テンプレート
  `created_at`        DATETIME    NOT NULL, -- 登録日時
  `updated_at`        DATETIME    NOT NULL, -- 更新日時
  PRIMARY KEY(`student_id`)
) ENGINE = InnoDB;
