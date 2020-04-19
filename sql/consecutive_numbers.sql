SELECT
    DISTINCT Num AS ConsecutiveNums
FROM
    (
        SELECT
            Num,
            IF (@pre = Num, @count := @count + 1, @count := 1) AS Nums,
            @pre := Num
        FROM
            LOGS AS l,
            (
                SELECT
                    @pre := NULL,
                    @count := 1
            ) AS pc
    ) AS n
WHERE
    Nums >= 3;
