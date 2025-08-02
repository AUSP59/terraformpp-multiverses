# 🔐 Security Policy

## 📅 Supported Versions

We release security patches for the latest major version of OmniForge.

| Version      | Supported          |
|--------------|--------------------|
| ULTRA PRIME  | ✅ Actively supported |
| Others       | ❌ No longer supported |

---

## 📣 Reporting a Vulnerability

If you discover a security vulnerability in OmniForge, please report it **privately and responsibly**.

### 📬 Contact Options:

- 📧 **Email**: security@omniforge.io *(placeholder – replace with real address)*
- 🐞 **GitHub Security Advisory**: [Open a confidential issue](https://github.com/AUSP59/terraformpp-multiverses/security/advisories)

Please **do not open public issues** that might expose the vulnerability.

---

## 🔎 What to Include

When reporting a vulnerability, please include:

- A clear description of the issue
- Steps to reproduce it (if applicable)
- Any logs, stack traces or code snippets
- The version of OmniForge used
- Your name or handle (optional, for credit)

We aim to respond within **72 hours** and issue patches promptly after triage.

---

## 🧪 Security Tools Used

- ✅ [GoSec](https://github.com/securego/gosec) — Static security analysis
- ✅ [Dependabot](https://docs.github.com/en/code-security) — Dependency audit
- ✅ [CodeQL](https://codeql.github.com/) — Code vulnerability scanning
- ✅ SPDX & REUSE compliance for open source licenses

---

## 🧾 Disclosure Policy

We practice **coordinated disclosure**. After a fix is released:

- We will publish a **GitHub Security Advisory**
- We will credit the reporter (if permitted)
- We will update `CHANGELOG.md` and `SECURITY_AUDIT.md`

---

## 🧰 Preventive Measures

OmniForge is built with **security-first architecture**, including:

- RBAC (Role-Based Access Control)
- Optional JWT-based authentication
- Strong input validation
- OpenAPI schema enforcement
- No secrets stored in the repo
- CI/CD pipelines with linting, tests, and static analysis

---

## 🙏 Thank You

We appreciate responsible disclosures that help us make OmniForge safer for everyone.

---
