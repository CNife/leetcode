tr -s ' ' '\n' <file.txt | sort | uniq -c | sort -r | awk '{ print $2, $1 }'
