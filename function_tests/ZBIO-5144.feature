Feature: Validate the API endpoints
Scenario: Validate correct API endpoint
Given the correct API endpoint as per the specification
When I send a GET request to the endpoint
Then the response status should be 200

Scenario: Validate incorrect API endpoint
Given an incorrect API endpoint
When I send a GET request to the endpoint
Then the response status should be 404
