Feature: React JS Components Testing
  As a QA engineer
  I want to test each React JS component
  So that I can ensure the application is working as expected

Scenario Outline: Test React JS Component
  Given I have a "<component_name>" component
  When I render the "<component_name>" component
  Then I should see the "<component_name>" component rendered correctly

Examples:
| component_name     |
| "Header"           |
| "Footer"           |
| "Navigation Menu"  |
| "Profile Section"  |
| "Content Area"     |
