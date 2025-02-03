## без сертификата (локалка):
http://localhost:3000 {
  root * /usr/share/caddy
  file_server
  encode gzip

  @root_path path /
  respond @root_path "Caddy is running on port 3000" 200

  handle_path /api/* {
      reverse_proxy backend:8080
  }

  try_files {path} /index.html
}
 ### OR

http://0.0.0.0:3000 {
root * /usr/share/caddy
file_server
encode gzip

    handle_path /api/* {
        reverse_proxy backend:8080
    }

    try_files {path} /index.html
}

## Без SSL для сервера:

## редирект (когда все работает:)

http://garbagegogoriki.ru {
redir https://garbagegogoriki.ru{uri}
}

https://garbagegogoriki.ru {
root * /usr/share/caddy
file_server
encode gzip

    @root_path path /
    respond @root_path "Caddy is running on port 3000" 200

    handle_path /api/* {
        reverse_proxy backend:8080
    }

    try_files {path} /index.html
}


## без редиректа:
http://garbagegogoriki.ru {
root * /usr/share/caddy
file_server
encode gzip

    handle_path /.well-known/acme-challenge/* {
        respond "ACME Challenge OK" 200
    }

    handle_path /api/* {
        reverse_proxy backend:8080
    }

    try_files {path} /index.html
}

https://garbagegogoriki.ru {
root * /usr/share/caddy
file_server
encode gzip

    handle_path /api/* {
        reverse_proxy backend:8080
    }

    try_files {path} /index.html
}

