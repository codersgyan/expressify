### Expressify-CLI 🚀<br/>

<img src="./assets/expressify-logo.png" height="200px" style="border-radius:50px"/><br/>

**Expressify-CLI** is a command-line tool designed to generate a **production-grade** scaffold for Express applications in a single command. This tool streamlines the setup process for developing Express applications with JavaScript/TypeScript, making it faster and more efficient.

#### Inspiration 💡

Setting up a new Express project can often feel like reinventing the wheel. Developers frequently encounter the challenge of integrating a myriad of components such as TypeScript configurations, testing frameworks, linters, Docker configurations, and more. This not only consumes valuable time but also poses a risk of inconsistency and configuration errors.

#### How Expressify-CLI Solves the Problem 🛠️

Expressify-CLI tackles these challenges head-on by automating the creation of a comprehensive and robust scaffold for Express applications. With just one command, developers can now generate a project structure that includes:

-   TypeScript (TS) configuration for robust typing. 📐
-   Test configurations for reliable code testing. ✅
-   Linters for maintaining code quality and consistency. 🔍
-   Docker files for easy containerization. 🐳
-   Logger setup for effective logging and monitoring. 📝
-   Graceful shutdowns for better resource management. 🧘
-   Error handling middleware for improved error management. 🚫
-   Optional authentication module with JWT for secure access. 🔒
-   Pre-configured Git setup with README, license file, and other production-level settings. 📄

#### Run Locally (Development) 🚀

Running Expressify-CLI locally is straightforward with the help of the provided Makefile.
Follow these steps to set up and run the project in a development environment:

##### Clone the Repository

Clone the repository to your local machine using Git:

```bash
git clone https://github.com/codersgyan/expressify.git
cd expressify
```

##### Using the Makefile

The Makefile includes commands to simplify the build and test process. Here's how you can use it:

###### Build the Project:

Compiles the project and creates an executable.

```bash
make build
```

###### Run the Tool:

Executes the compiled binary.

```bash
make run
```

###### Run Tests:

Runs the automated tests.

```bash
make test
```

###### Clean Up:

Removes the generated binary and any other temporary files.

```bash
make clean
```

##### Prerequisites:

Ensure that Go is installed on your system. Verify this with go version.
Familiarity with basic make commands can be helpful.

**_Notes:_**
The Makefile simplifies common tasks but can be modified if your workflow requires it.
Additional configuration or steps might be necessary depending on the specific setup of your project.

#### Open Source Contribution 🤝

Expressify-CLI is an open source project, and contributions are greatly appreciated! If you have ideas for improvements or have found a bug, feel free to open an issue. We also warmly welcome pull requests. Let's build a stronger tool together! 🌍🛠️
