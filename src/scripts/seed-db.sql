INSERT INTO
    faculties_entity (name, created_at, is_active)
VALUES
    ('Engineering', current_timestamp, true),
    ('Business', current_timestamp, true),
    ('Arts and Sciences', current_timestamp, true),
    ('Education', current_timestamp, true),
    ('Medicine', current_timestamp, true);

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
    (
        1,
        'CS',
        'Computer Science',
        current_timestamp,
        true,
        1
    ),
    (
        2,
        'MATH',
        'Mathematics',
        current_timestamp,
        true,
        1
    ),
    (3, 'PHYS', 'Physics', current_timestamp, true, 1),
    (
        4,
        'CHEM',
        'Chemistry',
        current_timestamp,
        true,
        2
    ),
    (5, 'BIO', 'Biology', current_timestamp, true, 2);