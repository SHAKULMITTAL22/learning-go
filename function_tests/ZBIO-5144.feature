Feature: Functional Testing
  Scenario Outline: Repository Details
    Given the system has repository details "<repository_details>"
    When the system reads and understands the repository details
    Then the source code should compile without errors
    Examples:
      | repository_details |
      | Valid details |
      | Invalid details |
      | Incomplete details |

  Scenario Outline: API Integration
    Given the system has APIs exposed at the HSBC Developer Portal "<api>"
    When the system accesses the APIs
    Then the APIs should return the correct responses as per the API specifications
    Examples:
      | api |
      | Valid API endpoints |
      | Invalid API endpoints |
      | Unavailable API endpoints |

  Scenario Outline: Cards – Collections & Recovery Functionality
    Given the system has actions related to Cards – Collections & Recovery functionality "<actions>"
    When the system executes the actions
    Then the actions should be executed correctly and the system should respond as per the requirements
    Examples:
      | actions |
      | Valid actions |
      | Invalid actions |
      | Invalid data |

  Scenario Outline: TDD
    Given the system has functional Test Cases from Requirements and Source Code "<test_cases>"
    When the system generates test specs based on the functional test cases and source code
    Then the system should be able to generate test specs correctly
    Examples:
      | test_cases |
      | Complete test cases |
      | Incomplete test cases |
      | Invalid test cases |

Feature: Non-Functional Testing
  Scenario Outline: System Performance
    Given the system is under high load or stress conditions "<conditions>"
    When the system operates under these conditions
    Then the system should be able to handle these conditions without significant performance degradation
    Examples:
      | conditions |
      | High loads |
      | Extremely high loads |

  Scenario Outline: API Response Time
    Given the system has API requests "<requests>"
    When the system processes these requests
    Then the APIs should return responses within an acceptable time frame as per the API specifications
    Examples:
      | requests |
      | Regular data |
      | Large amounts of data |

  Scenario Outline: System Security
    Given the system is exposed to various security threats or attacks "<threats>"
    When the system operates under these threats
    Then the system should be able to withstand these threats without compromising the system's integrity
    Examples:
      | threats |
      | Regular threats |
      | Sophisticated threats |

  Scenario Outline: System Availability
    Given the system operates under various operating conditions "<conditions>"
    When the system is used under these conditions
    Then the system should be available for use
    Examples:
      | conditions |
      | Regular conditions |
      | Extreme conditions |

  Scenario Outline: Mockito Virtualisation
    Given the system has dependencies that need to be virtualised/mocked "<dependencies>"
    When Mockito virtualises/mocks these dependencies
    Then Mockito should be able to correctly virtualise/mock the dependencies
    Examples:
      | dependencies |
      | Regular dependencies |
      | Complex dependencies |
