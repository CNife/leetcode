import test.Tester;

import java.util.*;
import java.util.stream.Collectors;

public class EmployeeImportance {
    public static int getImportance(List<Employee> employees, int id) {
        Map<Integer, Employee> map = new HashMap<>();
        for (Employee employee : employees) {
            map.put(employee.id, employee);
        }

        Deque<Employee> queue = new ArrayDeque<>();
        queue.addLast(map.get(id));
        int result = 0;
        while (!queue.isEmpty()) {
            Employee employee = queue.removeFirst();
            result += employee.importance;
            for (Integer subordinate : employee.subordinates) {
                queue.addLast(map.get(subordinate));
            }
        }
        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(EmployeeImportance.class);
        tester.addTestCase(
                Arrays.asList(
                        new Employee(1, 5, 2, 3),
                        new Employee(2, 3),
                        new Employee(3, 3)
                ), 1,
                11
        );
        tester.runTestCases();
    }

    private static class Employee {
        int id;
        int importance;
        List<Integer> subordinates;

        Employee(int id, int importance, int... subordinates) {
            this.id = id;
            this.importance = importance;
            this.subordinates = Arrays.stream(subordinates).boxed().collect(Collectors.toList());
        }

        @Override
        public String toString() {
            return "Employee{" +
                    "id=" + id +
                    ", importance=" + importance +
                    ", subordinates=" + subordinates +
                    '}';
        }
    }
}
