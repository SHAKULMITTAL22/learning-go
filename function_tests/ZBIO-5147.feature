Feature: Testing Credit Card API
Scenario: Fetch Credit Card Details
Given a valid credit card number '1234567890123456'
When I send a GET request to '/credit-card-details'
Then the response status should be 200
And the response should contain 'due date' and 'balance'
But when I send a GET request with an invalid credit card number '123'
Then the response status should be 400
And the response should contain 'Error: Invalid credit card number'
