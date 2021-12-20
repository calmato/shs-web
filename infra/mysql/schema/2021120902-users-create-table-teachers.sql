CREATE SCHEMA `users` DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE `users`.`teachers` (
  `id`              VARCHAR(21)  NOT NULL,             -- ユーザーID
  `mail`            VARCHAR(256) NOT NULL,             -- メールアドレス
  `role`            INT          NOT NULL,             -- 権限
  `last_name`       VARCHAR(16)  NOT NULL,             -- 姓
  `first_name`      VARCHAR(16)  NOT NULL,             -- 名
  `last_name_kana`  VARCHAR(32)  NOT NULL,             -- 姓(かな)
  `first_name_kana` VARCHAR(32)  NOT NULL,             -- 名(かな)
  `created_at`      DATETIME     NOT NULL,             -- 登録日時
  `updated_at`      DATETIME     NOT NULL,             -- 更新日時
  `deleted_at`      DATETIME     NULL     DEFAULT NULL -- 退会日時
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_teachers_mail` ON `users`.`teachers` (`mail` ASC) VISIBLE;
