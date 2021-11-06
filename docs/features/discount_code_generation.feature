Feature: Discount code generation
  As a brand
  I want to have discount codes generated for me
  So that I don’t have to deal with this administration myself.

  Background:
    Given I have "Brand A" created

  Scenario: Successfully generate discount codes
    Given I have logged in as Brand A user
    When I generate N discount codes for Brand A
    Then I should have N discount codes generated for Brand A

  Scenario Outline: Unsuccessfully generate discount codes due to auth
    Given I have logged in as <user type> user
    When I generate N discount codes for Brand A
    Then I should recieve error about missing permissions for <user type> user
    Examples:
      | user type       |
      | Brand B         |
      | Customer        |
      | Unauthenticated |

  Scenario Outline: Unsuccessfully generate discount codes providing wrong information about desired number of codes
    Given I have logged in as Brand A user
    When I generate <number of codes> discount codes for Brand A
    Then I should recieve error about incorrect <number of codes> discount codes provided
    Examples:
      | number of codes         |
      | negative                |
      | zero                    |
      | bigger than max allowed |

  Scenario: Unsuccessfully generate discount codes for non existing brand
    Given I have logged in as Brand A user
    When I generate N discount codes for Brand Z
    Then I should recieve error that Brand Z does not exist
