# Studious Doodle
Is a simple way to share lab work within class, with no logins read and writes.<br>

## Motivation
Since sharing lab codes via email or whatsapp is not convenient and requires phones (which is not allowed in our college labs). So I created a simple system where people can share lab codes inside the terminal (using curl) which makes the whole process faster and hastle-free. 

## How to use
An admin can upload documents to the server and these documents can be access by other student. But due to limitations of the free server by render, I created a special script in [https://studious-doodle.onrender.com/admin/scripts/session_script] which can be used to keep the server awake for longer times to decrease the response times. 

### Routes
#### Admin Routes
 POST `https://studious-doodle.onrender.com/<section>/<subject>/<manual>` requires admin token to post manual which has to be provided in Authorization header (type 'Bearer') and also provide the file to be uploaded.<br>
 DELETE  `https://studious-doodle.onrender.com/<section>/<subject>/<manual>` requires admin token to delete manual which has to be provided in Authorization header (type 'Bearer').<br>
#### User Routes
  GET `https://studious-doodle.onrender.com/<section>/<subject>/<manual>` to access the given manual.<br>
  GET `https://studious-doodle.onrender.com/<section>/<subject>` to get the list of manuals.<br>
  GET `https://studious-doodle.onrender.com/<section>` to get the list of subjects.<br>
  GET `https://studious-doodle.onrender.com` to get the list of sections.<br>
