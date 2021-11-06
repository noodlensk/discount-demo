# glossary

List of terms used across feature definition.

## Terms

- `Brand` is an entity with uniq `id` (and possible other fields like `name`, `description`, etc. but it out of the
  scope of current project).
- `User` - person with personal details such as email address. User can have two types of permissions:
  - Generate codes for given `Brand`
  - Getting codes from any `Brand`
    Theoretically it is possible that User has permissions for multiple brands and also can be a consumer(i.e. fetch
    code)
- `Discount code` - entity with `ID` / `Code` which is belongs to the `Brand` and can be assigned to the `User`. `ID`
  should be uniq per brand.
- `Brand A user` - user with permissions for generating codes for `Brand A`
- `Customer user` - user with permissions for getting codes from any brands
