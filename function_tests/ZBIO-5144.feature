Feature: Validate the React component rendering
Scenario: Testing React component rendering under normal and slow network conditions
Given the API base URL 'http://localhost:3000'
When I load the application
Then all components should render as expected without any errors
And when I simulate slow network speed
Then all components should still render correctly
