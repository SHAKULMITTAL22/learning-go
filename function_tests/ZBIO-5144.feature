Feature: Functional Test Cases from Requirements Testing
Scenario: Test the server response to correct Functional Test Cases from Requirements
Given the correct Functional Test Cases from Requirements
When I send a POST request to '/api/test-cases' with the correct Functional Test Cases in the body
Then the response status should be 200
And the response body should display the correct test specs
And when I send a POST request to '/api/test-cases' with incorrect Functional Test Cases in the body
Then the response status should be 400
And the response should contain an error message
