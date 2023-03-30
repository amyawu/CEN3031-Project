# Sprint 3 - Group 38
# Github Link to Repository: https://github.com/amyawu/CEN3031-Project.git
# Youtube Link: https://youtu.be/lRR2__19xAI

# Work completed in Sprint 3

  In the backend, we have worked on retrieving and displaying all images stored by a single user by building on implementation
from Sprint 2. This allows us to give users the options to save multiple images to to their profile, allowing them to track their progress over time of any possible risk. We have also managed to begin using bcrypt to hash passwords and store them as well as perform a password check
when logging into the application based on the inputed password. Until now, we've only had verification through the email that
was inputted. Lastly, we have created JWT Authentication tokens to serve to the client in order for requests from a given user (once logged in) be verified easily. 

Hopefully by Sprint 4, this functionality is secure and ironed out to give a level of protection for user data. Moreover, we plan to implement another level of abstraction for the python scripts by building a FAST API to communicate with the model. This will alleviate current latency problems of having to load relatively large deep learning models each time the function is called, and instead have it run once in an initializtion step.

In the frontend, we worked on image upload connection to the backend, profile page, and Cypress testing on the profile page. Backend needed the whole document location whereas Angular frontend development only provides just the image destination's directory. Therefore, we managed to figure out how to add one image to the backend. Our next steps are to be able to upload not just one image but multiple. Furthermore, we managed to expand more on user registration. We have not just the username and password registration but also a development of the user profile page. 

We also conducted Cypress testing on the profile page. Mainly, we focused on e2e and component testing for both developer and user inputs. Specifically, we required the user to enter his or her name, age, gender, and ethnicity to determine who he or she is before the individual logs in as a user. Furthermore, we enabled if the user tried to jump straight to profile without making an account, he or she would have to sign in first to identify who he or she is. In an issue with what we addressed within the recording, the e2e testing is supposed to be for the user while the component testing is supposed to be for the developer. In the Sprint 3 file, we meant to correct what we said within the video with the preceding statement for future acknowledgement of future Cypress testing.

In Sprint 4, we are looking to get each image's whole destination to fully integrate the backend and the frontend together. We are also hoping for more functionality in connecting the user to his or her account by starting the process to store the name, age, gender, and ethnicity of the individual from the front end sides to our backend database. For Sprint 4 or even past Sprint 4, we have to Cypress test the image upload button of the profile page once we fix the matter of uploading more than one image to the site and can connect the user to his or her profile information. However, we may also need extra time past Sprint 4 to fully connect the image upload to the identified account to the user as we are looking to upload more than one image first.

# Unit Tests and Cypress Test for Frontend

The frontend developed more Cypress testing with e2e and component testing specifically with:

Component testing

ProfileComponent mounts: check if the developer can access the profile page after registering

profile page without response: check if the developer can submit without filling out the profile page form

fill out profile page with response: checks if the developer can fill out the profile page

fill out profile page with response and submit: checks if the developer can fill out the profile page and submit

check if we can switch to login page from profile: checks if the developer can swap to login page

switches to login page from profile: checks if the developer clicks the button and it swaps to the login page

E2E testing

visiting profile spec: check if the user can access the profile page after visiting the register and login pages

fills out the profile form and submits it: check if the user can fill out the profile form and submit it

check if the user wants to switch to login from profile: check if the user can click the login button to swap from the profile page to the login site

# Unit Tests for Backend

The backend has implemented more unit tests to test the hashing ability of bcrypt, as well as compare two hashed passwords together. 

TestRouter - Checks if Gin router is properly being created

TestGetUser - Checks if a User can be fetched by ID using a GET request

TestVerifyUser - Checks if a given user's email and password matches that of an existing user in the database as POST request

TestGetUserByEmail - Similar to GetUser, but uses email as query parameter instead

TestCreateUser - Checks if a user can be created using a POST request, requiring for the email and password to be unique.

TestHashPassword - Tests to see if a password can be hashed for storage in bcrypt. Will fail if there's an error or if the hash is empty.

TestUserLogin - Tests the comparison function from bcrypt to make sure that inputted password matches the saved hash password.

# Backend Documentation
Link: https://github.com/amyawu/CEN3031-Project/blob/main/Back-end%20Documentation.pdf
