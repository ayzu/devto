# Devto Client

This application reads your article and publishes or edits it in `dev.to` portal.

## File Format

The application expects an article in the following format. 

```markdown
    # Header
    
    ## Meta
    
    tags: go, programming

    ## Part 1, Intro
```

The level one header (`#`) will be used at a title for the article. Information in section `## Meta` represents a meta information about your article and won't be included into its text. `Tags` represent a list of tags of the article.

## Usage

Publish an article:

    devto  --key=$KEY $FILEPATH

where

    $KEY - is API Key for dev.to portal
    $FILEPATH -is filepath to the article.