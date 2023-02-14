INSERT INTO departments_entity (department_code, name, description, dean, vice_dean, email, phone_number, dean_email, dean_phone, created_at, is_active, offered_semesters , number_of_years)
VALUES
    ('D001', 'Computer Science', 'Study of computers and computational systems', 'Dr. John Doe', 'Dr. Jane Doe', 'cs@university.com', '555-555-5555', 'dean_cs@university.com', '555-555-5551', current_timestamp , true, '{"Fall","Spring"}',4),
    ('D002', 'Mathematics', 'Study of numbers and shapes', 'Dr. Tom Doe', 'Dr. Sarah Doe', 'math@university.com', '555-555-5556', 'dean_math@university.com', '555-555-5552', current_timestamp , true,'{"Fall","Spring"}',4),
    ('D003', 'Physics', 'Study of matter and energy', 'Dr. James Doe', 'Dr. Emily Doe', 'physics@university.com', '555-555-5557', 'dean_physics@university.com', '555-555-5553', current_timestamp , true,'{"Fall","Spring"}',4),
    ('D004', 'Chemistry', 'Study of chemicals and their reactions', 'Dr. David Doe', 'Dr. Lily Doe', 'chemistry@university.com', '555-555-5558', 'dean_chemistry@university.com', '555-555-5554', current_timestamp , true,'{"Fall","Spring"}',4),
    ('D005', 'Biology', 'Study of living organisms', 'Dr. Michael Doe', 'Dr. Grace Doe', 'biology@university.com', '555-555-5559', 'dean_biology@university.com', '555-555-5555', current_timestamp , true,'{"Fall","Spring"}',4);


INSERT INTO faculties_entity (name, code, description, dean, vice_dean, email, phone_number, dean_email, dean_phone, department_id , created_at, is_active)
VALUES
    ('Mathematics', 'MATH', 'Department of Mathematics', 'John Doe', 'Jane Doe', 'math@example.com', '555-555-5555', 'johndoe@example.com', '555-555-5556', 1 , current_timestamp , true),
    ('Computer Science', 'CS', 'Department of Computer Science', 'Jane Smith', 'John Smith', 'cs@example.com', '555-555-5557', 'janesmith@example.com', '555-555-5558', 2 , current_timestamp , true),
    ('Physics', 'PHYS', 'Department of Physics', 'John Johnson', 'Jane Johnson', 'physics@example.com', '555-555-5559', 'johnjohnson@example.com', '555-555-5560', 3 , current_timestamp , true),
    ('Chemistry', 'CHEM', 'Department of Chemistry', 'Jane Wilson', 'John Wilson', 'chemistry@example.com', '555-555-5561', 'janewilson@example.com', '555-555-5562', 4 , current_timestamp , true),
    ('Biology', 'BIO', 'Department of Biology', 'John Davis', 'Jane Davis', 'biology@example.com', '555-555-5563', 'johndavis@example.com', '555-555-5564', 5 , current_timestamp , true);

