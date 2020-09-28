# Cloud Build Bestiary

This repo contains cloudbuild.yaml configs that exercise esoteric features of
the service.

To run them all:

```
go test ./... -v
```

To run one test (e.g., `workspace.yaml`):

```
go test ./... -v -run=.*/workspace.yaml
```

or 

```
gcloud builds submit --config=workspace.yaml
```
