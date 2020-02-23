public class TestCase<E, T> {
    E expected;
    T testSource;

    public TestCase(E expected, T testSource) {
        this.expected = expected;
        this.testSource = testSource;
    }
}