INSERT INTO instructors_entity (first_name, last_name, email, phone_number, password, dob, place_of_birth, sex, nationality, role , created_at,is_active)
VALUES
    ('John', 'Doe', 'johndoe@example.com', '555-555-5555', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1980-01-01', 'New York', 'Male', 'American', 'Instructor',current_timestamp,true),
    ('Jane', 'Doe', 'janedoe@example.com', '555-555-5556', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1981-02-02', 'Los Angeles', 'Female', 'American', 'Instructor',current_timestamp,true),
    ('Bob', 'Smith', 'bobsmith@example.com', '555-555-5557', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1982-03-03', 'Chicago', 'Male', 'American', 'Instructor',current_timestamp,true),
    ('Alice', 'Johnson', 'alicejohnson@example.com', '555-555-5558', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1983-04-04', 'Houston', 'Female', 'American', 'Instructor',current_timestamp,true),
    ('Tom', 'Williams', 'tomwilliams@example.com', '555-555-5559', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1984-05-05', 'Philadelphia', 'Male', 'American', 'Instructor',current_timestamp,true),
    ('Emily', 'Jones', 'emilyjones@example.com', '555-555-5560', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1985-06-06', 'Phoenix', 'Female', 'American', 'Instructor',current_timestamp,true),
    ('David', 'Brown', 'davidbrown@example.com', '555-555-5561', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1986-07-07', 'San Antonio', 'Male', 'American', 'Instructor',current_timestamp,true),
    ('Sophie', 'Davis', 'sophiedavis@example.com', '555-555-5562', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1987-08-08', 'San Diego', 'Female', 'American', 'instructor',current_timestamp,true),
    ('Jacob', 'Miller', 'oliviawilson@example.com2', '555-555-5563', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1988-09-09', 'Dallas', 'Male', 'American', 'admin',current_timestamp,true),
    ( 'Olivia', 'Wilson', 'oliviawilson@example.com', '555-555-5564', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim0', '1989-10-10', 'San Jose', 'Female', 'American', 'admin',current_timestamp,true);

INSERT INTO students_entity (student_id, first_name, surname, email, nationality, dob, place_of_birth, sex, password, role, status, access_status, acceptance_type, semester, graduation_date, is_graduated, department_id, supervisor_id , created_at, is_active)
VALUES
    (21906778, 'John', 'Doe', 'johndoe@example.com', 'American', '01-01-1998', 'New York', 'Male', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', 'student', 'registered', 'active', 'Admitted', 'Fall 2021', '2025-05-25', false, 1, 10, current_timestamp, true),
    (21906779, 'John', 'Doe', 'johndoe2@example.com', 'American', '01-01-1998', 'New York', 'Male', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', 'student', 'registered', 'active', 'Admitted', 'Fall 2021', '2025-05-25', false, 1, 10,current_timestamp, true);

INSERT INTO accounts_entity (approaching_dept, current_dept, discount, discount_type, installments, scholarship, student_id, total_dept, total_fee)
VALUES
    (0, 0, 10, 'Sports', 2, 50, 21906778, 1575, 3500),
    (0, 0, 10, 'Sports', 2, 50, 21906778, 1575, 3500);

INSERT INTO courses_entity (name, code, description, credits, ects, theoretical, practical, created_at, is_active)
VALUES
    ('Introduction to Computer Science', 'CS101', 'An overview of computer science', 3, 6, 45, 15, current_timestamp, true),
    ('Programming Fundamentals', 'CS102', 'Introduction to programming concepts', 4, 8, 60, 20, current_timestamp, true),
    ('Data Structures and Algorithms', 'CS203', 'Study of data structures and algorithms', 4, 8, 60, 20, current_timestamp, true),
    ('Web Development', 'CS304', 'Introduction to web development', 4, 8, 60, 20, current_timestamp, true),
    ('Database Systems', 'CS405', 'Introduction to database systems', 3, 6, 45, 15, current_timestamp, true),
    ('Operating Systems', 'CS406', 'Study of operating systems', 4, 8, 60, 20, current_timestamp, true),
    ('Computer Networks', 'CS407', 'Introduction to computer networks', 3, 6, 45, 15, current_timestamp, true),
    ('Object-Oriented Programming', 'CS408', 'Introduction to object-oriented programming', 4, 8, 60, 20, current_timestamp, true),
    ('Software Engineering', 'CS409', 'Study of software engineering concepts', 4, 8, 60, 20, current_timestamp, true),
    ( 'Artificial Intelligence', 'CS410', 'Introduction to artificial intelligence', 3, 6, 45, 15, current_timestamp, true),
    ( 'Machine Learning', 'CS411', 'Study of machine learning algorithms', 4, 8, 60, 20, current_timestamp, true),
    ( 'Data Science', 'CS412', 'Introduction to data science', 3, 6, 45, 15, current_timestamp, true),
    ( 'Cloud Computing', 'CS413', 'Introduction to cloud computing', 4, 8, 60, 20, current_timestamp, true),
    ( 'Cybersecurity', 'CS414', 'Introduction to cybersecurity', 3, 6, 45, 15, current_timestamp, true),
    ( 'Blockchain Technology', 'CS415', 'Introduction to blockchain technology', 4, 8, 60, 20, current_timestamp, true),
    ( 'Virtual Reality', 'CS416', 'Introduction to virtual reality', 3, 6, 45, 15, current_timestamp, true),
    ( 'Augmented Reality', 'CS417', 'Introduction to augmented reality', 4, 8, 60, 20, current_timestamp, true),
    ( 'Internet of Things', 'CS418', 'Introduction to internet of things', 3, 6, 45, 15, current_timestamp, true),
    ( 'Mobile Application Development', 'CS419', 'Introduction to mobile application development', 4, 8, 60, 20, current_timestamp, true),
    ( 'Quantum Computing', 'CS420', 'Introduction to quantum computing', 3, 6, 45, 15, current_timestamp,true),
    ( 'Human-Computer Interaction', 'CS421', 'Introduction to human-computer interaction', 4, 8, 60, 20, current_timestamp, true),
    ( 'Computer Graphics', 'CS422', 'Introduction to computer graphics', 3, 6, 45, 15, current_timestamp, true),
    ( 'Compiler Design', 'CS423', 'Study of compiler design', 4, 8, 60, 20, current_timestamp, true),
    ( 'Parallel Computing', 'CS424', 'Introduction to parallel computing', 3, 6, 45, 15, current_timestamp, true),
    ( 'Computer Vision', 'CS425', 'Introduction to computer vision', 4, 8, 60, 20, current_timestamp, true),
    ( 'Natural Language Processing', 'CS426', 'Introduction to natural language processing', 3, 6, 45, 15, current_timestamp, true),
    ( 'Information Retrieval', 'CS427', 'Introduction to information retrieval', 4, 8, 60, 20, current_timestamp, true),
    ( 'Cryptography', 'CS428', 'Introduction to cryptography', 3, 6, 45, 15, current_timestamp, true),
    ( 'Embedded Systems', 'CS429', 'Introduction to embedded systems', 4, 8, 60, 20, current_timestamp, true),
    ( 'High-Performance Computing', 'CS430', 'Introduction to high-performance computing', 3, 6, 45, 15, current_timestamp, true),
    ( 'Distributed Systems', 'CS431', 'Introduction to distributed systems', 4, 8, 60, 20, current_timestamp, true),
    ( 'Computer Architecture', 'CS432', 'Study of computer architecture', 3, 6, 45, 15, current_timestamp, true),
    ( 'Linear Algebra for Computer Science', 'CS433', 'Introduction to linear algebra for computer science', 4, 8, 60, 20, current_timestamp, true),
    ( 'Discrete Mathematics for Computer Science', 'CS434', 'Introduction to discrete mathematics for computer science', 3, 6, 45, 15, current_timestamp, true),
    ( 'Calculus for Computer Science', 'CS435', 'Introduction to calculus for computer science', 4, 8, 60, 20, current_timestamp, true),
    ( 'Statistics for Computer Science', 'CS436', 'Introduction to statistics for computer science', 3, 6, 45, 15, current_timestamp, true),
    ( 'Ethics in Technology', 'CS437', 'Study of ethics in technology', 4, 8, 60, 20, current_timestamp, true),
    ( 'History of Computing', 'CS438', 'Introduction to the history of computing', 3, 6, 45, 15, current_timestamp, true),
    ( 'Philosophy of Computing', 'CS439', 'Introduction to the philosophy of computing', 4, 8, 60, 20, current_timestamp, true),
    ( 'Social Implications of Computing', 'CS440', 'Study of the social implications of computing', 3, 6, 45, 15, current_timestamp, true);


INSERT INTO curriculum_entity (department_id, created_at, is_active)
VALUES
    (1,current_timestamp , true),
    (1 ,current_timestamp , true),
    (1,current_timestamp , true),
    (1, current_timestamp , true),
    (1,current_timestamp , true),
    (1, current_timestamp , true),
    (1,current_timestamp , true),
    (1, current_timestamp , true),
    (1,current_timestamp , true),
    (1 ,current_timestamp , true);

