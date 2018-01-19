<%= title("Model Forms") %>

The `form_for` helper can be used to generate HTML forms for a specified model. This makes the code easier to write, and maintains a level of "consistency" across your application.

The `form_for` helper behaves in a similar matter to the `form` helper, with several key differences.

The first difference is that the `form_for` takes a "model" as a first argument. This "model" only needs to be a `struct` it does not have to be database backed.

The second difference is in the tag calls themselves. These tags, such as `InputTag`, take the name of the attribute on the model you want to build a field for, then they take an optional set of options as the second argument.

<%= codeTabs() { %>
```go
// models/talk.go
type Talk struct {
  ID            int          `json:"id" db:"id"`
  CreatedAt     time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
  UserID        int          `json:"user_id" db:"user_id"`
  Title         string       `json:"title" db:"title"`
  Description   nulls.String `json:"description" db:"description"`
  Notes         nulls.String `json:"notes" db:"notes"`
  ParentID      nulls.Int    `json:"parent_id" db:"parent_id"`
  Abstract      string       `json:"abstract" db:"abstract"`
  AudienceLevel string       `json:"audience_level" db:"audience_level"`
  IsPublic      nulls.Bool   `json:"is_public" db:"is_public"`
  TalkFormatID  int          `json:"talk_format_id" db:"talk_format_id"`
}
```

```erb
// templates/talks/edit.html
&lt;%= form_for( talk, {action: talkPath({id: 3}), method: "PUT"}) { %&gt;
  &lt;div class="row"&gt;
    &lt;div class="col-md-12"&gt;
      &lt;%= f.InputTag("Title") %&gt;
    &lt;/div&gt;
    &lt;div class="col-md-6"&gt;
      &lt;%= f.TextArea("Abstract", {hide_label: true}) %&gt;
    &lt;/div&gt;


    &lt;div class="col-md-6"&gt;
      &lt;%= f.SelectTag("TalkFormatID", {options: talk_formats}) %&gt;
      &lt;%= f.SelectTag("AudienceLevel", , {options: audience_levels}) %&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;%= f.TextArea("Description", {rows: 10}) %&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;%= f.TextArea("Notes", {rows: 10}) %&gt;
    &lt;/div&gt;
  &lt;/div&gt;
&lt;% } %&gt;
```

```html
// OUTPUT
&lt;form action="/talks/3" id="talk-form" method="POST"&gt;
  &lt;input name="authenticity_token" type="hidden" value="cd998be98a99b452481c43fd3e4715e4e85333a45b982ac999"&gt;
  &lt;input name="_method" type="hidden" value="PUT"&gt;
  &lt;div class="row"&gt;
    &lt;div class="col-md-12"&gt;
      &lt;div class="form-group"&gt;
        &lt;label&gt;Title&lt;/label&gt;
        &lt;input class="form-control" id="talk-Title" name="Title" type="text" value="My Title"&gt;
      &lt;/div&gt;
    &lt;/div&gt;
    &lt;div class="col-md-6"&gt;
      &lt;div class="form-group"&gt;
        &lt;textarea class="form-control" id="talk-Abstract" name="Abstract"&gt;some data here&lt;/textarea&gt;
      &lt;/div&gt;
    &lt;/div&gt;

    &lt;div class="col-md-6"&gt;
      &lt;div class="form-group"&gt;
      &lt;label&gt;TalkFormatID&lt;/label&gt;
        &lt;select class="form-control" id="talk-TalkFormatID" name="TalkFormatID"&gt;
          &lt;option value="0" selected&gt;Talk&lt;/option&gt;
          &lt;option value="1"&gt;Lightning Talk&lt;/option&gt;
          &lt;option value="2"&gt;Workshop&lt;/option&gt;
          &lt;option value="3"&gt;Other&lt;/option&gt;
        &lt;/select&gt;
      &lt;/div&gt;
      &lt;div class="form-group"&gt;
        &lt;label&gt;AudienceLevel&lt;/label&gt;
        &lt;select class=" form-control" id="talk-AudienceLevel" name="AudienceLevel"&gt;
          &lt;option value="All" selected&gt;All&lt;/option&gt;
          &lt;option value="Beginner"&gt;Beginner&lt;/option&gt;
          &lt;option value="Intermediate"&gt;Intermediate&lt;/option&gt;
          &lt;option value="Advanced"&gt;Advanced&lt;/option&gt;
        &lt;/select&gt;
      &lt;/div&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;div class="form-group"&gt;
        &lt;label&gt;Description&lt;/label&gt;
        &lt;textarea class=" form-control" id="talk-Description" name="Description" rows="10"&gt;some data here&lt;/textarea&gt;
      &lt;/div&gt;
    &lt;/div&gt;

    &lt;div class="col-md-12"&gt;
      &lt;div class="form-group"&gt;
        &lt;label&gt;Notes&lt;/label&gt;
        &lt;textarea class=" form-control" id="talk-Notes" name="Notes" rows="10"&gt;some data here&lt;/textarea&gt;
      &lt;/div&gt;
    &lt;/div&gt;
  &lt;/div&gt;
&lt;/form&gt;
```
<% } %>

