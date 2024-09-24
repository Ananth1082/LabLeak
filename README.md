# Lab Leak
Is a simple way to share lab work within class, with no logins read 


## How to use
An admin can upload documents to the server and these documents can be access by other student. But due to limitations of the free server by render, I created a special script in [https://lableak.onrender.com/admin/scripts/session_script] which can be used to keep the server awake for longer times to decrease the response times. 

### Routes
#### Admin Routes
 POST `https://lableak.onrender.com/<section>/<subject>/<manual>` requires admin token to post manual which has to be provided in Authorization header (type 'Bearer') and also provide the file to be uploaded.<br>
 DELETE  `https://lableak.onrender.com/<section>/<subject>/<manual>` requires admin token to delete manual which has to be provided in Authorization header (type 'Bearer').<br>
#### User Routes
  GET `https://lableak.onrender.com/<section>/<subject>/<manual>` to access the given manual.<br>
  GET `https://https://lableak.onrender.com/<section>/<subject>` to get the list of manuals.<br>
  GET `https://https://lableak.onrender.com/<section>` to get the list of subjects.<br>
  GET `https://https://lableak.onrender.com/` to get the list of sections.<br>


## Features List
- Enable folder uploading system.
- Scripts to unzip and place folders in required place.
- 'Get' directly should send zip folder of the code.

### Lab Leak CLI
![Screenshot 2024-08-11 171407](https://github.com/user-attachments/assets/c20a75f0-8402-418b-9c74-5e99ec1f9f73)
