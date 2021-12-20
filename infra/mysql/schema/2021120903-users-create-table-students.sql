CREATE TABLE `users`.`students` (
  `id`              VARCHAR(21)  NOT NULL,             -- ユーザーID
  `mail`            VARCHAR(256) NOT NULL,             -- メールアドレス
  `last_name`       VARCHAR(16)  NOT NULL,             -- 姓(漢字)
  `first_name`      VARCHAR(16)  NOT NUll,             -- 名(漢字)
  `last_name_kana`  VARCHAR(32)  NOT NULL,             -- 姓(かな)
  `first_name_kana` VARCHAR(32)  NOT NULL,             -- 名(かな)
  `birth_year`      INT(4)       NOT NULL,             -- 誕生年
  `created_at`      DATETIME     NOT NULL,             -- 登録日時
  `updated_at`      DATETIME     NOT NULL,             -- 更新日時
  `deleted_at`      DATETIME     NULL     DEFAULT NULL -- 退会日時
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_students_mail` ON `users`.`students` (`mail` ASC) VISIBLE;
