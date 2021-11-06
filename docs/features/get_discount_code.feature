Feature: Getting discount code
  As a logged in user
  I want to be able to get a discount code
  So that I can get a discount on a purchase.

  Scenario: Successfully get discount code
    Given I have logged in as Customer user
    And I have brand page created
    When I get discount code for Brand A
    Then I should receive discount code for Brand A assigned to current user

  Scenario Outline: Unsuccessfully get discount code due to auth
    Given I have logged in as <user type> user
    When I get discount code for Brand A
    Then I should recieve error about missing permissions for <user type> user
    Examples:
      | user type       |
      | Brand B         |
      | Unauthenticated |

  Scenario: Unsuccessfully get discount code for non existing brand
    Given I have logged in as Customer user
    When I get discount code for Brand Z
    Then I should recieve error that Brand Z does not exist
