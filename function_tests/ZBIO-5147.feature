Feature: Credit Card Management API

Scenario: Fetching credit card details with a valid credit card number
Given the API base URL 'http://localhost:3000' 
When I send a GET request to '/credit-card-details' with valid credit card number
Then the response status should be 200
And the response should contain credit card details including due date and balance

Scenario: Fetching credit card details with an invalid credit card number
Given the API base URL 'http://localhost:3000'
When I send a GET request to '/credit-card-details' with invalid credit card number
Then the response status should be 400
And the response should contain error message "Invalid credit card details provided"

Scenario: Arranging a call if balance is unpaid and overdue
Given the API base URL 'http://localhost:3000'
When I send a POST request to '/arrange-call' with credit card number, due date, and contact details
Then the response status should be 200
And the response should contain message "Call arranged successfully"

Scenario: Arranging a call with invalid input
Given the API base URL 'http://localhost:3000'
When I send a POST request to '/arrange-call' with invalid input
Then the response status should be 400
And the response should contain error message "Invalid input provided"

Scenario: Updating card balance after payment is collected
Given the API base URL 'http://localhost:3000'
When I send a PUT request to '/update-balance' with payment amount and credit card number
Then the response status should be 200
And the response should contain message "Card balance updated"

Scenario: Updating card balance after payment processing error
Given the API base URL 'http://localhost:3000'
When I send a PUT request to '/update-balance' and a payment processing error occurs
Then the response status should be 400
And the response should contain error message "Payment processing error"
