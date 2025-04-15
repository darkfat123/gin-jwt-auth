<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExeGZjandja2RleXFyM3NlbnRqbTR2emRybjB0emtxN3Y4dGZwdWdieCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/myAFzJ8hJnlJiMN4hB/giphy.gif" alt="Logo" height="100">
  </a>

<h3 align="center">Gin JWT Auth API</h3>

  <p align="center">
    A Go (Gin) starter template for auth – includes register, login, and token-protected routes to save time when starting new projects.
    <br />
    <br />
    <a href="https://github.com/darkfat123/typing-race-web-multiplayer/issues">🚨 Report Bug</a>
    ·
    <a href="https://github.com/darkfat123/typing-race-web-multiplayer/issues">✉️ Request Feature</a>
    .
    <a href="https://github.com/darkfat123/typing-race-web-multiplayer?tab=readme-ov-file#-getting-started-for-development-only">🚀 Getting Started</a>
  </p>
</div>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">✨ Features:</h3>

  * 🔒 Hashes password securely with `bcrypt`
  * 🛢️ Uses `sqlx` for communicate with PostgreSQL
  * 📝 Stores `username`,` email`, and hashed password in database
  * ✅ Validates login credentials and returns a `JWT token`
  * 🔐 Protects all `/api/*` endpoints from unauthorized access
  * 📥 Reads `Authorization: Bearer <token>` header to verify identity
  * 👤 Provides user endpoint: `GET /api/users/:id`
  * 🌐 Handles CORS with custom origin configuration
  * 📄 Uses `zap` logger for structured and leveled logging

</br>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

### 🚀 Getting Started (for development only)

#### 1. Clone the project
```bash
git clone https://github.com/darkfat123/gin-jwt-auth.git
cd gin-jwt-auth
```
#### 2. Backend
```bash
go mod tidy
go run cmd/main.go
```

#### 3. Environment Variables
```bash
# root .env
DB_USER=<YOUR_DB_USER>
DB_PASSWORD=<YOUR_DB_PASSWORD>
DB_HOST=<YOUR_DB_HOST>
DB_PORT=<YOUR_DB_PORT>
DB_NAME=<YOUR_DB_NAME>
SERVER_PORT=8080
JWT_SECRET=<YOUR_SECRET>
ENVIRONMENT=dev
ALLOWED_ORIGINS=<YOUR_ORIGINS>
```
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

### 🚨 Example Usage
```bash
POST /auth/register
{
  "username": "test01",
  "email": "test01@example.com",
  "password": "secure123test"
}

POST /auth/login
{
  "username": "test01",
  "password": "secure123test"
}
# → Returns JWT Token

GET /api/users/1
# → Requires Bearer Token
Authorization: Bearer <your-jwt-token>

```

<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">🖥️ Programming languages and tools:</h3>

- Backend
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" />
  </a>
</p>

- Databases
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=postgresql" />
  </a>
</p>

- Tools
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=git,github,vscode,postman" />
  </a>
</p>

<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3> Connect with me 🎊: <h3>
  <a href="https://www.linkedin.com/in/supakorn-yookack-39a730289/">
   <img align="left" alt="Supakorn Yookack | Linkedin" width="30px" src="https://www.vectorlogo.zone/logos/linkedin/linkedin-icon.svg" />
  </a>
  <a href="mailto:supakorn.yookack@gmail.com">
    <img align="left" alt="Supakorn Yookack | Gmail" width="32px" src="https://www.vectorlogo.zone/logos/gmail/gmail-icon.svg" />
  </a>
  <a href="https://medium.com/@yookack_s">
    <img align="left" alt="Supakorn Yookack | Medium" width="32px" src="https://www.vectorlogo.zone/logos/medium/medium-tile.svg" />
  </a>
   <a href="https://www.facebook.com/supakorn.yookaek/">
    <img align="left" alt="Supakorn Yookack | Facebook" width="32px" src="https://www.vectorlogo.zone/logos/facebook/facebook-tile.svg" />
  </a>
   <a href="https://github.com/darkfat123">
    <img align="left" alt="Supakorn Yookack | Github" width="32px" src="https://www.vectorlogo.zone/logos/github/github-tile.svg" />
  </a>
    <p align="right" > Created by <a href="https://github.com/darkfat123">darkfat</a></p> <p align="right" > <img src="https://komarev.com/ghpvc/?username=darkfat123&label=Profile%20views&color=0e75b6&style=flat" alt="darkfat123" /> </p>
