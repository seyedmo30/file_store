# file_store



        curl -X POST \
            -b "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDY1NjI3ODMsInVzZXJuYW1lIjoiYWRtaW4ifQ.HECkBXg41I4iHtKOX7Sa2Y_guD1PT4u7oY5oTow8n_U; Path=/; Expires=Mon, 29 Jan 2024 21:13:03 GMT;" \
            -F "file=@2.mp3" \
            -F "name=ax" \
            -F "type=File_Type" \
            http://localhost:8086/upload