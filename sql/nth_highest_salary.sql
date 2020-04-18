CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT BEGIN RETURN (
    SELECT
        DISTINCT e.salary
    FROM
        employee e
    WHERE
        (
            SELECT
                count(DISTINCT salary)
            FROM
                employee
            WHERE
                salary > e.salary
        ) = N - 1
);

END