# Studious Doodle
Is a simple way to share lab work within class, with no logins and can be veiwed directly.

## Motivation
Since reading data should'nt require auth, I tried to create a system an admin a share data with a simple code which can be viewed by all directly. Which make reading data very quick.  

## How to use
An admin can upload documents to the server and these documents can be access by other student. But due to limitations of the free server by render, I created a special script in [https://studious-doodle.onrender.com/admin/scripts/session_script] which can be used to keep the server awake for longer times to decrease the response times. 

### Routes
#### Admin Routes
\tPOST `https://studious-doodle.onrender.com/<section>/<subject>/<manual>` requires admin token to post manual which has to be provided in Authorization header (type 'Bearer') and also provide the file to be uploaded.
\tDELETE  `https://studious-doodle.onrender.com/<section>/<subject>/<manual>` requires admin token to delete manual which has to be provided in Authorization header (type 'Bearer').
#### User Routes
\tGET `https://studious-doodle.onrender.com/<section>/<subject>/<manual>` to access the given manual.
\tGET `https://studious-doodle.onrender.com/<section>/<subject>` to get the list of manuals.
\tGET `https://studious-doodle.onrender.com/<section>` to get the list of subjects.
\tGET `https://studious-doodle.onrender.com` to get the list of sections.
