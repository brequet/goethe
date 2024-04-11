# GOETHE: A Fullstack Template

GOETHE is a fullstack template designed to streamline the development of modern web applications. It integrates a Golang backend with a Svelte frontend, leveraging Tailwind CSS for styling and Shadcn for additional UI components. The project is structured to serve a Single Page Application (SPA) with routing capabilities provided by `svelte-spa-router`.

## Features

- **Golang Backend**: Efficiently serves the frontend SPA and any API endpoints.
- **Svelte Frontend**: A reactive UI framework for building user interfaces.
- **Tailwind CSS**: Utilizes utility-first CSS for rapid UI development.
- **Shadcn**: Enhances the UI with additional components and styles.
- **svelte-spa-router**: Handles client-side routing for the SPA.
- **Makefile**: Simplifies the build process by automating frontend and backend builds.

## Getting Started

To get started with GOETHE, follow these steps:

1. **Clone the Repository**

2. **Install Dependencies**:
   - For the frontend, navigate to the frontend directory and run:

```sh
pnpm install
```

- Ensure you have Go installed for the backend.

3. **Build and Run**:
   - Use the provided Makefile to build both the frontend and backend:

```sh
make
```

- This command compiles the Svelte frontend, injects it into the Golang backend, and starts the server.

## Development

During development, you can run the frontend and backend separately for a more interactive experience.
