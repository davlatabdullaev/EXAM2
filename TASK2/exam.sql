-- TASK 2
SELECT
    d.name AS department,
    e.name AS employee,
    COUNT(s.employee_id) AS count
FROM
    departments d
JOIN
    employees e ON d.id = e.department_id
LEFT JOIN
    salaries s ON e.id = s.employee_id
GROUP BY
    d.name, e.name;
