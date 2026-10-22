# Security remediation (2026-10-22)

Go 1.26.10 is required by go.mod and used by Docker and CI. Removed routing logs containing full request arguments, email addresses/bodies, and chat prompts/responses. Dependabot auto-merge requires the successful workflow to validate the current PR head and merges with an expected SHA.

Validation: go test ./src/... passed. govulncheck ./src/... found zero vulnerable symbols and zero vulnerable imported packages.

Residual findings: github.com/aws/aws-sdk-go v1.55.8 includes GO-2022-0646 and GO-2022-0635 in the S3 Crypto SDK, for which the database lists no fixed version. The application does not import or call those packages. Eliminating the module-level findings requires migrating the legacy AWS SDK and DynamoDB adapter to SDK v2; this compatibility migration is outside this patch.
