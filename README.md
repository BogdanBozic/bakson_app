## bakson_app

![CodeCheck](https://github.com/BogdanBozic/bakson_app/actions/workflows/checkcode.yaml/badge.svg)


This simple GO gin webapp is a part of the interview assignment for Bakson ltd. It has two endpoints, /hello and /time that return a JSON response and date and time on the server as a string on a GET request without any authentication needed.

The Infrastructure repo is located [here](https://github.com/BogdanBozic/bakson_infra).

### GO Code

Go code is rather simple. Written with gin, it consists of a single module only and it exposes two endoints. Code is "covered" with unittests, rather clumsy ones, but unittests nonetheless. The tests are run as part of the CI/CD.

### Dockerfile

Dockerfile is, well, Dockerfile. It prepares the code to build an image out of it that will be used for running the app on K8s.

### GitHub Actions

Repo uses GitHub Actions for CI/CD. There are two workflows:

- CheckCode: runs GO and Dockerfile lints, and GO unittests in parallel. Runs on all pushes, except pushes to `main` branch.
- Deploy: builds the container image, uploads it to AWS ECR and updates the infra repo with the version of the app. Runs only on push to `main` branch.

### Branches and Environments

I haven't built separate environments or thought of any deployment strategies: everything that is pushed to the `main` branch will be deployed to the only environment that exists and will be reflected straight away. This approach, of course, I would never recommend to anyone that has a project running in production.


