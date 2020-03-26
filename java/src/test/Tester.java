package test;

import java.lang.reflect.Method;
import java.lang.reflect.Modifier;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Objects;
import java.util.function.Function;

public class Tester {
    private List<TestCase> testCases;
    private Method testMethod;
    private Function<Object, Object> transform;

    public Tester(Class<?> testClass) {
        this.testMethod = getTestMethod(testClass);
        testCases = new ArrayList<>();
    }

    public Tester(Class<?> testClass, Function<Object, Object> transform) {
        this(testClass);
        this.transform = transform;
    }

    private static Method getTestMethod(Class<?> testClass) {
        for (Method method : testClass.getDeclaredMethods()) {
            int modifiers = method.getModifiers();
            if ((modifiers & Modifier.PUBLIC) != 0 // public
                    && (modifiers & Modifier.STATIC) != 0 // static
                    && (!Objects.equals(method.getName(), "main"))) { // not main
                return method;
            }
        }
        throw new IllegalStateException("找不到测试方法");
    }

    private static boolean isInstance(Class<?> clazz, Object instance) {
        switch (clazz.getCanonicalName()) {
            case "boolean":
                return instance instanceof Boolean;
            case "byte":
                return instance instanceof Byte;
            case "char":
                return instance instanceof Character;
            case "short":
                return instance instanceof Short;
            case "int":
                return instance instanceof Integer;
            case "long":
                return instance instanceof Long;
            case "float":
                return instance instanceof Float;
            case "double":
                return instance instanceof Double;
            default:
                return instance == null || clazz.isInstance(instance);
        }
    }

    private static boolean smartEquals(Object lhs, Object rhs) {
        if (lhs == rhs) {
            return true;
        } else if (lhs == null | rhs == null) {
            return false;
        } else if (lhs.getClass().isArray()) {
            switch (lhs.getClass().getComponentType().getCanonicalName()) {
                case "boolean":
                    return Arrays.equals((boolean[]) lhs, (boolean[]) rhs);
                case "byte":
                    return Arrays.equals((byte[]) lhs, (byte[]) rhs);
                case "char":
                    return Arrays.equals((char[]) lhs, (char[]) rhs);
                case "short":
                    return Arrays.equals((short[]) lhs, (short[]) rhs);
                case "int":
                    return Arrays.equals((int[]) lhs, (int[]) rhs);
                case "long":
                    return Arrays.equals((long[]) lhs, (long[]) rhs);
                case "float":
                    return Arrays.equals((float[]) lhs, (float[]) rhs);
                case "double":
                    return Arrays.equals((double[]) lhs, (double[]) rhs);
                default:
                    return Arrays.deepEquals((Object[]) lhs, (Object[]) rhs);
            }
        } else {
            return lhs.equals(rhs);
        }
    }

    public void addTestCase(Object... args) {
        checkTestCase(args);

        Object expect = args[args.length - 1];
        Object[] source = Arrays.copyOf(args, args.length - 1);
        TestCase testCase = new TestCase(source, expect);
        testCases.add(testCase);
    }

    public void runTestCases() {
        try {
            for (TestCase testCase : testCases) {
                Object expect = testCase.expect;
                Object actual = testMethod.invoke(null, testCase.source);
                if (transform != null) {
                    expect = transform.apply(expect);
                    actual = transform.apply(actual);
                }

                if (!smartEquals(expect, actual)) {
                    System.err.println("测试失败：");
                    System.err.println("测例：" + Arrays.toString(testCase.source));
                    System.err.println("实际结果：" + actual);
                    System.err.println("期望结果：" + expect);
                    System.exit(1);
                }
            }
            System.out.println("通过测试：" + testMethod.getDeclaringClass().getName());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private void checkTestCase(Object... args) {
        Class<?>[] paramTypes = testMethod.getParameterTypes();
        Class<?> returnType = testMethod.getReturnType();

        int expectArgsLength = paramTypes.length;
        if (returnType != Void.class) {
            expectArgsLength += 1;
        }
        if (expectArgsLength != args.length) {
            throw new IllegalArgumentException("测例参数数量错误：" +
                    "需要 " + expectArgsLength + "，" +
                    "实际 " + args.length);
        }

        for (int i = 0; i < paramTypes.length - 1; i++) {
            if (!isInstance(paramTypes[i], args[i])) {
                throw new IllegalArgumentException("第" + (i + 1) + "个参数类型错误：" +
                        "需要 " + paramTypes[i].getCanonicalName() + "，" +
                        "实际 " + args[i].getClass().getCanonicalName());
            }
        }
        if (returnType != Void.class) {
            if (!isInstance(returnType, args[args.length - 1])) {
                throw new IllegalArgumentException("返回类型错误：" +
                        "需要 " + returnType + "，" +
                        "实际 " + args[args.length - 1]);
            }
        }
    }
}


