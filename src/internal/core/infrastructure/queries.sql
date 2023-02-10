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
       cc.curriculum_id, cc.year,cc.semester, cc.department_id,
       d.name AS department_name, d.department_code , d.number_of_years,
       co.course_id, co.code, co.name,co.credits,co.ects,co.practical,co.theoretical
FROM course_curriculum_entity sc
         JOIN curriculum_entity cc ON sc.curriculum_id = cc.curriculum_id
         JOIN departments_entity d ON cc.department_id = d.department_id
         JOIN courses_entity co ON sc.course_id = co.course_id
WHERE cc.department_id=2 AND sc.is_active=true;
