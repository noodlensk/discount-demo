Feature: Brand notification of user getting discount code
  As a brand
  I want to be notified about a user getting a discount code
  So that I can process information about the user for my loyalty programme.

  Scenario: Successfully notified about user getting discount code
    Given I have configured notification for Brand A of user discount code generation
    When User get discount code for Brand A
    Then I should receive notification about user getting a discount code
    And Notification contains information about user details
