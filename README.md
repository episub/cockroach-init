# cockroach-init

Small program to initialise a docker database with SQL files.

To use, create a Dockerfile for your project, and place scripts into the specified folder.  For example:

```
FROM episub/cockroach-init as init

FROM cockroachdb/cockroach:v2.0.1
COPY --from=init /built/* /importer/
ENV SCRIPTS_FOLDER /importer/Scripts
ADD Scripts /importer/Scripts
ENTRYPOINT ["/importer/init.sh"]
```
