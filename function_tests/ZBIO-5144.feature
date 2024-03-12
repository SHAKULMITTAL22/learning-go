Feature: Testing System Performance
Scenario: Testing system performance under multiple concurrent requests
Given the API base URL 'http://localhost:3000'
When I send multiple concurrent GET requests to '/api-details'
Then the system should handle the requests without any performance degradation

Scenario: Testing system performance under extreme load
Given the API base URL 'http://localhost:3000'
When I send an extremely high number of concurrent GET requests to '/api-details'
Then the system should handle the requests without crashing
