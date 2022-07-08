CREATE TABLE groups(
  group_id INTEGER PRIMARY KEY  AUTOINCREMENT,
  user_id         INTEGER       NOT NULL,
  created_at      DATETIME      NOT NULL,
  title           VARCHAR(50)   NOT NULL,
  body            VARCHAR(255)  NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);
