# Gonads
Web server for frontend development. Compiles SASS and JavaScript containing JSX on the fly

## Usage
`gonads -root="./app"`

## Example output
### Input SCSS
**main.scss**
```
html {
     body {
          background: red;
     }
}
```

### Output SCSS
```
html body {
  background: red; }
  ```

### Input JavaScript
**main.jsx**
```
(function() {
    var i = <div />;
    console.log("Hello, world!")
})
```


### Output JavaScript
```
(function () {
    var i = React.createElement("div", null);
    console.log("Hello, world!");
});
```

## Build
Gonads depends on the golang wrapper for libsass, [go-libsass](https://github.com/wellington/go-libsass). To build against libsass you need to set `PKG_CONFIG_PATH` to the directory containing libsass.pc.
