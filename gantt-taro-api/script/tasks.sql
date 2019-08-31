CREATE TABLE IF NOT EXISTS gantt_taro.tasks(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    user_id INT,
    project_id INT,
    parent_task_id INT,
    started_at DATE,
    ended_at DATE
);

INSERT INTO gantt_taro.tasks(name, started_at, ended_at)
VALUES 
('tasks1', '2019-08-05', '2019-08-23'),
('tasks2', '2019-09-03', '2019-09-27'),
('tasks3', '2019-08-01', '2019-08-31'),
('tasks4', '2019-09-07', '2019-09-09'),
('tasks5', '2019-09-05', '2019-09-06'),
('tasks6', '2019-11-05', '2019-11-11');