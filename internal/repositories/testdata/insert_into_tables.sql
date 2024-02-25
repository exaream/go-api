INSERT INTO articles (title, body, user_name, nice_num, created_at, updated_at) VALUES
	('article title 1', 'article body 1', 'Alice', 2, NOW(), NOW());

INSERT INTO articles (title, body, user_name, nice_num, created_at, updated_at) VALUES
	('article title 2', 'article body 2', 'Alice', 4, NOW(), NOW());

INSERT INTO comments (article_id, body, user_name, created_at, updated_at) values
	(1, 'article 1, comment 1', "Bob", NOW(), NOW());

INSERT INTO comments (article_id, body, user_name, created_at, updated_at) values
	(1, 'article 1, comment 2', "Chris", NOW(), NOW());

