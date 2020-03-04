class TestCase<E, S> {
    E expected;
    S source;

    TestCase(E expected, S source) {
        this.expected = expected;
        this.source = source;
    }
}
