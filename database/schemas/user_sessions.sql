CREATE TABLE IF NOT EXISTS user_sessions (
 session_key VARCHAR(128) PRIMARY KEY,
 user_id INT NOT NULL, -- Make this a key?
 login_time TIMESTAMP NOT NULL,
 last_seen_time TIMESTAMP NOT NULL,
 FOREIGN KEY (user_id) REFERENCES users(user_id)
)
