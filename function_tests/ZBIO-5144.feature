Feature: Testing APIs
Scenario: Testing HTTP GET method for fetching API details
Given the API base URL 'http://localhost:3000'
When I send a GET request to '/api-details'
Then the response status code should be 200
And the response should contain 'API details'

Scenario: Testing HTTP GET method for fetching API details with incorrect details
Given the API base URL 'http://localhost:3000'
When I send a GET request to '/api-details-invalid'
Then the response status code should be 404
