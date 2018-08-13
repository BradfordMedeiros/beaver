
REPO=$(echo $OPTIONS | awk '{ if ($1 == "repo") { print($2) }}')
echo "$ID $REPO">> resources 
UNIQUE_IDS=$(cat resources | uniq)
echo $UNIQUE_IDS > resources

