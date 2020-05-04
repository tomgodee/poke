
### Heroku
To setup, follow this guide.
https://devcenter.heroku.com/articles/getting-started-with-go#set-up
First, install the heroku's CLI
Once finished log in (i've logged in the web so it's automatically logged in for me)

Since i've created this app b4 i followed this guide next
https://devcenter.heroku.com/articles/preparing-a-codebase-for-heroku-deployment

First, i created a new empty application on Heroku, along with an associated empty Git repository.
```
heroku create
```

Then check if the remote i've just created is there 
```
git remote -v
```
You can also rename the remote (ik i did, to "poke-heroku")
```
git remote rename heroku poke-heroku
```

To deploy we need to push to heroku's remote master branch
```
git push [remote_name] [branch_name]

git push poke-heroku master
```

Heroku recommends us to add a Procfile



To push our local db to the heroku's remote db we need to set up like the following
```
SET PGUSER=[your_db_username] => in my case it's postgres
SET PGPASSWORD=[your_db_password] => in my case it's zxc321
```
Then to push
```
heroku pg:push [your_db_name] [remote_db_name] --app [your_app_name]
// For me:
// [your_db_name] = poke-development
// [remote_db_name] = DATABASE_URL
// [your_app_name] = postgresql-tapered-69911 
```

We shouldn't commit the ```.env``` file since this file contains config vars that we use on the local enviroment. On deployed enviroment we use other config vars. We can look at the heroku's .env file by typing
```
heroku config
```
We can also set and delete config vars through either the CLI or GUI
To read more about this: https://devcenter.heroku.com/articles/config-vars