SELECT
    a.Score AS Score,
    (
        SELECT
            count(DISTINCT b.Score)
        FROM
            Scores b
        WHERE
            b.Score >= a.Score
    ) AS Rank
FROM
    Scores a
ORDER BY
    Score DESC;
