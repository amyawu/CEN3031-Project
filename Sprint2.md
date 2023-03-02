# Sprint 2 - Group 38
# Github Link to Repository: https://github.com/amyawu/CEN3031-Project.git

# Work completed in Sprint 2 
During the duration of Sprint 2, we have connected the frontend and the backend as well as made progress on both ends. Integration currently works for login and register page, and we are working on integrating the image uplaod functionality.

In the frontend, we have...

In the backend, we have worked on the program to create individual user files based off email which is a user story we didn't
complete in Sprint 1. We have also created unit tests to test basic functionality of the user data, we weren't able to test
image functionality and are working on finding a way to successfully test image upload functions. Additionally, we have integrated python scripts that load a sample neural network for binary image classification (just a placeholder) that will eventually contain the actual model for melanoma classification.

# Unit Tests and Cypress Test for Frontend

# Unit Tests for Backend
TestRouter - Checks if Gin router is properly being created
TestGetUser - Checks if a User can be fetched by ID using a GET request
TestVerifyUser - Checks if a given user's email and password matches that of an existing user in the database as POST request
TestGetUserByEmail - Similar to GetUser, but uses email as query parameter instead
TestCreateUser -  Checks if a user can be created using a POST request, requiring for the email and password to be unique.

# Backend Documentation
Link: https://github.com/amyawu/CEN3031-Project/blob/main/Back-end%20Documentation.pdf
