# Demo of a Website Analytics using Cube.js

Demo of the [Cube.js](https://cube.dev/) Capabilities. 

This is a example of a very simple way to push click stream events into a backend and then use Cube.js to create dashboards.

### [analytics-backend](analytics-backend/README.md)

Small GO backend that exposes a REST end point to consume events and stores them into PostgreSQL.

### [analytics-cubejs](analytics-cubejs/README.md)

[Cube.js](https://cube.dev/) API with the schema for the events.

### [analytics-frontend](analytics-frontend/README.md)

Very simple Javascript file that capture clicks and page views and send them to the backend using the REST end point.
