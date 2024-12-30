# Getting Started with Go and the Web: Deploy to Railway

![Cover](/doc/cover.png)

This guide aims to show you how to build a web app using Go. You will create and test a simple web server built using Go on your local machine then deploy your web app to Railway.

Railway is a hosting service you can use to deploy your apps. It supports all sorts of deployment configs and has a plethora of features to customize the hosting set up for your app.

## Before We Start

To follow along with the guide, you must have Go installed on your machine. Download and install Go here: [Download and Install Go]().

Git and the GitHub CLI should be installed on your machine. Download and install Git from here: [Downloads - Git](https://git-scm.com/downloads). Download and install the GitHub CLI from here: [Download and Install GitHub CLI](https://cli.github.com).

A GitHub account and a Railway account are required. Sign up for GitHub here: [Create GitHub account](https://github.com/join). And sign up for Railway here: [Create Railway account](https://railway.com/login)

You must have some familiarity with HTML, CSS, JavaScript, HTTP to keep up with the concepts discussed in the guide.

Lastly, knowledge of Go is useful but not mandatory. This is a beginner guide.

## Hello, World! Go

Open your work directory and create a new file and name it `main.go`. 

Add the following code to `main.go`:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

Save the file and test it by running the following command in your terminal in your working directory:
```sh
go run main.go
```

You should the text "Hello, World!" in your terminal output. 

## Hello, World! Go on the Web

We just created a simple CLI app using Go. Next, let's create a simple web server using Go. Comment out the code you just wrote in the previous step. Above the comment block, add the following code:

```go
package main

import "fmt"
import "net/http"

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, World!</h1>")
  })

  http.ListenAndServe(":80", nil)
}
```
Save the file and run it. Visit `localhost` in your browser and you should see the text "Hello, World!" on a webpage.

![Hello World web page](https://res.cloudinary.com/craigsims808/image/upload/v1735436523/freelance/sturdy-train/hello-world-web-page_vrhx5b.png)

## Create and Upload GitHub Repo

Initialize your project directory as a Git repo.
```sh
git init
```

Commit your project files. In this case it's just the `main.go` file.
```sh
git add main.go
```

```sh
git commit -m "Initial commit"
```

Authenticate with GitHub
```sh
gh auth login
```

Upload your repo to GitHub
```sh
gh repo create my-repo --public --source=. --remote=origin
```

```sh
git push --set-upstream origin master
```

Replace `my-repo` with your desired repo name.

## Deploy to Railway

Railway offers many ways to deploy your web app. You can use the dashboard, the CLI, the API etc. This article: [Railway Deployment Options](https://docs.railway.com/quick-start) explains all the numerous ways you can deploy an app on the Railway platform.

We will use the dashboard for a start.

### Create a new Railway project

Visit dev.new in your browser. This will redirect you to railway.com/new and you will see a **New Project** modal with deployment options.

![Railway New Project Modal](https://res.cloudinary.com/craigsims808/image/upload/v1735561144/freelance/falcon-feather/railway-create-new-project-dashboard_x3f2le.png)

Select **Deploy from GitHub repo** and choose the repo you created previously.

![Deploy From GitHub repo](https://res.cloudinary.com/craigsims808/image/upload/v1735561971/freelance/falcon-feather/deploy-from-github-repo_wv8yvf.png)

The Railway platform will read the contents of your repo, initialize the project, build and then deploy it automatically as a service.

![Railway Build successful](https://res.cloudinary.com/craigsims808/image/upload/v1735586216/freelance/falcon-feather/railway-initial-build_ibv604.png)

### Generate a Domain for your project

Select **Settings** inside your Railway project's service. Under **Networking** click **Generate Domain**. This allows you to access your service on the internet.

![Generate Domain](https://res.cloudinary.com/craigsims808/image/upload/v1735586216/freelance/falcon-feather/generate-domain_mdpf51.png)

Railway will generate a domain name for your app. The URL will appear a few seconds after clicking **Generate Domain**. You will use this URL to access your Go web app on the internet. 

![View Generated Domain](https://res.cloudinary.com/craigsims808/image/upload/v1735586216/freelance/falcon-feather/view-generated-domain_epgsxm.png)

Test out your newly generated domain name in your browser. You should see the text "Hello, World!" when you visit the link.

![Test Generated Domain](https://res.cloudinary.com/craigsims808/image/upload/v1735586215/freelance/falcon-feather/test-generated-domain_jpjcrj.png)

## Serve Static files using Go

In your local project folder, create a new folder, `static`. Create a new file named `index.html` inside `static` and add the following code:

```html
<body>
  <h1>My Static Website</h1>
</body>
```

Update the `main.go` file to serve static files from the `static` folder. Comment out all the code you have written. Above the comment block, add the following code to `main.go`:

```go
package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":80", nil)
}
```

Test your code:
```sh
go run main.go
```

Visit `localhost` in your browser and you should see a webpage with the text: "My Static Website"

![My Static Website Local](https://res.cloudinary.com/craigsims808/image/upload/v1735587731/freelance/falcon-feather/my-static-website-local_mlaruy.png)

## Update Deployment on Railway

To update your deployment, simply stage and commit the file changes you have made using Git.

```sh
git add .
```

```sh
git commit -m "App update"
```

Then push the updates to GitHub. 

```sh
git push
```

Railway will automatically update your project right after pushing your changes to GitHub. Wait a few seconds and visit your app URL to test your changes

![Test Updated Site Live](https://res.cloudinary.com/craigsims808/image/upload/v1735588268/freelance/falcon-feather/test-updated-site-live_eynli9.png)

## Conclusion

Congratulations! You have successfully built a simple web app using Go and deployed it to Railway. You learned how to create a basic web server, serve static files, and utilize GitHub for version control and deployment. Railway makes it easy to deploy and manage your applications with its user-friendly platform and powerful features. Keep experimenting with Go and Railway to build more complex and scalable web applications. Happy coding!

## References

1. [Download and Install Go](https://golang.org/dl/)
2. [Downloads - Git](https://git-scm.com/downloads)
3. [Download and Install GitHub CLI](https://cli.github.com)
4. [Create GitHub account](https://github.com/join)
5. [Create Railway account](https://railway.com/login)
6. [Railway Deployment Options](https://docs.railway.com/quick-start)