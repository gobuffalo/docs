<%= title("Writing Non-Go Plugins") %>

Plugins do not need to be written in Go. They can be written in any language you would like, as long as they comply with the rules above.

For example, we can write the following plugin using Ruby:

```ruby
#!/usr/bin/env ruby
# ./plugins/buffalo-hello.rb

require 'json'

command = ARGV[0]

case command
when 'available'
  puts JSON.generate([{ name: 'hello', buffalo_command: 'root', description: 'says hello to you' }])
when 'hello'
  puts 'Hi there!'

end
```

To activate the plugin we need to add the file as `buffalo-hello.rb` to somewhere in the `$PATH` or in a directory called `plugins/` inside of a Buffalo application.

Finally the file needs to be made executable. On a Mac/Linux it can be done with `chmod +x buffalo-hello.rb`.

```bash
$ buffalo plugins list

Bin              |Command                    |Description
---              |---                        |---
buffalo-auth     |buffalo generate auth      |Generates a full auth implementation
buffalo-pop      |buffalo db                 |[DEPRECATED] please use `buffalo pop` instead.
buffalo-goth     |buffalo generate goth-auth |Generates a full auth implementation use Goth
buffalo-goth     |buffalo generate goth      |generates a actions/auth.go file configured to the specified providers.
buffalo-hello.rb |buffalo hello              |says hello to you
buffalo-heroku   |buffalo heroku             |helps with heroku setup and deployment for buffalo applications
buffalo-plugins  |buffalo events listen      |listens to github.com/gobuffalo/events
buffalo-pop      |buffalo destroy model      |Destroys model files.
buffalo-plugins  |buffalo generate plugin    |generates a new buffalo plugin
buffalo-plugins  |buffalo plugins            |tools for working with buffalo plugins
buffalo-pop      |buffalo pop                |A tasty treat for all your database needs
buffalo-trash    |buffalo trash              |destroys and recreates a buffalo app
buffalo-upgradex |buffalo upgradex           |updates Buffalo and/or Pop/Soda as well as your app
```
