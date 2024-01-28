# BlindSight Server
![BlindSight Logo](assets/blindsight_logo.png)
## Endpoints
- `/ping` _GET_
> Returns the string 'pong' if successful.
   
  
***
- `/users` _GET_
> Returns a **JSON of an array containing all usernames and emails**.
  
  
***
- `/register` _POST_  
> Enters user data into the Postgres database and sends verification email to user.

Receives a **multipart/form-data** with the following fields:  
-- `fname` _string_  
-- `lname` _string_  
-- `email` _string_  
-- `username` _string_  
-- `password` _string_  
  
***
- `/verify` _POST_
> Marks user as 'verified' in the database. Returns a **JSON of the user object** or an `AuthError` (see below).
  
Receives a **multipart/form-data** with the following fields:  
-- `verification_code` _string_
  
Possible `AuthError` status codes:
-- `201` - user is verified
-- `202` - user input is not a number
-- `203` - credentials are wrong
-- `204` - unknown error
  
***
- `/login` _POST_
> Checks if user has the correct credentials. If they are correct, a **JSON of the user object** is returned, otherwise an `AuthError`
is sent (see below).  
   
Receives a **multipart/form-data** with the following fields:    
-- `username` _string_  
-- `password` _string_  
    
NOTE: `username` accepts either user's email or username.  
   
Possible `AuthError` status codes:
-- `201` - user credentials are correct
-- `202` - user input is not a number
-- `203` - credentials are wrong
-- `204` - unknown error
    
***
- `/image` _POST_  
> Decodes `image`, analyzes it and sends an instruction back.    
  
Receives a **multipart/form-data** with the following fields:  
-- `name` _string_  
-- `image` _string_    
  
NOTE: `image` must be **base 64 encoded**.
  

