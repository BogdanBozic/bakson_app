
## bakson_app

![CodeCheck](https://github.com/BogdanBozic/bakson_app/actions/workflows/checkcode.yaml/badge.svg) ![Deploy](https://github.com/BogdanBozic/bakson_app/actions/workflows/deploy.yaml/badge.svg)


This simple GO gin webapp is a part of the interview assignment for Bakson ltd. It has two endpoints, [bakson-app.bastovansurcinski.click/](bakson-app.bastovansurcinski.click/) and [bakson-app.bastovansurcinski.click/time](bakson-app.bastovansurcinski.click/time) that return a JSON response and date and time UTC as a string on a GET request without any authentication needed.

The Infrastructure repo is located [here](https://github.com/BogdanBozic/bakson_infra). Please, take a look at it as it contains much more logic and resources than this repo, even though it was said to prioritize the Go app.

### GO Code

Go code is rather simple. Written with gin, it consists of a single module only, and it exposes two endpoints. Code is "covered" with unittests, rather clumsy ones, but unittests nonetheless. 
The tests are run as part of the CI/CD.

### Dockerfile

Dockerfile is, well, Dockerfile. It prepares the code to build an image out of it that will be used for running the app on K8s.
I have seen examples of multi-stage builds for Go apps where building the code is in one, and running the executable is in the other stage (or the last stage, anyway). I haven't tried this simply because I spent too much time building the infrastructure.

### GitHub Actions

Repo uses GitHub Actions for CI/CD. There are two workflows:

- CheckCode: runs GO and Dockerfile lints, and GO unittests in parallel. Runs on all pushes, except pushes to `main` branch.
- Deploy: builds the container image, uploads it to AWS ECR and updates the infra repo with the version of the app. Runs only on push to `main` branch.

### Branches and Environments

I haven't built separate environments or thought of any deployment strategies: everything that is pushed to the `main` branch will be deployed to the only environment that exists and will be reflected straight away.
This approach, of course, I would never recommend to anyone that has a project running in production.
