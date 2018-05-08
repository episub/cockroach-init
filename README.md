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

# Notes:

* Files are imported in an order sorted by filename.  Therefore, prepend files with, e.g., a number like 00001, so that they will be imported in the desired order
* Specify folder containing files via environment variable SCRIPTS_FOLDER
* Use above generated image as you would the base cockroach image
