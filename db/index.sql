CREATE TABLE IF NOT EXISTS test_db (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT,
    age INT,
    profession TEXT
);

INSERT INTO test_db (name, age, profession) VALUES
('John Doe', 28, 'Engineer'),
('Jane Smith', 34, 'Doctor'),
('Emily Johnson', 22, 'Student'),
('Michael Brown', 45, 'Teacher'),
('Sarah Davis', 30, 'Architect'),
('Chris Wilson', 38, 'Designer'),
('Patricia Garcia', 26, 'Nurse'),
('Robert Martinez', 50, 'Manager'),
('Linda Rodriguez', 29, 'Scientist'),
('James Hernandez', 41, 'Lawyer');

