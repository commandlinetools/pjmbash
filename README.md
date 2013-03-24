pjmbash
=======

Project Management - Bash
-------------------------

> A command line tool for viewing, changing, and quickly switching between different development and runtime environments.  Manages the environment variables on a per project basis with the ability to group/filter similar projects using tags.  Inspects the base environment and shows differences per project.  Categorizes commonly used tools by name, version, os, user, and location.

Notes:

1. pjmbash pre-beta development currently being done on the develop branch.
1. pjmwin is being developed in parallel for Windows environments

Status
------

* current work being done on feat-lspj branch
* next action: test lspj

Plan
----

* lspj - list no projects
* shpj - show nonexistent project
* adpj - add first project - show first project, list one project
* adpj - add second project - show second project, list two projects
* chpj - change to first project, show current project, list two projects marking first as current
* chpj - change to second project, show current project, list two projects marking second as current
* shpj - show default environment, show delta from default
* cmpj - compare two project environments
* misc - work on managing project dependencies

        > e.g., ant, dcevm, idea, java, mvn, tomcat (nickname, vendor, name, type, version, os, user, location, url)

Future
------

 * edpj - edit a project, add a tag, change a tag, remove a tag, list with filter to include or exclude tag
 * rmpj - remove a project, show removed (nonexistent) project, list no longer shows removed project
 * pspj - push the current project onto a stack, list projects shows stacked projects
 * pppj - pop the top project off of the stack, list projects shows empty stack
 * expj - export a project to a script or data file (bash, windows cmd, json)
 * impj - import a project from a script or data file (bash, windows cmd, json)
