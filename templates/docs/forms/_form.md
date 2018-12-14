<%= title("Basic Forms") %>

The `form` helper can be used to generate HTML forms. Since this type of form isn't attached to any particular "model" all information must be passed as options to the form and it's methods.

<%= codeTabs() { %>

```erb
// templates/talks/edit.html

&lt;%= form({action: talkPath({id: 3}), method: "PUT"}) { %&gt;
  &lt;div class="row"&gt;
    &lt;div class="col-md-12"&gt;
      &lt;%= f.InputTag({name:"Title", value: talk.Title }) %&gt;
    &lt;/div&gt;

    &lt;div class="col-md-6"&gt;
      &lt;%= f.TextArea({value: talk.Abstract, hide_label: true }) %&gt;
    &lt;/div&gt;

    &lt;div class="col-md-6"&gt;
      &lt;%= f.SelectTag({name: "TalkFormatID", value: talk.TalkFormatID, options: talk_formats}) %&gt;
      &lt;%= f.SelectTag({name: "AudienceLevel", value: talk.AudienceLevel, options: audience_levels }) %&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;%= f.TextArea({name: "Description", value: talk.Description, rows: 10}) %&gt;
    &lt;/div&gt;
    &lt;div class="col-md-12"&gt;
      &lt;%= f.TextArea({notes:"Notes", value: talk.Notes, rows: 10 }) %&gt;
    &lt;/div&gt;

  &lt;/div&gt;
&lt;% } %&gt;
```

```html
// OUTPUT
&lt;form action="/talks/3" method="POST"&gt;
  &lt;input name="authenticity_token" type="hidden" value="e0c536b7a1a7d752066727b771f1e5d02220ceff5143f6c77b"&gt;
  &lt;input name="_method" type="hidden" value="PUT"&gt;
  &lt;div class="row"&gt;
    &lt;div class="col-md-12"&gt;
      &lt;div class="form-group"&gt;
        &lt;input class=" form-control" name="Title" type="text" value="My Title"&gt;
      &lt;/div&gt;
    &lt;/div&gt;
    &lt;div class="col-md-6"&gt;
      &lt;div class="form-group"&gt;
        &lt;textarea class=" form-control"&gt;some data here&lt;/textarea&gt;
      &lt;/div&gt;
    &lt;/div&gt;

    &lt;div class="col-md-6"&gt;
      &lt;div class="form-group"&gt;
        &lt;select class=" form-control" name="TalkFormatID"&gt;
          &lt;option value="0" selected&gt;Talk&lt;/option&gt;
          &lt;option value="1"&gt;Lightning Talk&lt;/option&gt;
          &lt;option value="2"&gt;Workshop&lt;/option&gt;
          &lt;option value="3"&gt;Other&lt;/option&gt;
        &lt;/select&gt;
      &lt;/div&gt;
      &lt;div class="form-group"&gt;
        &lt;select class=" form-control" name="AudienceLevel"&gt;
          &lt;option value="All" selected&gt;All&lt;/option&gt;
          &lt;option value="Beginner"&gt;Beginner&lt;/option&gt;
          &lt;option value="Intermediate"&gt;Intermediate&lt;/option&gt;
          &lt;option value="Advanced"&gt;Advanced&lt;/option&gt;
        &lt;/select&gt;
      &lt;/div&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;div class="form-group"&gt;
        &lt;textarea class=" form-control" name="Description" rows="10"&gt;some data here&lt;/textarea&gt;
      &lt;/div&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;div class="form-group"&gt;
        &lt;textarea class=" form-control" notes="Notes" rows="10"&gt;some data here&lt;/textarea&gt;
      &lt;/div&gt;
    &lt;/div&gt;
  &lt;/div&gt;
&lt;/form&gt;
```
<% } %>

