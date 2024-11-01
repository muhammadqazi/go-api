-- In SQL, there are three types of join operations: inner join, left join, and right join.
-- A left join returns all the rows from the left table (courses_entity in your case), and the matching rows from the right table (course_curriculum_entity). If there is no match, the result will contain NULL values for the right table columns.
-- A right join is similar to a left join, but the roles of the tables are reversed. The right join returns all the rows from the right table (course_curriculum_entity), and the matching rows from the left table (courses_entity). If there is no match, the result will contain NULL values for the left table columns.
-- Here's an example of a left join and a right join between the two tables:

-- Inner join example
SELECT courses_entity.name, courses_entity.code, courses_entity.description, courses_entity.credits, courses_entity.ects, courses_entity.theoretical, courses_entity.practical, course_curriculum_entity.curriculum_id
FROM courses_entity
INNER JOIN course_curriculum_entity ON courses_entity.course_id = course_curriculum_entity.course_id;


-- Left join example
SELECT courses_entity.name, courses_entity.code, courses_entity.description, courses_entity.credits, courses_entity.ects, courses_entity.theoretical, courses_entity.practical, course_curriculum_entity.curriculum_id
FROM courses_entity
LEFT JOIN course_curriculum_entity ON courses_entity.course_id = course_curriculum_entity.course_id;

-- Right join example
SELECT courses_entity.name, courses_entity.code, courses_entity.description, courses_entity.credits, courses_entity.ects, courses_entity.theoretical, courses_entity.practical, course_curriculum_entity.curriculum_id
FROM courses_entity
RIGHT JOIN course_curriculum_entity ON courses_entity.course_id = course_curriculum_entity.course_id;


SELECT sc.course_id, sc.course_load, sc.created_at, sc.updated_at,sc.deleted_at,
       cc.curriculum_id, sc.year,sc.semester, cc.department_id, sc.course_load,
       d.name AS department_name, d.department_code , d.number_of_years,
       co.course_id, co.code, co.name,co.credits,co.ects,co.practical,co.theoretical
FROM course_curriculum_entity sc
         JOIN curriculum_entity cc ON sc.curriculum_id = cc.curriculum_id
         JOIN departments_entity d ON cc.department_id = d.department_id
         JOIN courses_entity co ON sc.course_id = co.course_id
WHERE cc.department_id=2 AND sc.is_active=true;

SELECT en.student_enrollment_id AS enrollment_id,
       ins.last_name AS supervisor_name, ins.last_name AS supervisor_surname, ins.instructor_id AS supervisor_id,
       en.created_at,en.updated_at,en.deleted_at,en.is_approved,en.semester,en.year,en.student_id,
       std.first_name AS student_name,std.surname AS student_surname, std.status AS student_status, std.access_status,
       req.course_id, co.name AS course_name,co.code AS course_code,co.credits AS course_credits,co.is_active AS course_status
FROM student_course_request_entity req
         JOIN student_enrollments_entity en ON req.student_enrollment_id = en.student_enrollment_id
         JOIN instructors_entity ins ON ins.instructor_id = en.instructor_id
         JOIN students_entity std ON std.student_id = en.student_id
         JOIN courses_entity co ON req.course_id = co.course_id
WHERE en.is_active=true AND en.is_approved=false AND en.instructor_id = 10;


SELECT en.student_enrollment_id, en.student_id , req.course_id , req.student_course_request_id,
       co.name,co.code , sch.day , sch.start_time , sch.end_time , co.credits , sch.lecture_venue, en.year,en.semester
FROM student_enrollments_entity en
         JOIN student_course_request_entity req ON en.student_enrollment_id = req.student_enrollment_id
         JOIN courses_entity co ON req.course_id = co.course_id
         JOIN course_schedule_entity sch ON req.course_id = sch.course_id
WHERE en.student_id = 21906778 AND en.semester = 'spring' AND en.year = 2023 AND req.is_approved;


SELECT ex.created_at , ex.is_active , ex.exam_venue , ex.date ,  ex.exam_type, ex.duration , req.course_id , co.code , co.name , co.credits
FROM student_course_request_entity req
         JOIN student_enrollments_entity en ON en.student_enrollment_id = req.student_enrollment_id
         JOIN exam_schedule_entity ex ON ex.course_id = req.course_id
         JOIN courses_entity co ON co.course_id = req.course_id
WHERE en.student_id = 21906778 AND en.semester = 'spring' AND en.year = 2023 AND req.is_approved;

SELECT
    en.created_at AS enrollment_date , en.is_active , en.instructor_id , ins.first_name , ins.last_name , ins.email , ins.is_active AS instructor_status,
    co.course_id , co.name , co.code , co.credits , co.theoretical , co.practical
FROM instructor_enrollments_entity en
         JOIN instructor_courses_entity inco ON en.instructor_enrollment_id = inco.instructor_enrollment_id
         JOIN courses_entity co ON inco.course_id = co.course_id
         JOIN instructors_entity ins ON ins.instructor_id = en.instructor_id
WHERE en.instructor_id=9;