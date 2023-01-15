# 🎣 PhishReport's IOK contribution guidelines

Thank you for taking your time to contribute to this project, we appreciate it a lot 🎉

The following document contains several guidelines that you should follow when contributing to this project. 
If you believe the guidelines require alteration please propose your own changes too via a Pull Request!

## Styleguides

### Pull Requests

- Please follow the format below to help speed up the process of reviewing PRs

```
🟢Additions:

- Added [`rule-name-lowercase`](<COMMIT_URL>)

🟠Changes:

- Added rule condition to [`rule-name-lowercase`](<COMMIT_URL>)

🔴Removals:

- Removed [`rule-name-lowercase`](<COMMIT_URL>)

```

### Commit Messages

- Use the present tense
- Reference related issues & pull requests in the description (*if applicable*)
- Limit the title to 60 characters at most
- It is advised to prepend the following emojis to the start of a commit message:
  - 💎 `:gem:` when changing underlying IOK code
  - 💡 `:bulb:` when updating dependencies 
  - ✂️ `:scissors:` when removing dependencies
  - ✨ `:sparkles:` when modifying an IOK rule
  - 🚀 `:rocket:` when creating a new IOK rule
  - 📦 `:package:` when modifying the CI workflows
  - 📌 `:pushpin:` when modifying the guidelines outlined in this file (`CONTRIBUTING.md`)

### IOK Rules

- Titles must follow the format of:
  - `<context> Phishing Kit <unique_id>`
    - `context` is either the brand or process being imitated (e.g. `Facebook Account Recovery`)
    - `unique_id` can be generated from a random assortment of `8` alphanumeric characters in the ranges of `[0-9]` and `[a-f]` (e.g. `0e420f8e`) and must be written in <ins>lowercase</ins>

- Descriptions must at the least accurately describe what the rule detects (e.g. `Detects a Facebook phishing kit, telling the victim to enter their details to reactivate their account.`)
- References must include between to `2` to `5` unique URLScan URLs which refer to the same kit
- All detection fields must follow camel casing (e.g. `camelCase`)
- Tags must include the targeted company/brand <ins>**OR**</ins> technique(s) used, and at the very least should include targeted country (*if applicable*) and any other tags that you deem to be sufficient (eg. `kit`, `target.facebook`, `target_country.germany`)

**Rule Structure:**
```yaml
title: 
description: 
references:
[related:]
detection:
  
  fieldName:
  
  condition: 
tags:
```

**NOTICE:** Descriptions do in fact support markdown rendering, when being displayed on the rule's dedicated page on the PhishReport website [here](https://phish.report/IOK/indicators)
