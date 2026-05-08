-- Worker table
CREATE TABLE IF NOT EXISTS worker (
    name VARCHAR(255) PRIMARY KEY,
    max_process INT NOT NULL
);

-- Queue table
CREATE TABLE IF NOT EXISTS queue (
    name VARCHAR(255) PRIMARY KEY,
    worker_name VARCHAR(255) NOT NULL,
    process_order INT NOT NULL,
    FOREIGN KEY (worker_name) REFERENCES worker(name)
);

-- Job table
CREATE TABLE IF NOT EXISTS job (
    id VARCHAR(255) PRIMARY KEY,
    queue VARCHAR(255) NOT NULL,
    source VARCHAR(255) NOT NULL,
    procedure VARCHAR(50) NOT NULL,
    json_params TEXT NOT NULL,
    json_result TEXT NOT NULL,
    error_json TEXT NOT NULL,
    register_date TIMESTAMP NOT NULL,
    start_date TIMESTAMP,
    finish_date TIMESTAMP,
    status_code BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (queue) REFERENCES queue(name)
);

-- Create indexes
CREATE INDEX idx_job_queue ON job(queue);
CREATE INDEX idx_job_register_date ON job(register_date);
