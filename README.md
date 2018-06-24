aerie
=====

Aerie -- shortened to `ae` -- is a tool for developers, used for overviewing the status and recent changes in projects.

"Projects" is defined to mean "git repositories".  Aerie also has some features which integrate lightly with some popular git hosts, but does not force you to use these things; we :heart: decentralization.

Aerie is an overview tool only -- it's read-only views.  You'll still need to use git normally or with other tools; aerie just points you towards things that need some love.


Roadmap
-------

Aerie is currently a pre-alpha prototype.  Use at your own risk.  (That said, it *is* read-only, so the "risk" is... well, none.)

Here are some of the things we want to cover:

- List repositories.
- List the date of the last change on the 'master' branch in each repository.
- List how many non-'master' branches are in each repository.
- List the date of the last change on each branch; contextualize this into how long this branch has been diverging.
- List the date of the last change on each branch; contextualize this into how stale it has become.
- List how many commits diverged from 'master' each branch is.
- Provide notice when remote branches are updated and local references have not.
- more to come...

Generally, these are all features aimed at making sure that work which has been started doesn't sit around unfinished and become stale.
