CREATE TABLE users(
  user_id INTEGER PRIMARY KEY AUTOINCREMENT,
  email         VARCHAR(50)   NOT NULL,
  password      VARCHAR(50)   NOT NULL,
  first_name    VARCHAR(50)   NOT NULL,
  last_name     VARCHAR(50)   NOT NULL,
  date_of_birth DATETIME      NOT NULL,
  is_private    BOOLEAN       NOT NULL,
  avatar        VARCHAR(255),
  nickname      VARCHAR(50),
  about         VARCHAR(255)
);
