# Go Rook todo
- [x] Make sure that lexer catches const declarations: `const name = expression`. and that it's reflected in the tokens.
- [x] Make sure the lexer does not capture trailing whitespaces.
- [x] Capture type declarations in lexer.
- [-] Build AST and Establish basic syntax.
  - [x] Modules
  - [ ] Types
  - [ ] Functions (Types)
  - [ ] Declarations (kinda covered above)
  - [ ] Expressions
  - [ ] Value Literals
- [ ] Establish where types are expected when declaring a new type or aliasing an existing type.
- [ ] Document Syntax in syntax test.
- [x] Get started on Pratt Parser.

# old todo
- [x] Catch malformed type declarations and populate an error array.
- [x] Move FlagError from small_compiler.js to /error/index.js and generalize.
- [ ] Compile Function declarations and put them into the type module.
- [ ] Prevent redeclaration of builtin numeric types.
- [ ] Catch redefined and circular Types.
- [ ] Replace Type identifiers with their expanded types.
- [ ] Flesh out small_compiler with functions and more basics.
- [ ] Catch malformed function declarations.
- [ ] Enable void return values for Functions.
- [ ] Check that Functions produce a stack equal to their result type. like. If the result is i32, then we need to leave i32 on the stack. [reference](https://webassembly.github.io/spec/core/syntax/modules.html#functions)
- [ ] Ensure that functions that return void leave nothing on their stack.
- [ ] Make an npm package for Rook
- [ ] Test CLI of the npm package locally.
