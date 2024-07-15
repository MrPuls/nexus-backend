CREATE TABLE projects (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT,
                       created_at TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'UTC')
);

INSERT INTO projects (name, description) VALUES
                                        ('Demo project 1', 'demo project 1 description'),
                                        ('Demo project 2', 'demo project 2 description');