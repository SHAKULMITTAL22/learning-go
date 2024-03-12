Feature: Testing API that interacts with React JS application
Scenario: Get all React components
Given the API base URL 'http://localhost:3000'
When I send a GET request to '/react-app-components'
Then the response status should be 200
And the response should contain 'component details'
