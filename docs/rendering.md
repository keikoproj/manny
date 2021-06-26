# Rendering

Manny is not cognizant of the underlying CloudFormation that it is packaging up. When a user submits a template, whether
in JSON or YAML format, Manny will simply read the file bytes and attach them to the Manny config. All user related input
goes untouched in that regard. 

Formatting _does_ apply to how Manny sends the Cloud Resource to ArgoCD. Manny can send in JSON or YAML formats, although,
traditional Kubernetes styled JSON (which is not actually valid JSON) seems to cause problems with ArgoCD.

By default, Manny uses YAML formatting.