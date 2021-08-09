curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jasonasante\"}}){id}}"}' |cut -d "i" -f2 |cut -d ":" -f2 |cut -d "}" -f1

