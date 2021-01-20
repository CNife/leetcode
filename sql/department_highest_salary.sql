SELECT
    Department.Name AS Department,
    Employee.Name AS Employee,
    Employee.Salary AS Salary
FROM
    (
        SELECT
            DepartmentId,
            Max(Salary) AS MaxSalary
        FROM
            Employee
        GROUP BY
            DepartmentId
    ) MaxSalaryInDepartment
    INNER JOIN Employee ON MaxSalaryInDepartment.DepartmentId = Employee.DepartmentId
    INNER JOIN Department ON MaxSalaryInDepartment.DepartmentId = Department.Id
WHERE
    MaxSalaryInDepartment.MaxSalary = Employee.Salary;