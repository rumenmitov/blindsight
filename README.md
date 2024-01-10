# BlindSight Server
![BlindSight Logo](assets/blindsight_logo.png)
## Endpoints
- `/register` _POST_  
> Enters user data into the Postres database **Users**.

Receives a **multipart/form-data** with the following fields:  
-- `fname` _string_  
-- `lname` _string_  
-- `email` _string_  
-- `password` _string_  
  
***
- `/images` _POST_  
> Decodes `image` and saves to **images** directory as a **PNG**.    

Receives a **multipart/form-data** with two fields:  
-- `name` _string_
-- `image` _string (base64 encoded)_  
