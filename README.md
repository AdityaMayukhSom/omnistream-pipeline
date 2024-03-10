## Pixelated Pipeline

The backend infrastructure of a **Pixelated Pen** written using GoLang.
Frontend is written using Svelte Kit (Pixelated Pen).

**Currently the backend uses these libraries:**

- `The Go Standard Library`
- `gorilla/mux`
- `github.com/lib/pq`
- `google/uuid`
- `dgrijalva/jwt-go`
- `github.com/charmbracelet/log`
- `github.com/spf13/viper`

## Features currently implemented

- REST API
  - `GET` retrieves posts.
  - `POST` writes new posts.
  - `PUT` updates posts.
  - `DELETE` deletes posts.
- PostgreSQL database to store all posts and user info.
  - `users` table stores the user account info.
  - `posts` table stores the post data.
- User authentication
  - Users who sign up have their account data recorded to the PostgreSQL database.
  - Users who log in have their form data checked by comparing it to the database records, then receives a JWT access token if authenticated.
  - Users also receive refresh tokens that last for 7 days. This allows the users to stay logged in even if their access token has been invalidated after 15 minutes.
  - Protected endpoints require a valid access token to access.
  - Each token has a UUID, allowing one user to be logged into multiple devices at a time.
- Where to store the JWT
  - Access token is stored in memory as a Svelte store, while the refresh token is stored in a HttpOnly cookie.
  - Backend handles everything CORS, and this allows cookies to be sent over cross origin (same domain, but different ports).
- Frontend
  - Frontend now starts a timer that counts until the access token expiration time, and when it does, it sends a request to the backend for a refresh token.
  - Clicking on a post title links to a page dedicated for that single post.
  - Clicking on the username in the navbar after login redirects to the user profile page.
  - Each post component shows the title, author and the content of the post.
  - Users can write their own posts.
  - Posts are shortened in the homepage if they exceed around 100 words or so.
- Security
  - Passwords are hashed using bcrypt.
