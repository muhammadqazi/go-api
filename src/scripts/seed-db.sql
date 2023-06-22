INSERT INTO buildings_entity (building_id, name, code, description, number_of_rooms, created_at)
VALUES
    (1, 'Building 1', 'B1', 'Description 1', 10, current_timestamp),
    (2, 'Building 2', 'B2', 'Description 2', 15,current_timestamp),
    (3, 'Building 3', 'B3', 'Description 3', 8,current_timestamp),
    (4, 'Building 4', 'B4', 'Description 4', 12,current_timestamp),
    (5, 'Building 5', 'B5', 'Description 5', 20,current_timestamp);

INSERT INTO building_rooms_entity (building_room_id, room_number, room_description, number_of_seats, is_available, building_id, created_at)
VALUES
    (1, 'Room 1', 'Description 1', 10, true, 1,current_timestamp),
    (2, 'Room 2', 'Description 2', 15, true, 1,current_timestamp),
    (3, 'Room 3', 'Description 3', 8, false, 2,current_timestamp),
    (4, 'Room 4', 'Description 4', 12, true, 2,current_timestamp),
    (5, 'Room 5', 'Description 5', 20, true, 3,current_timestamp);


INSERT INTO instructors_entity (first_name, last_name, email, phone_number, password, dob, place_of_birth, sex, nationality, role , created_at,is_active, salary,office_id)
VALUES
    ('John', 'Doe', 'johndoe@example.com', '555-555-5555', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1980-01-01', 'New York', 'Male', 'American', 'Instructor',current_timestamp,true,900,1),
    ('Jane', 'Doe', 'janedoe@example.com', '555-555-5556', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1981-02-02', 'Los Angeles', 'Female', 'American', 'Instructor',current_timestamp,true,900,1),
    ('Bob', 'Smith', 'bobsmith@example.com', '555-555-5557', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1982-03-03', 'Chicago', 'Male', 'American', 'Instructor',current_timestamp,true,900,1),
    ('Alice', 'Johnson', 'alicejohnson@example.com', '555-555-5558', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1983-04-04', 'Houston', 'Female', 'American', 'Instructor',current_timestamp,true,900,1),
    ('Tom', 'Williams', 'tomwilliams@example.com', '555-555-5559', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1984-05-05', 'Philadelphia', 'Male', 'American', 'Instructor',current_timestamp,true,900,1),
    ('Emily', 'Jones', 'emilyjones@example.com', '555-555-5560', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1985-06-06', 'Phoenix', 'Female', 'American', 'Instructor',current_timestamp,true,900,1),
    ('David', 'Brown', 'davidbrown@example.com', '555-555-5561', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1986-07-07', 'San Antonio', 'Male', 'American', 'Instructor',current_timestamp,true,900,1),
    ('Sophie', 'Davis', 'sophiedavis@example.com', '555-555-5562', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1987-08-08', 'San Diego', 'Female', 'American', 'instructor',current_timestamp,true,900,1),
    ('Jacob', 'Miller', 'oliviawilson@example.com2', '555-555-5563', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', '1988-09-09', 'Dallas', 'Male', 'American', 'admin',current_timestamp,true,900,1),
    ( 'Olivia', 'Wilson', 'oliviawilson@example.com', '555-555-5564', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim0', '1989-10-10', 'San Jose', 'Female', 'American', 'admin',current_timestamp,true,900,1);

INSERT INTO faculties_entity (name, description, email, phone_number , created_at, is_active , dean_id,vice_dean_id)
VALUES
    ('Mathematics','Department of Mathematics','math@example.com', '555-555-5555',current_timestamp , true,1 , 2),
    ('Computer Science','Department of Computer Science','cs@example.com', '555-555-5557',current_timestamp , true,1,2 ),
    ('Physics','Department of Physics', 'physics@example.com', '555-555-5559',current_timestamp , true,2,1),
    ('Chemistry','Department of Chemistry','chemistry@example.com', '555-555-5561',current_timestamp , true,2,2),
    ('Biology','Department of Biology', 'biology@example.com', '555-555-5563',current_timestamp , true,2,2);

INSERT INTO departments_entity (department_code, name, description, email, phone_number, created_at, is_active , number_of_years, faculty_id)
VALUES
    ('D001', 'Computer Science', 'Study of computers and computational systems', 'cs@university.com', '555-555-5555', current_timestamp , true,4 , 1),
    ('D002', 'Mathematics', 'Study of numbers and shapes', 'math@university.com', '555-555-5556',current_timestamp , true,4,2),
    ('D003', 'Physics', 'Study of matter and energy', 'physics@university.com', '555-555-5557',current_timestamp , true,4,3),
    ('D004', 'Chemistry', 'Study of chemicals and their reactions',  'chemistry@university.com', '555-555-5558',current_timestamp , true,4,1),
    ('D005', 'Biology', 'Study of living organisms', 'biology@university.com', '555-555-5559',current_timestamp , true,4,1);

INSERT INTO students_entity (student_id, first_name, surname, email, nationality, dob, place_of_birth, sex, password, role, status, access_status, acceptance_type, semester, graduation_date, is_graduated, department_id, supervisor_id , created_at, is_active)
VALUES
    (21906778, 'John', 'Doe', 'johndoe@example.com', 'American', '01-01-1998', 'New York', 'Male', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', 'student', 'registered', 'active', 'Admitted', 'Fall 2021', '2025-05-25', false, 1, 10, current_timestamp, true),
    (21906779, 'John', 'Doe', 'johndoe2@example.com', 'American', '01-01-1998', 'New York', 'Male', '$2a$14$Lo3A2ZYhrl3oHxAKqbgif.RaPFkm77XIOBiU05veQuEXl9T5xfXim', 'student', 'registered', 'active', 'Admitted', 'Fall 2021', '2025-05-25', false, 1, 10,current_timestamp, true);

INSERT INTO accounts_entity (approaching_dept, current_dept, discount, discount_type, installments, scholarship, student_id, total_dept, total_fee,department_fee)
VALUES
    (0, 0, 10, 'Sports', 2, 50, 21906778, 1575, 1750 , 3500),
    (0, 0, 10, 'Sports', 2, 50, 21906778, 1575, 1600 , 3500);


INSERT INTO invoices_entity (invoice_id, date, amount, description, installment, term, account_id)
VALUES
    (1, current_timestamp, 100.0, 'Invoice 1', 1, 'Term 1', 1),
    (2, current_timestamp, 150.0, 'Invoice 2', 2, 'Term 1', 1),
    (3, current_timestamp, 200.0, 'Invoice 3', 3, 'Term 2', 1),
    (4, current_timestamp, 250.0, 'Invoice 4', 4, 'Term 2', 1),
    (5, current_timestamp, 300.0, 'Invoice 5', 5, 'Term 3', 1);

INSERT INTO courses_entity (name, code, description, credits, ects, theoretical, practical, created_at, is_active, department_id)
VALUES
    ('Introduction to Computer Science', 'CS101', 'An overview of computer science', 3, 6, 45, 15, current_timestamp, true, 1),
    ('Programming Fundamentals', 'CS102', 'Introduction to programming concepts', 4, 8, 60, 20, current_timestamp, true, 1),
    ('Data Structures and Algorithms', 'CS203', 'Study of data structures and algorithms', 4, 8, 60, 20, current_timestamp, true, 1),
    ('Web Development', 'CS304', 'Introduction to web development', 4, 8, 60, 20, current_timestamp, true, 1),
    ('Database Systems', 'CS405', 'Introduction to database systems', 3, 6, 45, 15, current_timestamp, true, 1),
    ('Operating Systems', 'CS406', 'Study of operating systems', 4, 8, 60, 20, current_timestamp, true, 1),
    ('Computer Networks', 'CS407', 'Introduction to computer networks', 3, 6, 45, 15, current_timestamp, true, 1),
    ('Object-Oriented Programming', 'CS408', 'Introduction to object-oriented programming', 4, 8, 60, 20, current_timestamp, true, 1),
    ('Software Engineering', 'CS409', 'Study of software engineering concepts', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Artificial Intelligence', 'CS410', 'Introduction to artificial intelligence', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Machine Learning', 'CS411', 'Study of machine learning algorithms', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Data Science', 'CS412', 'Introduction to data science', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Cloud Computing', 'CS413', 'Introduction to cloud computing', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Cybersecurity', 'CS414', 'Introduction to cybersecurity', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Blockchain Technology', 'CS415', 'Introduction to blockchain technology', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Virtual Reality', 'CS416', 'Introduction to virtual reality', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Augmented Reality', 'CS417', 'Introduction to augmented reality', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Internet of Things', 'CS418', 'Introduction to internet of things', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Mobile Application Development', 'CS419', 'Introduction to mobile application development', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Quantum Computing', 'CS420', 'Introduction to quantum computing', 3, 6, 45, 15, current_timestamp,true, 1),
    ( 'Human-Computer Interaction', 'CS421', 'Introduction to human-computer interaction', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Computer Graphics', 'CS422', 'Introduction to computer graphics', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Compiler Design', 'CS423', 'Study of compiler design', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Parallel Computing', 'CS424', 'Introduction to parallel computing', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Computer Vision', 'CS425', 'Introduction to computer vision', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Natural Language Processing', 'CS426', 'Introduction to natural language processing', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Information Retrieval', 'CS427', 'Introduction to information retrieval', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Cryptography', 'CS428', 'Introduction to cryptography', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Embedded Systems', 'CS429', 'Introduction to embedded systems', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'High-Performance Computing', 'CS430', 'Introduction to high-performance computing', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Distributed Systems', 'CS431', 'Introduction to distributed systems', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Computer Architecture', 'CS432', 'Study of computer architecture', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Linear Algebra for Computer Science', 'CS433', 'Introduction to linear algebra for computer science', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Discrete Mathematics for Computer Science', 'CS434', 'Introduction to discrete mathematics for computer science', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Calculus for Computer Science', 'CS435', 'Introduction to calculus for computer science', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Statistics for Computer Science', 'CS436', 'Introduction to statistics for computer science', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Ethics in Technology', 'CS437', 'Study of ethics in technology', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'History of Computing', 'CS438', 'Introduction to the history of computing', 3, 6, 45, 15, current_timestamp, true, 1),
    ( 'Philosophy of Computing', 'CS439', 'Introduction to the philosophy of computing', 4, 8, 60, 20, current_timestamp, true, 1),
    ( 'Social Implications of Computing', 'CS440', 'Study of the social implications of computing', 3, 6, 45, 15, current_timestamp, true, 1);

INSERT INTO course_schedule_entity (created_at, day, start_time, end_time, lecture_venue, course_id,is_theoretical)
VALUES (current_timestamp, 'Monday', '09:00', '11:00', 'Room 101', 6, false),
       (current_timestamp, 'Tuesday', '13:00', '15:00', 'Room 102', 7,false),
       (current_timestamp, 'Wednesday', '10:00', '12:00', 'Room 103', 8,true),
       (current_timestamp, 'Thursday', '11:00', '13:00', 'Room 104', 9,true),
       (current_timestamp, 'Friday', '14:00', '16:00', 'Room 105', 10,false);

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




