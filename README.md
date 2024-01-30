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






### file_store

curl -X POST \
    -b "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2MTUwNjcsInVzZXJuYW1lIjoiYWRtaW4ifQ.tBwtVeXwpK9Y6XTZ9_qtKOMG0fQgYliI7_hpmX88TWw; Expires=Tue, 30 Jan 2024 11:44:27 GMT;" \        
    -F "file=@4.jpg" \
    -F "name=ax 4" \
    -F "type=jpg" \
    -F "tags=personal" \
    -F "tags=sport" \
    http://0.0.0.0/upload


### retrieve


curl -i -X GET \
   -b "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY2MTUwNjcsInVzZXJuYW1lIjoiYWRtaW4ifQ.tBwtVeXwpK9Y6XTZ9_qtKOMG0fQgYliI7_hpmX88TWw; Expires=Tue, 30 Jan 2024 11:44:27 GMT;" \
 'http://0.0.0.0:8080/retrieve?name=ax&tags&tags=sport&tags=personal'

