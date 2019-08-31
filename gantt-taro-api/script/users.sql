CREATE TABLE IF NOT EXISTS gantt_taro.users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    unique_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP NOT NULL default current_timestamp on update current_timestamp
);
CREATE INDEX unique_id_index ON gantt_taro.users(unique_id);
