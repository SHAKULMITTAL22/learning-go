Feature: Functional API Testing
  Scenario: Validate the API endpoints
    Given the API base URL 'http://localhost:3000'
    When I send a GET request to '/valid-endpoint'
    Then the response status should be 200
    And when I send a GET request to '/invalid-endpoint'
    Then the response status should be 404

  Scenario: Test the server response to correct Functional Test Cases from Requirements
    Given the API base URL 'http://localhost:3000'
    When I send a POST request to '/test-cases' with correct test specs
    Then the response status should be 200
    And the response should display the correct test specs
    And when I send a POST request to '/test-cases' with incorrect test specs
    Then the response status should be 400
    And the response should contain an error message

  Scenario: Test the functionality of Cards – Collections & Recovery
    Given the API base URL 'http://localhost:3000'
    When I send a POST request to '/cards' with valid card details
    Then the response status should be 200
    And the response should be as per the business logic
    And when I send a POST request to '/cards' with invalid card details
    Then the response should contain an error message with appropriate status code
