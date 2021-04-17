#  Analytics Backend

GO backend to store events.


Endpoint: `POST /events`

Body:
```
{
    "event_type": "view",
    "page_path": "/path/",
    "is_admin": false,
    "click_text": "here",
    "click_id": "",
    "click_class": "",
    "event_time" : "2021-03-10T10:56:28.200Z"
}
```

## RUN

- `made db`: Starts PostgreSQL instance
- `make run`: Run server on `localhost:8000`
- `make stop`: Stops PostgreSQL
