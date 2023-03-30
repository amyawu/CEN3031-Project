# Sprint 3 - Group 38
# Github Link to Repository: https://github.com/amyawu/CEN3031-Project.git
# Youtube Link:

# Work completed in Sprint 3

  In the backend, we have worked on retrieving and displaying all images stored by a single user by building on implementation
from Sprint 2. This allows us to give users the options to save multiple images to to their profile, allowing them to track their progress over time of any possible risk. We have also managed to begin using bcrypt to hash passwords and store them as well as perform a password check
when logging into the application based on the inputed password. Until now, we've only had verification through the email that
was inputted. Lastly, we have created JWT Authentication tokens to serve to the client in order for requests from a given user (once logged in) be verified easily. Hopefully by Sprint 4, this functionality is secure and ironed out to give a level of protection for user data. Moreover, we plan to implement another level of abstraction for the python scripts by building a FAST API to communicate with the model. This will alleviate current latency problems of having to load relatively large deep learning models each time the function is called, and instead have it run once in an initializtion step.

In the frontend, 

# Unit Tests and Crypress Test for Frontend

# Unit Tests for Backend

The backend has implemented more unit tests to test the hashing ability of bcrypt, as well as compare two hashed passwords together. 

# Backend Documentation
Link: https://github.com/amyawu/CEN3031-Project/blob/main/Back-end%20Documentation.pdf
