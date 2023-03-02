# Sprint 2 - Group 38
# Github Link to Repository: https://github.com/amyawu/CEN3031-Project.git
# YouTube Link: https://youtu.be/p1Q5EuL-jcI

# Work completed in Sprint 2 
During the duration of Sprint 2, we have connected the frontend and the backend as well as made progress on both ends.
In the frontend, we have worked on the unit and Cypress tests. The E2E Cypress testing is within the spec.cy.ts. It is on the login, register, and upload pages. Specifically, it tests if the user can login or register with an email and password and if the user can switch between login and register pages. Within upload, it checks if the user can access the upload page since we are in the process of connecting it physically to the backend in choosing and submitting a file upload. In terms of component testing, we did a component test for each login, register, and upload components. 
Similar to E2E testing, the site ran for each component and made sure that the developer can fill out and submit a form for the login and register pages as well as testing if the form is responsive without sending an empty output to the backend. We also checked if the developer can switch between the login and register pages via the button on each alternative page of login and register. As stated for the upload page, we verified that the developer can access the upload page by testing the upload component. Moreover, we are in the process of importing all the CSS modules for the component Cypress tests so that the developers' view of the site matches the users' view of the site. We had only imported some of the CSS modules for the component testing, and the rest needs to be tested one by one with the component testing.
We also developed the start of the upload page as the webpage, but it is not connected. In other words, we can upload any files, but the backend is believed to be only able to save images as a theory that needs to be tested. We are in the process of connecting the backend with Cloudinary and photo storage with the upload page so the path of the files can be saved.
In the backend, we have worked on the program to create individual user files based off email which is a user story we didn't
complete in Sprint 1. We have also created unit tests to test basic functionality of the user data, we weren't able to test
image functionality and are working on finding a way to successfully test image upload functions.

# Unit Tests and Cypress Test for Frontend
Cypress Test
E2E Testing via './cypress/e2e/spec.cy.ts' and connected to backend via Go
- submit a filled out username and password for login page
- submit a filled out username and password for register page
- switch to login page from register page
- switch to register page from login page
- visit upload page
Unit Tests
Component testing via './login/login.component.cy.ts'
- submit an empty form for login page just to see if the page is responsive (won't send the data out to backend since it's empty)
- submit a filled out username and password for login page
- switch to register page from login page
Component testing via './register/register.component.cy.ts'
- submit an empty form for register page just to see if the page is responsive (won't send the data out to backend since it's empty)
- submit a filled out username and password for register page
- switch to register page from register page
Component testing via './upload/upload.component.cy.ts'
- visit upload page
# Unit Tests for Backend
TestRouter -
TestGetUser -
TestVerifyUser -
TestGetUserByEmail - 
TestCreateUser - 

# Backend Documentation
Link: https://github.com/amyawu/CEN3031-Project/blob/main/Back-end%20Documentation.pdf
