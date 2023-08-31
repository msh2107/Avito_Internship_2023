CREATE TABLE users (
    id INT PRIMARY KEY
);

CREATE TABLE segments (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE user_segments (
    user_id INT,
    segment_id INT,
    PRIMARY KEY (user_id, segment_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segments(id) ON DELETE CASCADE
);