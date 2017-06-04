# Changelog

## 0.8.0
* TBD

## 0.7.1
* Updated versions in makefile and specfile properly

## 0.7.0
* Changed template specification from position argument to -template flag
* Extracted template data parsing into subpackage to make adding new data sources easier

## 0.6.0
* Rename project to cauldron
* Improved version string handling:
  * Extracted version string handling into subpackage
  * Included Git details when built from checkout to make development easier
  * Stopped including build time in binary, for more consistent builds
* Added the -file flag for rendering to a specified file.
* Added the -exec flag for running a command after template rendering.

## 0.5.0
* Template data can now be read from JSON files with the -json flag. Template data provided on the command line is ignored when using this flag.
* Added the -version flag.
* Added a simple build system.
* Added a spec file for RPM builds.
* Introduced a changelog.

## 0.4.0
* Added the Sprig template library 

## 0.3.0
* Template files can now be written to in-place with the -inplace flag.
* Better error messages on argument and template-parsing failures.

## 0.2.0
* Arrays and maps can now be read from the command line in the form of JSON-formatted strings.

## 0.1.0
* First working version.
