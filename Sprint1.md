# Sprint 1 - Group 38

# User Stories Planned For The Future
As a researcher, I would like for the user to have the option to store the image and allow its usage for further model training.

As an app user, I would like the option to change my password so that I can access my account in the event I forget my log in.

As an app user, I would like the option to sort through my saved images based on the date I uploaded them.

As an app user, I would like to be able to delete images that I have stored so that I can get rid past images.

As an app user, I would like to be able to create an account on the application such that I am able to have a place to upload images safely.


# User Stories Planned For Sprint 1
As an app user, I would like to have the option for the image to be saved and sent for further training of the classification model, so that performance may be improved.

As an app user, I would like to be able to select an image to be submitted so that it may be processed through the model.

As an app user, I would like to receive useful information based on the specifics provided by my submitted data.

As an app user, I would like to be able to upload images that may be stored and retrieved later to track health progression.

As a researcher, I would like the user to be able to store relevant information regarding health or demographic data that may be useful for tracking of health status.

As an app user, I would like to be able to store several images at a time so that I don't waste time uploading them one at a time.

As an app user, I would like to count with an easy to use interface that provides me with all the options I need to undergo the process easily.

# Successful Stories
Out of the Sprint 1 user stories, we have accomplished storing of data in the form of users, which include images stored as part of the data. Within this data we have
different parameters we have stored per user including name, email, password, gender, and age. While we may add ethnicity to this list, these are the pieces of info we have for our users. Along with this, we have a mock register and login page created on the front-end as the start of an easy to use interface for the customer. It works with our current MongoDB mock database so we are hoping to achieve further functionality based on the back-end database with Go. After Sprint 1, we also are looking to add more parameters like the ones mentioned above to the account registration page to match the Go back-end database in the future sprints.

# Unsuccessful Stories
The failures we had for the back-end in the early stages of getting functionality include some bugs in the user struct of the code. While we are able to submit multiple images at once, only one of these images is tracked due to how our code is set up. In order to fix this and account for the other images uploaded, we would need to create individual directories for each account on the application that store all the images. This is something we will work on in the future Sprints, as well as deleting images and sorting them by date added. 
In the front-end, we experienced errors within various stages of difficulty- both in technical and human. For technical, we solved failures in having the same stable GitHub version by mostly with pulling and pushing changes on different GitHub branches. Furthermore, we overcame the connection errors to the mock database on MongoDB by working through VSCode instead of the terminal. We also concluded that our HTTP conflicts with the decryptions of the usernames and passwords were solved besides reverting the last stable version of our program by activating the Postman hidden headings. Moreover, for issues due to human error, we figured out the problems with GitHub and the different directories in our web application by watching tutorials and implementing what we learned as well as making sure we had installed all the dependencies needed.