Feature: Validate API endpoints
  Scenario: Test correct API endpoint
    Given the correct API endpoint as per the specification
    When I send a GET request to '/api/endpoint'
    Then the response status should be 200

  Scenario: Test incorrect API endpoint
    Given an incorrect API endpoint
    When I send a GET request to '/api/wrong-endpoint'
    Then the response status should be 404
