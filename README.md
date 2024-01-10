# BlindSight Server
![BlindSight Logo](assets/blindsight_logo.png)
## Endpoints
- `/images` _POST_  
Receives a **multipart/form-data** with two fields: `name` _string_, `image` _string (base64 encoded)_  
Decodes `image` and saves to in **images** directory as a **PNG** 
