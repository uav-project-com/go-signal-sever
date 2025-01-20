# Test create user
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"test1","phone":"0989455664"}' \
  http://127.0.0.1:9090/api/v1/user -vv