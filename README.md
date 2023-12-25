# go-ci-release-repository

Simple template of go repository with workflow for ci and release

## Usage

In ``main`` branch contains base repository without workflow of release artifacts and docs.

In ``docs`` branch add workflow for generating mkdocs.

In ``artifacts`` branch add workflow for generating release with artifacts as binaries.

You can merge needed branches to main to add new workflow functional.

Version expected start with ``v[n].[n].[n]*``, where ``[n]`` is any digit. Example ``v0.0.1-alpha``.