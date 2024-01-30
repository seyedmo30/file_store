### hello world

curl -i -X GET \
 'http://0.0.0.0/'



### create user

curl -i -X POST \
   -d "password=admin" \
   -d "username=admin" \
 'http://0.0.0.0/createuser'


### login

curl -i -X POST \
   -H "Content-Type:application/x-www-form-urlencoded" \
   -d "password=admin" \
   -d "username=admin" \
 'http://0.0.0.0/login'


### set cookie

like this :

Set-Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2MDM0OTksInVzZXJuYW1lIjoiYWRtaW4ifQ.gJQCB--gdpwE1DDwEcejG8GaPc9ZEyGZDx23Yk6SUrE; Expires=Tue, 30 Jan 2024 08:31:39 GMT






# file_store

curl -X POST \
    -b "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY1NjI3ODMsInVzZXJuYW1lIjoiYWRtaW4ifQ.HECkBXg41I4iHtKOX7Sa2Y_guD1PT4u7oY5oTow8n_U; Path=/; Expires=Mon, 29 Jan 2024 21:13:03 GMT;" \
    -F "file=@2.mp3" \
    -F "name=ax" \
    -F "type=File_Type" \
    http://localhost:8086/upload