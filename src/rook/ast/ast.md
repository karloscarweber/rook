# Abstract syntax Tree

We've got statements and expressions

1. Statements produce side effects or actions in the programming language or environment.
2. Expressions return values.

Every program is a series of statements, and within those statements there can be other statements or expressions. But the program is basically a series of statements. Our AST needs to become a tree of s expressions. So `1 + 2 + 3` needs to become `(+ 1 (+ 2 3)).

infix expression have an operator between them. To figure out where each number and operator goes, left or right side we need just a single step of lookahead. More complicated syntax trees require more steps, but let's not worry about that right now.

ASTs have an idea of a NODE, with a left or right branch. What's on the left, happens first. In operator expressions like `4 + 5 + 6` the node is `sum` and the left node is `5` so we would get `(sum 6 (sum 4 5))` when it's all done. That would be represented like this:
```
  exp
  / \
sum  O
    / \
   4  sum
      / \
     5   6
```
or 
```
 sum
 / \
4  sum
   / \
  5   6
```

The topmost node is always either a statement or an expression, depending on the programming language conventions. But it's usually a statement.

## Binding power
Making ASTs from operators and expressions requires making decions about which branch an expression or value goes to. For example: `5 * 2 + 3` binds `5` and `2` together to be executed before the sum of that expression is added to `3`. The S-Expression should look like this: `(sum 3 (mul 5 2))`. The deeper an S-Expression is nested, the higher it's priority to be executed. Translating this to commands would look something like this:
```
val  5
val  2
mul
val  3
sum
```
The values `5` and `2` are placed on the stack first, and then executed. So, knowing all that, we need to know, when we're reading a sequence of values and operators, if a value is attached to the operator on the right or the left, based on precedence of the operator, and what is executed first in a sequence.

the left side, or child is executed first.

so, to figure out what we do when parsing, we start by reading the tokens: values and operators. If we first get a value and then an operator, we make an incomplete s expression: `5 + 2` -> `(sum 5 <exp>)`. An Operator has a higher precedence than an expression, so
