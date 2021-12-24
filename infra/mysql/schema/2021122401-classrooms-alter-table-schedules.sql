CREATE TABLE `classrooms`.`schedules` (
  `weekday`    INT         NOT NULL DEFAULT 0, -- 曜日
  `is_closed`  TINYINT     NOT NULL,           -- 休校フラグ
  `lessons`    JSON        NOT NULL,           -- 授業コマ
  `created_at` DATETIME    NOT NULL,           -- 登録日時
  `updated_at` DATETIME    NOT NULL,           -- 更新日時
  PRIMARY KEY(`weekday`)
) ENGINE = InnoDB;
