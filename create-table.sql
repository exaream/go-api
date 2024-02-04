CREATE TABLE IF NOT EXISTS articles (
  id integer UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  title varchar(100) NOT NULL,
  body text NOT NULL,
  user_name varchar(100) NOT NULL,
  nice_num integer NOT NULL,
  created_at datetime,
  updated_at datetime
);

CREATE TABLE IF NOT EXISTS comments (
  id integer UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  article_id integer UNSIGNED NOT NULL,
  body text NOT NULL,
  user_name varchar(100) NOT NULL,
  created_at datetime,
  updated_at datetime,
  FOREIGN KEY (article_id) REFERENCES articles(id)
);