import test.Tester;

import java.util.*;

public class WordLadder2 {
    public static List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        HashSet<String> wordSet = new HashSet<>(wordList);
        if (wordSet.isEmpty() || !wordSet.contains(endWord)) {
            return Collections.emptyList();
        }

        Map<String, Set<String>> graph = new HashMap<>();
        boolean found = bfs(beginWord, endWord, wordSet, graph);
        if (!found) {
            return Collections.emptyList();
        }

        List<String> path = new ArrayList<>();
        path.add(beginWord);
        List<List<String>> result = new ArrayList<>();
        dfs(beginWord, endWord, graph, path, result);
        return result;
    }

    private static boolean bfs(String beginWord, String endWord,
                               HashSet<String> wordSet,
                               Map<String, Set<String>> graph
    ) {
        Queue<String> queue = new ArrayDeque<>();
        queue.offer(beginWord);
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        Set<String> nextLevelVisited = new HashSet<>();
        int wordLength = beginWord.length();

        boolean found = false;
        while (!queue.isEmpty()) {
            int queueLength = queue.size();
            for (int i = 0; i < queueLength; i++) {
                String currentWord = queue.remove();
                char[] currentWordArray = currentWord.toCharArray();
                for (int index = 0; index < wordLength; index++) {
                    char originChar = currentWordArray[index];
                    for (char ch = 'a'; ch <= 'z'; ch++) {
                        currentWordArray[index] = ch;
                        String nextWord = new String(currentWordArray);
                        if (wordSet.contains(nextWord) && !visited.contains(nextWord)) {
                            if (nextWord.equals(endWord))
                                found = true;
                            nextLevelVisited.add(nextWord);
                            queue.offer(nextWord);
                            graph.computeIfAbsent(currentWord, key -> new HashSet<>())
                                    .add(nextWord);
                        }
                    }
                    currentWordArray[index] = originChar;
                }
            }
            if (found)
                break;
            visited.addAll(nextLevelVisited);
            nextLevelVisited.clear();
        }
        return found;
    }

    private static void dfs(String beginWord, String endWord,
                            Map<String, Set<String>> graph,
                            List<String> path,
                            List<List<String>> result
    ) {
        if (beginWord.equals(endWord)) {
            result.add(new ArrayList<>(path));
            return;
        }

        Set<String> nextWords = graph.get(beginWord);
        if (nextWords != null) {
            for (String nextWord : nextWords) {
                path.add(nextWord);
                dfs(nextWord, endWord, graph, path, result);
                path.remove(path.size() - 1);
            }
        }
    }

    public static void main(String[] args) {
        @SuppressWarnings("unchecked")
        Tester tester = new Tester(WordLadder2.class, list -> new HashSet<>((List<List<String>>) list));
        tester.addTestCase(
                "hit", "cog",
                Arrays.asList("hot", "dot", "dog", "lot", "log", "cog"),
                Arrays.asList(
                        Arrays.asList("hit", "hot", "dot", "dog", "cog"),
                        Arrays.asList("hit", "hot", "lot", "log", "cog")
                )
        );
        tester.addTestCase(
                "hit", "cog",
                Arrays.asList("hot", "dot", "dog", "lot", "log"),
                Collections.emptyList()
        );
        tester.runTestCases();
    }
}
