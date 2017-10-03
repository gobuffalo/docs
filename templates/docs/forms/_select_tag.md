<%= title("Select Tags") %>

To build your `&lt;select&gt;` tags inside forms Tags provide 3 convenient ways to add your `&lt;select&gt;` options: `form.SelectOptions`, `map[string]interface{}` or `[]string`, all of them by passing an `options` field into the `form.SelectTag` options like:

```erb
&lt;%= f.SelectTag("TalkFormatID", {options: talkFormats}) %&gt;
```
or

```erb
&lt;%= f.SelectTag("TalkFormatID", {options: ["one", "two"]}) %&gt;
```

Which will use the same value for the `value` attribute and the body of the option, or:

```erb
&lt;%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}}) %&gt;
```

Which allows us to define the options map inside the view.

### Selectable Interface

Another alternative for the select options is to pass a list of structs that meet the `form.Selectable` interface.

Which consist of two functions:

```go
//Selectable allows any struct to become an option in the select tag.
type Selectable interface {
  SelectValue() interface{}
  SelectLabel() string
}
```

By implementing this interface tags will call `SelectValue` and `SelectLabel` to get the option Value and Label from implementer.

### Selected

Tags will add the `selected` attribute to the option that has the same value than the one it receives on the `value` option of the `form.SelectTag`, so you don't have to look for the option that has equal value than the selected one manually, p.e:

```erb
&lt;%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}, value: 2}) %&gt;
```

Produces:

```html
&lt;div class="form-group"&gt;
  &lt;label&gt;TalkFormatID&lt;/label&gt;
  &lt;select class="form-control" id="talk-TalkFormatID" name="TalkFormatID"&gt;
    &lt;option value="1"&gt;one&lt;/option&gt;
    &lt;option value="2" selected&gt;two&lt;/option&gt;
  &lt;/select&gt;
&lt;/div&gt;
```

And similarly with the `form.SelectOptions` slice:

```erb
&lt;%= f.SelectTag("TalkFormatID", {options: talkFormats, value: 2}) %&gt;
```
