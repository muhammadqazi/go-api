INSERT INTO
    faculties_entity (faculty_id,name, created_at, is_active)
VALUES
    (1,'Engineering', current_timestamp, true),
    (2,'Business', current_timestamp, true),
    (3,'Arts and Sciences', current_timestamp, true),
    (4,'Education', current_timestamp, true),
    (5,'Medicine', current_timestamp, true);

INSERT INTO
    departments_entity (
        department_id,
        department_code,
        name,
        created_at,
        is_active,
        faculty_id
    )
VALUES
    (1,'CS','Computer Science',current_timestamp,true,1),
    (2,'MATH','Mathematics',current_timestamp,true,1),
    (3, 'PHYS', 'Physics', current_timestamp, true, 1),
    (4,'CHEM', 'Chemistry',current_timestamp,true,2),
    (5, 'BIO', 'Biology', current_timestamp, true, 2);

INSERT INTO courses_entity (course_id, course_code, name, credit_hours, quota, status , is_active,created_at,department_id)
VALUES 
(1, 'CS101', 'Introduction to Computer Science', 3, 30, 'active',true,current_timestamp,1),
(2, 'MATH101', 'Calculus I', 4, 25, 'active',true,current_timestamp,1),
(3, 'PHYS101', 'Classical Mechanics', 3, 20, 'active',true,current_timestamp,1),
(4, 'CHEM101', 'General Chemistry', 3, 25, 'active',true,current_timestamp,2),
(5, 'BIO101', 'Introduction to Biology', 3, 30, 'active',true,current_timestamp,2);

INSERT INTO accounts_entity
    (accounts_id, total_fee, current_dept,approaching_dept,total_dept,scholarship, discount, installments) 
VALUES
    (1, 10000, 70000,1000,8000, 1000, 500, 4),
    (2, 12000, 70000,1000,8000,1500, 400, 5),
    (3, 8000, 70000,1000,8000,500, 300, 3),
    (4, 15000, 70000,1000,8000,2000, 600, 4),
    (5, 7000, 70000,1000,8000,700, 200, 2);


INSERT INTO 
    invoices_entity (invoice_id, date, amount, description, installment, term,accounts_id) 
VALUES 
    (1, current_timestamp, 1000, 'Monthly invoice', 1, 'Fall',1),
    (2, current_timestamp, 1000, 'Monthly invoice', 1, 'Fall',2),
    (3, current_timestamp, 1000, 'Monthly invoice', 1, 'Fall',3),
    (4, current_timestamp, 1000, 'Monthly invoice', 1, 'Fall',4),
    (5, current_timestamp, 1000, 'Monthly invoice', 1, 'Fall',5);

INSERT INTO payments_entity 
    (payment_id, date, amount, process_type,accounts_id) 
VALUES 
    (1, current_timestamp, 1000, 'Cash',1),
    (2, current_timestamp, 1000, 'Check',2),
    (3, current_timestamp, 1000, 'Credit card',3),
    (4, current_timestamp, 1000, 'Bank transfer',4),
    (5, current_timestamp, 1000, 'Paypal',5);


INSERT INTO personal_info_entity (personal_info_id, id_card_number, passport_number, father_name, mother_name) 
VALUES 
    (1, 1234567890, 'A12345678', 'John Smith', 'Jane Smith'),
    (2, 1234567891, 'A12345679', 'Bob Smith', 'Samantha Smith'),
    (3, 1234567892, 'A12345680', 'Alice Smith', 'Bob Smith'),
    (4, 1234567893, 'A12345681', 'Mike Smith', 'Sara Smith'),
    (5, 1234567894, 'A12345682', 'Sarah Smith', 'Chris Smith');

INSERT INTO contact_info_entity (contact_info_id, email, phone_number, local_address, emergency_name, emergency_phone) 
VALUES 
    (1, 'john@example.com', '123-456-7890', '123 Main Street', 'Jane Smith', '123-456-7891'),
    (2, 'bob@example.com', '123-456-7892', '456 Main Street', 'Samantha Smith', '123-456-7893'),
    (3, 'alice@example.com', '123-456-7894', '789 Main Street', 'Bob Smith', '123-456-7895'),
    (4, 'mike@example.com', '123-456-7896', '246 Main Street', 'Sara Smith', '123-456-7897'),
    (5, 'sarah@example.com', '123-456-7898', '369 Main Street', 'Chris Smith', '123-456-7899');

INSERT INTO advisors_entity 
    (advisor_id, name, surname, email, password, office, line, course_id,is_active,created_at) 
VALUES 
    (1, 'John', 'Smith', 'john@example.com', 'password123', '123 Main Street', 'Computer Science', 1,true,current_timestamp),
    (2, 'Bob', 'Smith', 'bob@example.com', 'password456', '456 Main Street', 'Computer Science', 1,true,current_timestamp),
    (3, 'Alice', 'Smith', 'alice@example.com', 'password789', '789 Main Street', 'Computer Science', 2,true,current_timestamp),
    (4, 'Mike', 'Smith', 'mike@example.com', 'password246', '246 Main Street', 'Computer Science', 2,true,current_timestamp),
    (5, 'Sarah', 'Smith', 'sarah@example.com', 'password369', '369 Main Street', 'Computer Science', 3,true,current_timestamp);


INSERT INTO addresses_entity 
    (address_id, state, city, province, address) 
VALUES
    (1, 'New York', 'New York City', 'NY', '123 Main St'),
    (2, 'California', 'Los Angeles', 'CA', '456 Maple Ave'),
    (3, 'Texas', 'Houston', 'TX', '789 Oak St'),
    (4, 'Georgia', 'Atlanta', 'GA', '321 Pine St'),
    (5, 'Illinois', 'Chicago', 'IL', '654 Cedar Ave');

INSERT INTO students_entity 
    (created_at, updated_at, deleted_at, is_active, first_name, surname, email, nationality, dob, place_of_birth, sex, password, role, status, semester, enrollment_date, graduation_date,student_id,is_deleted,is_graduated,faculty_id,personal_info_id,contact_info_id,address_id,accounts_id) 
VALUES 
    (current_timestamp, current_timestamp, NULL, true, 'John', 'Doe', 'john.doe@example.com', 'American', '01/01/1970', 'New York', 'male', 'password', 'student', 'enrolled', 'Spring', '01/01/2021', NULL,21906778,false,false,1,1,1,1,1), 
    (current_timestamp, current_timestamp, NULL, true, 'Jane', 'Doe', 'jane.doe@example.com', 'Canadian', '01/01/1980', 'Toronto', 'female', 'password', 'student', 'enrolled', 'Fall', '01/01/2021', NULL,22107446,false,false,1,1,1,1,1)
