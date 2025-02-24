# Contributing to semantic-release

As a contributor, here are the guidelines we would like you to follow:

- [How can I contribute?](#how-can-i-contribute)
- [Using the issue tracker](#using-the-issue-tracker)
- [Submitting a Pull Request](#submitting-a-pull-request)
- [Coding rules](#coding-rules)
- [Git Commit Type](#git-commit-type)
- [Working with the code](#working-with-the-code)

## How can I contribute?

### Give feedback on issues

Some issues are created without information requested in the [Bug report guideline](#bug-report).
Help make them easier to resolve by adding any relevant information.

Issues with the [discussion label](https://github.com/Blvckleg/BingusBoingus/labels/discussion) are meant to discuss the implementation of new features.

### Fix bugs and implement features

Confirmed bugs and ready-to-implement features are marked with the [help wanted label](https://github.com/Blvckleg/BingusBoingus/labels/help%20wanted) or the [bug label](https://github.com/Blvckleg/BingusBoingus/labels/bug).

## Using the issue tracker

The issue tracker is the channel for [bug reports](#bug-report), [features requests](#feature-request) and [submitting pull requests](#submitting-a-pull-request) only.

### Bug report

A good bug report shouldn't leave others needing to chase you for more information.
Please try to be as detailed as possible in your report and fill the information requested in the [bug report template](https://github.com/Blvckleg/BingusBoingus/blob/master/.github/ISSUE_TEMPLATE/bug_report.md).

### Feature request

Please provide as much detail and context as possible and fill the information requested in the [feature request template](https://github.com/Blvckleg/BingusBoingus/blob/master/.github/ISSUE_TEMPLATE/feature_request.md).

## Submitting a Pull Request

Good pull requests, whether patches, improvements, or new features, are a fantastic help.
They should remain focused in scope and avoid containing unrelated commits.

Here is a summary of the steps to follow:

1. Create issue

2. Create issue branch from main
3. Make your code changes, following the [Coding rules](#coding-rules)
4. Push your topic branch

```bash
git push origin <topic-branch-name>
```

5. [Open a Pull Request](https://help.github.com/articles/creating-a-pull-request/#creating-the-pull-request) with a clear title and description.

**Tips**:

- For ambitious tasks, open a Pull Request as soon as possible with the `[WIP]` prefix in the title, in order to get feedback and help from the community.


## Coding rules

### Source code

To ensure consistency and quality throughout the source code, all code modifications must have:

- No [linting](#lint) errors
- A [test](#tests) for every possible case introduced by your code change
- [Valid commit message(s)](#commit-message-guidelines)
- Documentation for new features
- Updated documentation for modified features

### Documentation

To ensure consistency and quality, all documentation modifications must:

- Refer to brand in [bold](https://help.github.com/articles/basic-writing-and-formatting-syntax/#styling-text) with proper capitalization, i.e. **GitHub**, **Nestjs**, **npm**
- Prefer [tables](https://help.github.com/articles/organizing-information-with-tables) over [lists](https://help.github.com/articles/basic-writing-and-formatting-syntax/#lists) when listing key values, i.e. List of options with their description
- Use [links](https://help.github.com/articles/basic-writing-and-formatting-syntax/#links) when you are referring to:
- Use the [single backtick `code` quoting](https://help.github.com/articles/basic-writing-and-formatting-syntax/#quoting-code) for:
  - programming language keywords, i.e. `function`, `async`, `String`
- Use the [triple backtick `code` formatting](https://help.github.com/articles/creating-and-highlighting-code-blocks) for:
  - code examples
  - configuration examples
  - sequence of command lines

### Commit message guidelines

#### Atomic commits

If possible, make [atomic commits](https://en.wikipedia.org/wiki/Atomic_commit), which means:

- a commit should contain exactly one self-contained functional change
- a functional change should be contained in exactly one commit
- a commit should not create an inconsistent state (such as test errors, linting errors, partial fix, feature with documentation etc...)

A complex feature can be broken down into multiple commits as long as each one maintains a consistent state and consists of a self-contained change.

#### Commit message format

Each commit message consists of a **header**, a **body** and a **footer**.
The header has a special format that includes a **type**, a **scope** and a **subject**:

```commit
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

The **header** is mandatory and the **scope** of the header is optional.

The **footer** can contain a [closing reference to an issue](https://help.github.com/articles/closing-issues-via-commit-messages).

#### Revert

If the commit reverts a previous commit, it should begin with `revert: `, followed by the header of the reverted commit.
In the body it should say: `This reverts commit <hash>.`, where the hash is the SHA of the commit being reverted.

## Git Commit Type

The type must be one of the following:

| Type         | Description                                                                                                 |
| ------------ | ----------------------------------------------------------------------------------------------------------- |
| **build**    | Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)         |
| **ci**       | Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs) |
| **docs**     | Documentation only changes                                                                                  |
| **feat**     | A new feature                                                                                               |
| **fix**      | A bug fix                                                                                                   |
| **perf**     | A code change that improves performance                                                                     |
| **refactor** | A code change that neither fixes a bug nor adds a feature                                                   |
| **style**    | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)      |
| **test**     | Adding missing tests or correcting existing tests                                                           |

#### Subject

The subject contains succinct description of the change:

- use the imperative, present tense: "change" not "changed" nor "changes"
- don't capitalize first letter
- no dot (.) at the end

#### Body

Just as in the **subject**, use the imperative, present tense: "change" not "changed" nor "changes".
The body should include the motivation for the change and contrast this with previous behavior.

#### Footer

The footer should contain any information about **Breaking Changes** and is also the place to reference GitHub issues that this commit **Closes**.

**Breaking Changes** should start with the word `BREAKING CHANGE:` with a space or two newlines.
The rest of the commit message is then used for this.

#### Examples

```commit
fix(pencil): stop graphite breaking when too much pressure applied
```

```commit
feat(pencil): add 'graphiteWidth' option

Fix #42
```

```commit
perf(pencil): remove graphiteWidth option

BREAKING CHANGE: The graphiteWidth option has been removed.

The default graphite width of 10mm is always used for performance reasons.
```

## Working with the code

### Set up the workspace

```bash
# Clone your fork of the repo into the current directory
$ git clone https://github.com/
# Navigate to the newly cloned directory
$ cd <repo-name>
# Assign the original repo to a remote called "upstream"
wip
```

### Lint


### Tests

Before pushing your code changes make sure all **tests pass** and the **coverage is good**: ``npm run test``

### Commits
