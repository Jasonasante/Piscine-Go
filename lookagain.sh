find -name "*.sh" |cut -d "/" -f2 | sed 's/.sh//g'|sort -r
