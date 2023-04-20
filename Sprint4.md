# Sprint 3 - Group 38
# Github Link to Repository: https://github.com/amyawu/CEN3031-Project.git
# Youtube Link: https://youtu.be/jzXi0Bdh-nc
# Front end readme: https://github.com/amyawu/CEN3031-Project/tree/main/front-end

# Work completed in Sprint 4

In the backend, we did some code clean up and separated our functions into two files endpoint.go and helpers.go. The endpoint.go holds the functions that have endpoints such as GET, POST, PUT and helpers.go holds the helper functions that are used within the endpoint functions. We abstracted the python script into it's own backend API. We also fixed up some existing code to help facilitate image display on the front end through use of tokens and authentication (see frontend for more info). Finally, in Sprint 4 we attempted to set up an online database using Microsoft Azure SQL and editted our code in order to support this new online database. This unfortunately didn't end up working out as there was conflict between the database and the existing code that we weren't anticipating and so we had to abandon the prospect of an online database for lack of time and resources before theclose of Sprint 4.

# Unit Tests and Cypress Test for Frontend
In the frontend, there were multiple new updates that we created. We made sure that the image upload is connected between frontend and backend. We checked that the image upload can be done by users who are logged in and those who are not. However, the image only saves for the users who are logged in. 

We also created the following new elements:

Account: The user can logged in and see what details that he/she had registered under

Profile: The user can edit what details he/she had registered under prior

Display: The user can see all images he/she had uploaded when they are also logged in.

Home: The user upon logging in can access three parts of the site from the home page: edit profile (edit what he/she had registered for), see recens (look at all the images he/she had uploaded), and upload an image (uploading an image)

Navbar: The user can access the different parts of the site wherever page he/she goes on the site. It shifts depending on whether the user is logged in or not. Specifically, if the user is logged in, he/she can log out, but if the user is not logged in, he/she will need to either register or login first.

We provided and edited cypress component and e2d testing respectively to developer and users for the following elements: home, profile, account, login, register, and upload.

However, due to the AuthGuard implementation, we had issues with Cypress only accepting certain tokens, resulting in some errors in our Cypress testing. If we had more time past this semester, we will have to isolate AuthGuard independently from Cypress in order to test all components. Either way, the backend is working and sending the requests as it should be so the code should work despite the token inconsistency with Cypress.

We ran into a CORE issue upon connecting with the backend when the user uploads an image; again, this had to deal with authentication of the user, so if we had more time past this semester, we will focus more on keeping the user logged in and maintaining that authentication for image upload.

# Unit Tests for Backend
The only unit test we created for Go was a test to see if connection to the backend database created with Microsoft Azure SQL would work, called TestDatabaseConnection.
Due to MSSQL failing to work with our existing code and lack of resources and time to find a replacement for Sprint 4's conclusion, this is the only unit test
for backend and it will never succeed. At least we tried.

On the other hand, we created unit tests for the secondary backend for FAST API. This unit tests ensures that the backend server can be spun up on a different port other than the default 8000, as that one is being used for the Go backend.

# Unit Tests for Frontend
We did multiple new Cypress unit testing for the following pages:

Account

Display

Upload

Home

We did new Cypress unit testing for the navbar which appears in every page.

We also provided updates to the following pages:

Login

Register

E2E Testing:
Note: Since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked
Essentially, all pages and navbar (Except when the user is not logged in for upload page) will face this issue and we need more time past the semester to independently test everything.

Navbar- the user can use the navbar to navigate to different parts of the site

Account- the user can see what they had registered their account as for the site

Display- the user can display all images that he/she had previously uploaded

Upload- the user can upload an image and see what image that they had uploaded

Home- the user had logged in and can go to three different parts of the site: edit profile, upload image, and see recent images.

Login- the user fills out a form and logs in

Register- the user fills out a form and registers

Component Testing:

Navbar- make sure the developer can go through other parts of the site for ease of access (since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

Account- make sure the developer can see the details of the profile that the user has registered (since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

Display- make sure the developer can see all the images that the user has uploaded (not including those they have yet to login as user for) (since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

Upload- make sure the developer can see that a logged in user as well as an anonymous user can upload a file and see what file they had uploaded (For logged in user upload, since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

Home- make sure the developer can see that a user is logged in and can go to different parts of the site upon logging in (since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

Login- make sure the developer can see that the user can fill out the login form and logged in (since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

Register- make sure the developer can see that the user can fill out the register form and register his/her account (since AuthGuard is implemented, some parts of the site the developer cannot access as the user needs to be logged in first in which Cypress only accepts certain tokens for, which is not the case for our testing. If we had more time past the semester, we will have to take out AuthGuard to test it independently without user login necessary to make sure all site areas worked)

# Backend Documentation
Link: https://github.com/amyawu/CEN3031-Project/blob/main/Back-end%20Documentation.pdf

# Frontpage Readme
