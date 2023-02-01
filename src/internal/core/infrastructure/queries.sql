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
