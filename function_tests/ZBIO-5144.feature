Feature: Testing Repository Details API
Scenario: Testing HTTP GET method for valid repository details
Given the API base URL 'http://localhost:3000' 
When I send a GET request to '/repository-details' with valid details
Then the response status should be 200
And the response should display the correct repository details

Scenario: Testing HTTP GET method for invalid repository details
Given the API base URL 'http://localhost:3000' 
When I send a GET request to '/repository-details' with invalid details
Then the response status should be 404
And the response should display an error message
