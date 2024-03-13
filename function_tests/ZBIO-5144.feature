Feature: API endpoints validation
Scenario: Testing correct API endpoint
Given the correct API base URL as per the specification
When I send a GET request to the endpoint
Then the response status should be 200

Scenario: Testing incorrect API endpoint
Given the incorrect API base URL
When I send a GET request to the endpoint
Then the response status should be 404
