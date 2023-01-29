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
    (0, 0, 10, 'Sports', 2, 50, 21906778, 1575, 3500),
    (0, 0, 10, 'Sports', 2, 50, 21906778, 1575, 3500);

INSERT INTO instructors_entity (instructor_id, first_name, last_name, email, phone_number, password, dob, place_of_birth, sex, nationality, role , created_at,is_active)
VALUES
    (1, 'John', 'Doe', 'johndoe@example.com', '555-555-5555', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1980-01-01', 'New York', 'Male', 'American', 'Instructor',current_timestamp,true),
    (2, 'Jane', 'Doe', 'janedoe@example.com', '555-555-5556', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1981-02-02', 'Los Angeles', 'Female', 'American', 'Instructor',current_timestamp,true),
    (3, 'Bob', 'Smith', 'bobsmith@example.com', '555-555-5557', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1982-03-03', 'Chicago', 'Male', 'American', 'Instructor',current_timestamp,true),
    (4, 'Alice', 'Johnson', 'alicejohnson@example.com', '555-555-5558', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1983-04-04', 'Houston', 'Female', 'American', 'Instructor',current_timestamp,true),
    (5, 'Tom', 'Williams', 'tomwilliams@example.com', '555-555-5559', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1984-05-05', 'Philadelphia', 'Male', 'American', 'Instructor',current_timestamp,true),
    (6, 'Emily', 'Jones', 'emilyjones@example.com', '555-555-5560', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1985-06-06', 'Phoenix', 'Female', 'American', 'Instructor',current_timestamp,true),
    (7, 'David', 'Brown', 'davidbrown@example.com', '555-555-5561', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1986-07-07', 'San Antonio', 'Male', 'American', 'Instructor',current_timestamp,true),
    (8, 'Sophie', 'Davis', 'sophiedavis@example.com', '555-555-5562', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1987-08-08', 'San Diego', 'Female', 'American', 'Instructor',current_timestamp,true),
    (9, 'Jacob', 'Miller', 'jacobmiller@example.com', '555-555-5563', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1988-09-09', 'Dallas', 'Male', 'American', 'Instructor',current_timestamp,true),
    (10, 'Olivia', 'Wilson', 'oliviawilson@example.com', '555-555-5564', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim0', '1989-10-10', 'San Jose', 'Female', 'American', 'Instructor',current_timestamp,true);




INSERT INTO courses_entity (course_id, name, code, description, credits, ects, theoretical, practical, created_at, is_active)
VALUES
    (1, 'Introduction to Computer Science', 'CS101', 'An overview of computer science', 3, 6, 45, 15, current_timestamp, true),
    (2, 'Programming Fundamentals', 'CS102', 'Introduction to programming concepts', 4, 8, 60, 20, current_timestamp, true),
    (3, 'Data Structures and Algorithms', 'CS203', 'Study of data structures and algorithms', 4, 8, 60, 20, current_timestamp, true),
    (4, 'Web Development', 'CS304', 'Introduction to web development', 4, 8, 60, 20, current_timestamp, true),
    (5, 'Database Systems', 'CS405', 'Introduction to database systems', 3, 6, 45, 15, current_timestamp, true),
    (6, 'Operating Systems', 'CS406', 'Study of operating systems', 4, 8, 60, 20, current_timestamp, true),
    (7, 'Computer Networks', 'CS407', 'Introduction to computer networks', 3, 6, 45, 15, current_timestamp, true),
    (8, 'Object-Oriented Programming', 'CS408', 'Introduction to object-oriented programming', 4, 8, 60, 20, current_timestamp, true),
    (9, 'Software Engineering', 'CS409', 'Study of software engineering concepts', 4, 8, 60, 20, current_timestamp, true),
    (10, 'Artificial Intelligence', 'CS410', 'Introduction to artificial intelligence', 3, 6, 45, 15, current_timestamp, true),
    (11, 'Machine Learning', 'CS411', 'Study of machine learning algorithms', 4, 8, 60, 20, current_timestamp, true),
    (12, 'Data Science', 'CS412', 'Introduction to data science', 3, 6, 45, 15, current_timestamp, true),
    (13, 'Cloud Computing', 'CS413', 'Introduction to cloud computing', 4, 8, 60, 20, current_timestamp, true),
    (14, 'Cybersecurity', 'CS414', 'Introduction to cybersecurity', 3, 6, 45, 15, current_timestamp, true),
    (15, 'Blockchain Technology', 'CS415', 'Introduction to blockchain technology', 4, 8, 60, 20, current_timestamp, true),
    (16, 'Virtual Reality', 'CS416', 'Introduction to virtual reality', 3, 6, 45, 15, current_timestamp, true),
    (17, 'Augmented Reality', 'CS417', 'Introduction to augmented reality', 4, 8, 60, 20, current_timestamp, true),
    (18, 'Internet of Things', 'CS418', 'Introduction to internet of things', 3, 6, 45, 15, current_timestamp, true),
    (19, 'Mobile Application Development', 'CS419', 'Introduction to mobile application development', 4, 8, 60, 20, current_timestamp, true),
    (20, 'Quantum Computing', 'CS420', 'Introduction to quantum computing', 3, 6, 45, 15, current_timestamp,true),
    (21, 'Human-Computer Interaction', 'CS421', 'Introduction to human-computer interaction', 4, 8, 60, 20, current_timestamp, true),
    (22, 'Computer Graphics', 'CS422', 'Introduction to computer graphics', 3, 6, 45, 15, current_timestamp, true),
    (23, 'Compiler Design', 'CS423', 'Study of compiler design', 4, 8, 60, 20, current_timestamp, true),
    (24, 'Parallel Computing', 'CS424', 'Introduction to parallel computing', 3, 6, 45, 15, current_timestamp, true),
    (25, 'Computer Vision', 'CS425', 'Introduction to computer vision', 4, 8, 60, 20, current_timestamp, true),
    (26, 'Natural Language Processing', 'CS426', 'Introduction to natural language processing', 3, 6, 45, 15, current_timestamp, true),
    (27, 'Information Retrieval', 'CS427', 'Introduction to information retrieval', 4, 8, 60, 20, current_timestamp, true),
    (28, 'Cryptography', 'CS428', 'Introduction to cryptography', 3, 6, 45, 15, current_timestamp, true),
    (29, 'Embedded Systems', 'CS429', 'Introduction to embedded systems', 4, 8, 60, 20, current_timestamp, true),
    (30, 'High-Performance Computing', 'CS430', 'Introduction to high-performance computing', 3, 6, 45, 15, current_timestamp, true),
    (31, 'Distributed Systems', 'CS431', 'Introduction to distributed systems', 4, 8, 60, 20, current_timestamp, true),
    (32, 'Computer Architecture', 'CS432', 'Study of computer architecture', 3, 6, 45, 15, current_timestamp, true),
    (33, 'Linear Algebra for Computer Science', 'CS433', 'Introduction to linear algebra for computer science', 4, 8, 60, 20, current_timestamp, true),
    (34, 'Discrete Mathematics for Computer Science', 'CS434', 'Introduction to discrete mathematics for computer science', 3, 6, 45, 15, current_timestamp, true),
    (35, 'Calculus for Computer Science', 'CS435', 'Introduction to calculus for computer science', 4, 8, 60, 20, current_timestamp, true),
    (36, 'Statistics for Computer Science', 'CS436', 'Introduction to statistics for computer science', 3, 6, 45, 15, current_timestamp, true),
    (37, 'Ethics in Technology', 'CS437', 'Study of ethics in technology', 4, 8, 60, 20, current_timestamp, true),
    (38, 'History of Computing', 'CS438', 'Introduction to the history of computing', 3, 6, 45, 15, current_timestamp, true),
    (39, 'Philosophy of Computing', 'CS439', 'Introduction to the philosophy of computing', 4, 8, 60, 20, current_timestamp, true),
    (40, 'Social Implications of Computing', 'CS440', 'Study of the social implications of computing', 3, 6, 45, 15, current_timestamp, true);


INSERT INTO curriculum_entity (department_id, course_id, semester, year , created_at, is_active)
VALUES
    (1, 1, 'Fall', 2019,current_timestamp , true), (1, 2, 'Fall', 2019,current_timestamp , true), (1, 3, 'Fall', 2019,current_timestamp , true), (1, 4, 'Fall', 2019,current_timestamp , true), (1, 5, 'Fall', 2019,current_timestamp , true),
    (1, 6, 'Spring', 2019,current_timestamp , true), (1, 7, 'Spring', 2019,current_timestamp , true), (1, 8, 'Spring', 2019,current_timestamp , true), (1, 9, 'Spring', 2019,current_timestamp , true), (1, 10, 'Spring', 2019,current_timestamp , true),
    (1, 11, 'Fall', 2020,current_timestamp , true), (1, 12, 'Fall', 2020,current_timestamp , true), (1, 13, 'Fall', 2020,current_timestamp , true), (1, 14, 'Fall', 2020,current_timestamp , true), (1, 15, 'Fall', 2020,current_timestamp , true),
    (1, 16, 'Spring', 2020,current_timestamp , true), (1, 17, 'Spring', 2020,current_timestamp , true), (1, 18, 'Spring', 2020,current_timestamp , true), (1, 19, 'Spring', 2020,current_timestamp , true), (1, 20, 'Spring', 2020,current_timestamp , true),
    (1, 21, 'Fall', 2021,current_timestamp , true), (1, 22, 'Fall', 2021,current_timestamp , true), (1, 23, 'Fall', 2021,current_timestamp , true), (1, 24, 'Fall', 2021,current_timestamp , true), (1, 25, 'Fall', 2021,current_timestamp , true),
    (1, 26, 'Spring', 2021,current_timestamp , true), (1, 27, 'Spring', 2021,current_timestamp , true), (1, 28, 'Spring', 2021,current_timestamp , true), (1, 29, 'Spring', 2021,current_timestamp , true), (1, 30, 'Spring', 2021,current_timestamp , true),
    (1, 31, 'Fall', 2022,current_timestamp , true), (1, 32, 'Fall', 2022,current_timestamp , true), (1, 33, 'Fall', 2022,current_timestamp , true), (1, 34, 'Fall', 2022,current_timestamp , true), (1, 35, 'Fall', 2022,current_timestamp , true),
    (1, 36, 'Spring', 2022,current_timestamp , true), (1, 37, 'Spring', 2022,current_timestamp , true), (1, 38, 'Spring', 2022,current_timestamp , true), (1, 39, 'Spring', 2022,current_timestamp , true), (1, 40, 'Spring', 2022,current_timestamp,true);
