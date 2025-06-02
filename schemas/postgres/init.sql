CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    fullname VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    telephone VARCHAR(20),
    birth_date DATE NOT NULL,
    register_date TIMESTAMP NOT NULL DEFAULT now(),
    last_access TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE post (
	post_id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	text TEXT NOT NULL,
	url_foto TEXT,
	like_count INTEGER DEFAULT 0
);

CREATE TABLE post_like (
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	liked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, post_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE
);

CREATE TABLE post_comment(
	id_comment SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	text TEXT NOT NULL,
	commented_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE
);

CREATE TABLE music_history (
    id_history SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    music_id INTEGER NOT NULL,
    listened_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE community (
	community_id SERIAL PRIMARY KEY,
	community_name VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	description TEXT
);

CREATE TABLE community_posts(
	community_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	PRIMARY KEY (community_id, post_id),
	FOREIGN KEY (community_id) REFERENCES community(community_id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE
);

CREATE TABLE user_community (
	community_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	PRIMARY KEY (community_id, user_id),
	FOREIGN KEY (community_id) REFERENCES community(community_id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE chat (
	chat_id SERIAL PRIMARY KEY,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat_participants (
	user_id INTEGER NOT NULL,
	chat_id INTEGER NOT NULL,
	PRIMARY KEY (user_id, chat_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (chat_id) REFERENCES chat(chat_id) ON DELETE CASCADE
);

CREATE TABLE chat_message (
	message_id SERIAL PRIMARY KEY,
	author_id INTEGER,
	chat_id INTEGER NOT NULL,
	message TEXT NOT NULL,
	sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (author_id) REFERENCES users(user_id) ON DELETE SET NULL,
	FOREIGN KEY (chat_id) REFERENCES chat(chat_id) ON DELETE CASCADE
);