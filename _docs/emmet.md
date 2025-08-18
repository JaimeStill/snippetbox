# Emmet Quick Reference

## Basic Elements

### Tag Generation
- `div` → `<div></div>`
- `p` → `<p></p>`
- `a` → `<a href=""></a>`
- `img` → `<img src="" alt="">`
- `input` → `<input type="text">`

## Nesting Operators

### Child: `>`
`div>ul>li` → 
```html
<div>
    <ul>
        <li></li>
    </ul>
</div>
```

### Sibling: `+`
`div+p+span` →
```html
<div></div>
<p></p>
<span></span>
```

### Climb-up: `^`
`div>p>span^p` →
```html
<div>
    <p><span></span></p>
    <p></p>
</div>
```

### Grouping: `()`
`div>(header>ul>li*2)+footer>p` →
```html
<div>
    <header>
        <ul>
            <li></li>
            <li></li>
        </ul>
    </header>
    <footer>
        <p></p>
    </footer>
</div>
```

## Multiplication: `*`
`ul>li*5` →
```html
<ul>
    <li></li>
    <li></li>
    <li></li>
    <li></li>
    <li></li>
</ul>
```

## Attributes

### ID: `#`
`div#header` → `<div id="header"></div>`

### Class: `.`
`div.container` → `<div class="container"></div>`
`div.container.active` → `<div class="container active"></div>`

### Custom Attributes: `[]`
`input[type=email placeholder="Email"]` → `<input type="email" placeholder="Email">`
`a[href=# target=_blank]` → `<a href="#" target="_blank"></a>`

### Combining ID, Class, and Attributes
`div#main.container.fluid[data-role=page]` → `<div id="main" class="container fluid" data-role="page"></div>`

## Text Content: `{}`
`p{Click here}` → `<p>Click here</p>`
`a[href=#]{Link text}` → `<a href="#">Link text</a>`

## Numbering: `$`

### Basic Numbering
`ul>li.item$*3` →
```html
<ul>
    <li class="item1"></li>
    <li class="item2"></li>
    <li class="item3"></li>
</ul>
```

### Zero Padding: `$$`, `$$$`
`ul>li.item$$$*3` →
```html
<ul>
    <li class="item001"></li>
    <li class="item002"></li>
    <li class="item003"></li>
</ul>
```

### Reverse Numbering: `@-`
`ul>li.item$@-*3` →
```html
<ul>
    <li class="item3"></li>
    <li class="item2"></li>
    <li class="item1"></li>
</ul>
```

### Starting Number: `@N`
`ul>li.item$@3*3` →
```html
<ul>
    <li class="item3"></li>
    <li class="item4"></li>
    <li class="item5"></li>
</ul>
```

## Implicit Tag Names
- `.container` → `<div class="container"></div>`
- `#header` → `<div id="header"></div>`
- `ul>.item` → `<ul><li class="item"></li></ul>`
- `table>.row>.cell` → `<table><tr class="row"><td class="cell"></td></tr></table>`
- `select>.option*3` → `<select><option class="option"></option>...</select>`

## Common HTML Boilerplate

### HTML5 Document: `!` or `html:5`
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    
</body>
</html>
```

## Practical Examples

### Navigation Menu
`nav>ul.nav>li.nav-item*5>a[href=#]{Item $}` →
```html
<nav>
    <ul class="nav">
        <li class="nav-item"><a href="#">Item 1</a></li>
        <li class="nav-item"><a href="#">Item 2</a></li>
        <li class="nav-item"><a href="#">Item 3</a></li>
        <li class="nav-item"><a href="#">Item 4</a></li>
        <li class="nav-item"><a href="#">Item 5</a></li>
    </ul>
</nav>
```

### Form Structure
`form>(.form-group>label[for=input$]{Label $}+input#input$[type=text])*3+button[type=submit]{Submit}` →
```html
<form>
    <div class="form-group">
        <label for="input1">Label 1</label>
        <input id="input1" type="text">
    </div>
    <div class="form-group">
        <label for="input2">Label 2</label>
        <input id="input2" type="text">
    </div>
    <div class="form-group">
        <label for="input3">Label 3</label>
        <input id="input3" type="text">
    </div>
    <button type="submit">Submit</button>
</form>
```

### Card Component
`.card>(header.card-header>h3{Title})+.card-body>p{Content}+a.btn[href=#]{Read More}` →
```html
<div class="card">
    <header class="card-header">
        <h3>Title</h3>
    </header>
    <div class="card-body">
        <p>Content</p>
        <a href="#" class="btn">Read More</a>
    </div>
</div>
```

### Table with Headers
`table>(thead>tr>th*3{Header $})+(tbody>tr*3>td*3{Cell $})` →
```html
<table>
    <thead>
        <tr>
            <th>Header 1</th>
            <th>Header 2</th>
            <th>Header 3</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Cell 1</td>
            <td>Cell 2</td>
            <td>Cell 3</td>
        </tr>
        <tr>
            <td>Cell 1</td>
            <td>Cell 2</td>
            <td>Cell 3</td>
        </tr>
        <tr>
            <td>Cell 1</td>
            <td>Cell 2</td>
            <td>Cell 3</td>
        </tr>
    </tbody>
</table>
```

## Tips
- Emmet works in most modern code editors (VS Code, Sublime, Atom, etc.)
- Press Tab or Enter to expand abbreviations
- Combine operators for complex structures
- Use parentheses for grouping when needed
- Implicit tags save typing (`.class` becomes `<div>`, `ul>.item` makes `<li>`)