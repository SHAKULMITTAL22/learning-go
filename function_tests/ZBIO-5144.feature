Feature: API Endpoint Testing
  Scenario: Validate the API endpoints
    Given the correct API endpoint as per the specification
    When I send a request to the API endpoint
    Then the response status should be 200
    Given an incorrect API endpoint
    When I send a request to the API endpoint
    Then the response status should be 404
