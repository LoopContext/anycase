// Package anycase converts strings to various cases. See the conversion table below:
// | Function                                  | Result               |
// | ----------------------------------------- | -------------------- |
// | `ToSnake(s)`                              | `any_kind_of_string` |
// | `ToSnakeAndIgnore(s, '.')`                | `any_kind.of_string` |
// | `ToSnakeUppercase(s)`                     | `ANY_KIND_OF_STRING` |
// | `ToKebab(s)`                              | `any-kind-of-string` |
// | `ToKebabUppercase(s)`                     | `ANY-KIND-OF-STRING` |
// | `ToCamel(s)`                              | `AnyKindOfString`    |
// | `ToLowerCamel(s)`                         | `anyKindOfString`    |
// | `ToDelimited(s, '.')`                     | `any.kind.of.string` |
// | `ToDelimitedUppercase(s, '.', ' ', true)` | `ANY.KIND OF.STRING` |
// | `ToDelimitedUppercase(s, '.', '', true)`  | `ANY.KIND.OF.STRING` |
package anycase
