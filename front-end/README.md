# Sprint 3 - Group 38
# Front end documentation and demo

# Demo
# YouTube Link: https://youtu.be/kEzdAiEbDjE

# Documentation
# Dependencies:
Angular
Go
Node.js

Run npm start to start the hosting service
Run go run . to start the backend service

# Upon going to http://localhost:4200/
You will have to login first in order to access your own individual data with the following categories:

Account

Profile

Register

Login

Home

However, uploading an image is generic and can be accessible by all users.

There is a navbar for users to navigate.

# Account: 
Page where users can review their current information. The page fetches the user info found in the backend, and displays them in a readable format for them to assess and manage.

# Profile: 
Page where users can edit their information and update it as they seem fit. Once the new data is selected, the page sends a request to update the user’s info with the new values chosen. It is possible to only update one specific value rather than having to fill them all entirely.

# Register: 
Page where users enter new credentials to register them. Once new credentials are entered, the user is able to fully access all resources provided by the website.

# Login: 
Page where users enter their existing credentials to login. The user cannot create any new data from this page, as it must be previously registered for them to access them in this page. Once logged in, the user has full access to all resources provided by the website.

# Upload: 
Page where images are uploaded. Even though login is not needed, the page’s behavior will change depending on if you are logged in or not. If so, all the images will be stored to the specific user uploading them. However, if no one is logged in, then the image will be sent to a general database where it will be used by our Melanoma AI to more easily recognize the condition.

# Display: 
Page where specific images are displayed. Unlike Image Upload, you can only access this page once you are logged in, as you can only see the images belonging your specific user.

# Home: 
Main page of the app. You must be logged in and have a valid token to access this page. From this page, you are able to access every single page in the website, and can easily access any type of information you might need.
