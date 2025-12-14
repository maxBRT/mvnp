# Project Plan: mvnp (Maven Plus)

## 1. Project Philosophy

* **Single Binary:** Fast startup, native compilation (no JVM warmup lag).
* **Convention over Configuration:** Sane defaults (e.g., latest Java version, JUnit 5, Git initialized) to minimize user prompts.
* **Preservation:** Respect the user's existing `pom.xml` structure. Never reformat the entire file; use string insertion to preserve comments and indentation.

---

## 2. Architecture Overview

The tool sits between the User and the underlying Maven binary/File System.

1.  **CLI Layer:** Parses commands (`init`, `get`, `make`) and handles flags.
2.  **Orchestration:** Delegate to the installed `mvn` binary and perform IO directly when possible.
3.  **Maven Central Client:** Interacts with the Maven Central search API to fetch dependency data.
4.  **XML Engine:** A specialized module to edit `pom.xml` as text/strings rather than full DOM parsing, ensuring format safety.

---

## 3. Core Functionalities

### Feature A: Scaffolding (`mvnx init`)
**Goal:** Replace `mvn archetype:generate` with a fast, zero-network template engine.

* **Logic:**
    1.  Create standard directory structure: `src/main/java`, `src/test/java`, `src/main/resources`.
    2.  Generate a `pom.xml` from an internal Handlebars/Go template.
    3.  Pre-fill modern plugins (`maven-compiler-plugin`, `maven-surefire-plugin`).
    4.  Initialize a Git repository (`git init`) and generate a Java-specific `.gitignore`.

### Feature B: Dependency Injection (`mvnx get`)
**Goal:** Eliminate manual GAV (Group, Artifact, Version) lookups and XML copy-pasting.

* **Logic:**
    1.  **Search:** Query `https://search.maven.org` with the user's input.
    2.  **Select:** Display an interactive TUI (Text User Interface) list for the user to pick the correct artifact.
    3.  **Fetch:** Retrieve the latest stable version.
    4.  **Inject:** Locate the closing `</dependencies>` tag in `pom.xml` and insert the `<dependency>` block immediately before it.

### Feature C: Code Generation (`mvnx make`)
**Goal:** Simplify the creation of boilerplate Java files.

* **Usage:** `mvnx make class users.UserService`
* **Logic:**
    1.  Parse `pom.xml` to locate `src/main/java`.
    2.  Infer the root package from the directory structure.
    3.  Create the file `src/main/java/com/app/users/UserService.java`.
    4.  Write the file with the correct `package com.app.users;` declaration.

---



## 4. Detailed Logic: The `get` Command

1.  **User Input:** `mvnx get jackson`
2.  **API Call:** GET `https://search.maven.org/solrsearch/select?q=jackson&rows=5&wt=json`
3.  **TUI Display:**
    ```text
    ? Select a dependency:
    > com.fasterxml.jackson.core:jackson-databind (2.15.2)
      com.fasterxml.jackson.core:jackson-core (2.15.2)
    ```
4.  **Selection:** User presses Enter.
5.  **File Operation:**
    * Read `pom.xml` as a string.
    * Check if `<dependencies>` block exists.
    * **If Yes:** String replace `</dependencies>` with:
        ```xml
            <dependency>
                <groupId>com.fasterxml.jackson.core</groupId>
                <artifactId>jackson-databind</artifactId>
                <version>2.15.2</version>
            </dependency>
        </dependencies>
        ```
    * **If No:** Create the block and insert.

---

## 5. Development Roadmap

### Phase 1: The Skeleton
* Initialize Go module.
* Set up Cobra for `init`, `get` and `make` commands.
* Implement `init` to use mvn archetype.
* Integrate Bubble Tea to query user for group and artefact.

### Phase 2: The Search Engine
* Implement the HTTP Client for Maven Central.
* Integrate Bubble Tea to display search results interactively.

### Phase 3: The XML Surgeon
* Implement the file reading/writing logic.
* Build the "safe injection" function (finding the tag and inserting text).

### Phase 4: Polish & Distribution
* Add the `g` (generate) command.
* Set up GitHub Actions to compile binaries for Windows, Linux, and Mac.

---

## 7. Suggested Directory Structure

```text
mvnx/
├── cmd/
│   ├── root.go       # Main entry point
│   ├── init.go       # Logic for 'init' command
│   ├── add.go        # Logic for 'add' command
│   └── generate.go   # Logic for 'g' command
├── internal/
│   ├── maven/        # Maven Central API client
│   ├── scaffold/     # Templates for new projects
│   ├── xml/          # Logic for editing pom.xml
│   └── tui/          # Bubble Tea UI components
├── main.go           # Application runner
├── go.mod
└── go.sum
