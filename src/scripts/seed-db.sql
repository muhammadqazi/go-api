INSERT INTO departments_entity (department_code, name, description, dean, vice_dean, email, phone_number, dean_email, dean_phone, created_at, is_active)
VALUES
    ('D001', 'Computer Science', 'Study of computers and computational systems', 'Dr. John Doe', 'Dr. Jane Doe', 'cs@university.com', '555-555-5555', 'dean_cs@university.com', '555-555-5551', current_timestamp , true),
    ('D002', 'Mathematics', 'Study of numbers and shapes', 'Dr. Tom Doe', 'Dr. Sarah Doe', 'math@university.com', '555-555-5556', 'dean_math@university.com', '555-555-5552', current_timestamp , true),
    ('D003', 'Physics', 'Study of matter and energy', 'Dr. James Doe', 'Dr. Emily Doe', 'physics@university.com', '555-555-5557', 'dean_physics@university.com', '555-555-5553', current_timestamp , true),
    ('D004', 'Chemistry', 'Study of chemicals and their reactions', 'Dr. David Doe', 'Dr. Lily Doe', 'chemistry@university.com', '555-555-5558', 'dean_chemistry@university.com', '555-555-5554', current_timestamp , true),
    ('D005', 'Biology', 'Study of living organisms', 'Dr. Michael Doe', 'Dr. Grace Doe', 'biology@university.com', '555-555-5559', 'dean_biology@university.com', '555-555-5555', current_timestamp , true);


INSERT INTO faculties_entity (name, code, description, dean, vice_dean, email, phone_number, dean_email, dean_phone, department_id , created_at, is_active)
VALUES
    ('Mathematics', 'MATH', 'Department of Mathematics', 'John Doe', 'Jane Doe', 'math@example.com', '555-555-5555', 'johndoe@example.com', '555-555-5556', 1 , current_timestamp , true),
    ('Computer Science', 'CS', 'Department of Computer Science', 'Jane Smith', 'John Smith', 'cs@example.com', '555-555-5557', 'janesmith@example.com', '555-555-5558', 2 , current_timestamp , true),
    ('Physics', 'PHYS', 'Department of Physics', 'John Johnson', 'Jane Johnson', 'physics@example.com', '555-555-5559', 'johnjohnson@example.com', '555-555-5560', 3 , current_timestamp , true),
    ('Chemistry', 'CHEM', 'Department of Chemistry', 'Jane Wilson', 'John Wilson', 'chemistry@example.com', '555-555-5561', 'janewilson@example.com', '555-555-5562', 4 , current_timestamp , true),
    ('Biology', 'BIO', 'Department of Biology', 'John Davis', 'Jane Davis', 'biology@example.com', '555-555-5563', 'johndavis@example.com', '555-555-5564', 5 , current_timestamp , true);


INSERT INTO students_entity (student_id, first_name, surname, email, nationality, dob, place_of_birth, sex, password, role, status, access_status, acceptance_type, semester, graduation_date, is_graduated, department_id, created_at, is_active)
VALUES
    (21906778, 'John', 'Doe', 'johndoe@example.com', 'American', '01-01-1998', 'New York', 'Male', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', 'Student', 'Enrolled', 'Active', 'Admitted', 'Fall 2021', '2025-05-25', false, 1, current_timestamp, true),
    (21906779, 'John', 'Doe', 'johndoe2@example.com', 'American', '01-01-1998', 'New York', 'Male', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', 'Student', 'Enrolled', 'Active', 'Admitted', 'Fall 2021', '2025-05-25', false, 1, current_timestamp, true);

INSERT INTO accounts_entity (approaching_dept, current_dept, discount, discount_type, installments, scholarship, student_id, total_dept, total_fee)
VALUES
    (0, 0, 10, 'Sports', 2, 50, 21906780, 1575, 3500),
    (0, 0, 10, 'Sports', 2, 50, 21906780, 1575, 3500);