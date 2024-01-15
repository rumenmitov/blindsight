# BlindSight Server
![BlindSight Logo](assets/blindsight_logo.png)
## Endpoints
- `/register` _POST_  
> Enters user data into the Postgres database and sends verification email to user.

Receives a **multipart/form-data** with the following fields:  
-- `fname` _string_  
-- `lname` _string_  
-- `email` _string_  
-- `username` _string_  
-- `password` _string_  
  
***
- `/verify/:user_id` _GET_
> Marks user as 'verified' in the database.
  
Receives the `user_id` through the URL.  
  
***
- `/login` _POST_
> Checks if user has the correct credentials. If they are correct, a **JSON of the user object** is returned.  
   
Receives a **multipart/form-data** with the following fields:    
-- `username` _string_  
-- `password` _string_  
    
NOTE: `username` accepts either user's email or username.  
    
***
- `/images` _POST_  
> Decodes `image` and saves to **images** directory as a **PNG**.    
  
Receives a **multipart/form-data** with the following fields:  
-- `name` _string_  
-- `image` _string_    
  
NOTE: `image` must be **base 64 encoded**.
  

