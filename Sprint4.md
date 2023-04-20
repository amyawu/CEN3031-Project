# Sprint 3 - Group 38
# Github Link to Repository: https://github.com/amyawu/CEN3031-Project.git
# Youtube Link: 

# Work completed in Sprint 4

In the backend, we did some code clean up and separated our functions into two files endpoint.go and helpers.go. The endpoint.go holds the functions that have endpoints such as GET, POST, PUT and helpers.go holds the helper functions that are used within the endpoint functions. We abstracted the python script into it's own backend API. We also fixed up some existing code to help facilitate image display on the front end through use of tokens and authentication (see frontend for more info). Finally, in Sprint 4 we attempted to set up an online database using Microsoft Azure SQL and editted our code in order to support this new online database. This unfortunately didn't end up working out as there was conflict between the database and the existing code that we weren't anticipating and so we had to abandon the prospect of an online database for lack of time and resources before theclose of Sprint 4.

# Unit Tests and Cypress Test for Frontend



# Unit Tests for Backend
The only unit test we created for Go was a test to see if connection to the backend database created with Microsoft Azure SQL would work, called TestDatabaseConnection.
Due to MSSQL failing to work with our existing code and lack of resources and time to find a replacement for Sprint 4's conclusion, this is the only unit test
for backend and it will never succeed. At least we tried.

On the other hand, we created unit tests for the secondary backend for FAST API. This unit tests ensures that the backend server can be spun up on a different port other than the default 8000, as that one is being used for the Go backend.


# Backend Documentation
Link: https://github.com/amyawu/CEN3031-Project/blob/main/Back-end%20Documentation.pdf
