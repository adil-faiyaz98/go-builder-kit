# Security Policy

## Supported Versions

We currently support the following versions with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0.0 | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability within Go Builder Kit, please send an email to adilmd98@gmail.com. All security vulnerabilities will be promptly addressed.

Please include the following information in your report:

- Type of issue (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit the issue

This information will help us triage your report more quickly.

## Security Update Process

When we receive a security bug report, we will:

1. Confirm the problem and determine the affected versions.
2. Audit code to find any potential similar problems.
3. Prepare fixes for all supported versions.
4. Release new versions as soon as possible.

## Public Disclosure Timing

A public disclosure date is negotiated by the Go Builder Kit team and the bug submitter. We prefer to fully disclose the bug as soon as possible once a user mitigation is available. It is reasonable to delay disclosure when the bug or the fix is not yet fully understood, the solution is not well-tested, or for vendor coordination. The timeframe for disclosure is from immediate (especially if already publicly known) to a few weeks. For vulnerabilities with a straightforward mitigation, we expect report date to disclosure date to be on the order of 7 days. The Go Builder Kit team holds the final say when setting a disclosure date.
