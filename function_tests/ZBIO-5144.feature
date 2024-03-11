FUNCTIONAL TEST CASE:
Functional Test Cases:

1. Test Case for Repository Details:
   - Input: Repository details provided by the HSBC team
   - Expected Output: Successful access and retrieval of source code, requirements, and logs
   - Edge Case: Repository details are incorrect or missing

2. Test Case for API Integration:
   - Input: APIs exposed at the HSBC Developer Portal
   - Expected Output: Successful interaction with the APIs and correct response received
   - Edge Case: APIs are unavailable or return an error

3. Test Case for Track 1- Green Field Development:
   - Input: Functionality of our choice from the area of Cards – Collections & Recovery
   - Expected Output: Successful implementation and testing of chosen functionality
   - Edge Case: Chosen functionality is not available or returns an error

Non-Functional Test Cases:

1. Test Case for Test framework/platform:
   - Input: REST Assured, Karate (API testing framework)
   - Expected Output: Successful execution of tests using the chosen framework/platform
   - Edge Case: Framework/platform is unavailable or incompatible

2. Test Case for API Specifications:
   - Input: RAML or Swagger
   - Expected Output: Successful creation and execution of API specifications
   - Edge Case: Specifications are incorrect or missing

3. Test Case for DevOps:
   - Input: Jenkins pipeline
   - Expected Output: Successful integration and deployment using Jenkins pipeline
   - Edge Case: Pipeline is broken or unavailable

Gherkin Scenarios:

Scenario: Test Repository Details
Given the repository details provided by the HSBC team
When I access the repository
Then I should be able to retrieve the source code, requirements, and logs

Scenario: Test API Integration
Given the APIs exposed at the HSBC Developer Portal
When I interact with the APIs
Then I should receive the correct response

Scenario: Test Green Field Development
Given the functionality of our choice from the area of Cards – Collections & Recovery
When I implement and test the chosen functionality
Then it should be successful and return the correct results

Scenario: Test Test framework/platform
Given the REST Assured, Karate (API testing framework)
When I execute the tests using the chosen framework/platform
Then it should be successful

Scenario: Test API Specifications
Given the RAML or Swagger
When I create and execute the API specifications
Then it should be successful and return the correct results

Scenario: Test DevOps
Given the Jenkins pipeline
When I integrate and deploy using the Jenkins pipeline
Then it should be successful and return the correct results.