Feature: Testing Repository Details
Scenario: Testing HTTP GET method for fetching repository details
Given the API base URL 'http://localhost:3000'
When I send a GET request to '/repository-details'
Then the response status code should be 200
And the response should contain 'repository details'

Scenario: Testing HTTP GET method for fetching repository details with incorrect details
Given the API base URL 'http://localhost:3000'
When I send a GET request to '/repository-details-invalid'
Then the response status code should be 404
