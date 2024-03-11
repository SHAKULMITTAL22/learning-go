Feature: React JS Components Testing
  As a QA engineer
  I want to test each React JS component
  So that I can ensure the application is working as expected

Scenario Outline: Test React JS Component
  Given I have a "<component>" to test
  When I render the "<component>"
  Then the "<component>" should render without errors

Examples:
  | component   |
  | "Header"    |
  | "Footer"    |
  | "Sidebar"   |
  | "MainBody"  |
  | "Login"     |
  | "Register"  |
